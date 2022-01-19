//+build mysql

package db

import (
	"bytes"
	"context"
	"flag"
	"fmt"

	"github.com/quanxiang-cloud/cabin/logger"
	msc "github.com/quanxiang-cloud/cabin/tailormade/db/mysql"
	"github.com/quanxiang-cloud/structor/internal/dorm"
	"github.com/quanxiang-cloud/structor/internal/dorm/clause"
	"gorm.io/gorm"
)

var (
	host     string
	db       string
	username string
	password string

	log      bool
	logLevel int

	maxIdleConns int
	maxOpenConns int
)

func init() {
	flag.StringVar(&host, "mysql-host", "127.0.0.1:3306", "MySQL host. default 127.0.0.1:3306")
	flag.StringVar(&db, "mysql-database", "", "MySQL database name")
	flag.StringVar(&username, "mysql-user", "", "The username for authentication")
	flag.StringVar(&password, "mysql-password", "", "The password for authentication")
	flag.BoolVar(&log, "mysql-log", false, "")
	flag.IntVar(&logLevel, "mysql-log-level", -1, "The log level. it cannot be make if disable log. Level options: -1 debug, 0 info, 1 warn, 2 error, 3 dPanic, 4 panic, 5 fatal. Default log level is debugLevel")
	flag.IntVar(&maxIdleConns, "mysql-maxIdleConns", 10, "The maximum number of connections in the idle connection pool. default 10")
	flag.IntVar(&maxOpenConns, "mysql-maxOpenConns", 20, "The maximum number of open connections to the database. default 20")

	clause.SetExpressions(map[string]clause.Expr{
		(&Terms{}).GetTag(): terms,
		(&Match{}).GetTag(): match,
		(&Term{}).GetTag():  term,
		(&LT{}).GetTag():    lt,
		(&LTE{}).GetTag():   lte,
		(&GT{}).GetTag():    gt,
		(&GTE{}).GetTag():   gte,

		(&MUST{}).GetTag():    must,
		(&MUSTNOT{}).GetTag(): mustNot,
		(&SHOULD{}).GetTag():  should,
		(&RANGE{}).GetTag():   range1,

		(&Sum{}).GetTag(): sum,
		(&Avg{}).GetTag(): avg,
		(&Min{}).GetTag(): min,
		(&Max{}).GetTag(): max,
	})
}

type Dorm struct {
	db *gorm.DB

	builder *MYSQL
	opt     *option
}

func New() (*Dorm, error) {
	conf := msc.Config{
		Host:     host,
		DB:       db,
		User:     username,
		Password: password,
		Log:      log,

		MaxIdleConns: maxIdleConns,
		MaxOpenConns: maxOpenConns,
	}

	log := logger.New(&logger.Config{
		Level: logLevel,
	})

	db, err := msc.New(conf, log)
	if err != nil {
		return nil, err
	}
	return &Dorm{
		db: db,
	}, nil
}

type option struct {
	offset int
	limit  int
	sort   []string
}

func (d *Dorm) Table(tablename string) dorm.Dorm {
	builder := &MYSQL{
		val: &ConditionVal{
			Condition: bytes.Buffer{},
		},
		agg: bytes.Buffer{},
	}

	return &Dorm{
		db:      d.db.Table(tablename),
		builder: builder,
		opt:     new(option),
	}
}

func (d *Dorm) Where(expr clause.Expression) dorm.Dorm {
	expr.Build(d.builder)
	return d
}

func (d *Dorm) Select(exprs ...clause.Expression) dorm.Dorm {
	for _, expr := range exprs {
		expr.Build(d.builder)
	}
	return d
}

func (d *Dorm) Limit(limit int64) dorm.Dorm {
	d.db = d.db.Limit(int(limit))
	return d
}

func (d *Dorm) Offset(offset int64) dorm.Dorm {
	d.db = d.db.Offset(int(offset))
	return d
}

func (d *Dorm) Order(order ...string) dorm.Dorm {
	d.db = d.db.Order(order)
	return d
}

func (d *Dorm) FindOne(ctx context.Context) (map[string]interface{}, error) {
	ret := make(map[string]interface{})
	err := d.db.Where(d.builder.val.Condition.String(), d.builder.val.vars...).Find(&ret).Error
	return ret, err
}

func (d *Dorm) Find(ctx context.Context) ([]map[string]interface{}, error) {
	ret := make([]map[string]interface{}, 0)
	if d.builder.agg.Len() != 0 {
		d.db = d.db.Select(d.builder.agg.String())
	}
	err := d.db.Where(d.builder.val.Condition.String(), d.builder.val.vars...).Find(&ret).Error
	return ret, err
}

func (d *Dorm) Count(ctx context.Context) (int64, error) {
	var ret int64
	err := d.db.Where(d.builder.val.Condition.String(), d.builder.val.vars...).Count(&ret).Error
	return ret, err
}

func (d *Dorm) Insert(ctx context.Context, entities ...interface{}) (int64, error) {
	var ret int64 = 0
	err := d.db.CreateInBatches(entities, len(entities)).Error
	ret = int64(len(entities))
	return ret, err
}

func (d *Dorm) Update(ctx context.Context, entity interface{}) (int64, error) {
	affected := d.db.Where(d.builder.val.Condition.String(), d.builder.val.vars...).Updates(entity).RowsAffected
	return affected, nil
}

func (d *Dorm) Delete(ctx context.Context) (int64, error) {
	affected := d.db.Where(d.builder.val.Condition.String(), d.builder.val.vars...).Delete(nil).RowsAffected
	return affected, nil
}

type MYSQL struct {
	table string

	val *ConditionVal
	agg bytes.Buffer
}

type ConditionVal struct {
	Condition bytes.Buffer
	vars      []interface{}
}

func (m *MYSQL) WriteQuoted(field string) {
	m.val.Condition.WriteString(field)
}

func (m *MYSQL) WriteByte(c byte) error {
	_, err := m.WriteString(string(c))
	return err
}

func (m *MYSQL) WriteString(str string) (int, error) {
	m.WriteQuoted(str)
	return len(str), nil
}

func (m *MYSQL) AddVar(value interface{}) {
	m.val.vars = append(m.val.vars, value)
}

func (m *MYSQL) GetVar() interface{} {
	return m.val
}

func (m *MYSQL) WriteQuotedAgg(field string) {
	m.agg.WriteString(field)
}

func (m *MYSQL) AddAggVar(key string, value interface{}) {
	if m.agg.Len() > 0 {
		m.agg.WriteString(", ")
	}
	m.agg.WriteString(fmt.Sprintf("%s %s", value, key))
}
