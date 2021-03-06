//+build mongo

package db

import (
	"context"
	"flag"
	"fmt"
	"reflect"
	"strings"

	mgc "github.com/quanxiang-cloud/cabin/tailormade/db/mongo"
	"github.com/quanxiang-cloud/structor/internal/dorm"
	"github.com/quanxiang-cloud/structor/internal/dorm/clause"
	"github.com/quanxiang-cloud/structor/internal/dorm/structor"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var _ clause.Builder = &MONGO{}

var (
	host          string
	direct        bool
	authMechanism string
	authSource    string
	username      string
	password      string
	passwordSet   bool
	database      string
	prefixes      VarPrefixArray
)

type VarPrefixArray []string

func (v *VarPrefixArray) String() string {
	return fmt.Sprint(*v)
}

func (v *VarPrefixArray) Set(s string) error {
	*v = append(*v, strings.Split(s, ",")...)
	return nil
}

func init() {
	flag.StringVar(&host, "mongo-host", "127.0.0.1:27017", "Mongo host. default 127.0.0.1:27017")
	flag.BoolVar(&direct, "mongo-direct", true, "Direct connections cannot be made if multiple hosts are specified or an SRV URI is used.")
	flag.StringVar(&authMechanism, "mongo-auth-mechanism", "SCRAM-SHA-1", `The mechanism to use for authentication. Supported values include "SCRAM-SHA-256", "SCRAM-SHA-1", "MONGODB-CR", "PLAIN", "GSSAPI", "MONGODB-X509", and "MONGODB-AWS". This can also be set through the "authMechanism" URI option.`)
	flag.StringVar(&authSource, "mongo-auth-source", "admin", `The name of the database to use for authentication. This defaults to "$external" for MONGODB-X509, GSSAPI, and PLAIN and "admin" for all other mechanisms. This can also be set through the "authSource" URI option.`)
	flag.StringVar(&username, "mongo-username", "", "The username for authentication. This can also be set through the URI as a username:password pair before the first @ character.")
	flag.StringVar(&password, "mongo-password", "", "The password for authentication. This must not be specified for X509 and is optional for GSSAPI authentication.")
	flag.BoolVar(&passwordSet, "mongo-password-set", false, "For GSSAPI, this must be true if a password is specified, even if the password is the empty string, and false if no password is specified, indicating that the password should be taken from the context of the running process.")
	flag.StringVar(&database, "mongo-database", "", "Database name.")
	flag.Var(&prefixes, "prefixes", "The prefix of the row name. It will be used to serial the table content. It can be multiple. default []")

	// expressions
	clause.SetDmlExpressions(map[string]clause.Expr{
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

		(&Sum{}).GetTag():   sum,
		(&Avg{}).GetTag():   avg,
		(&Min{}).GetTag():   min,
		(&Max{}).GetTag():   max,
		(&Count{}).GetTag(): count,

		(&Bool{}).GetTag(): bool1,
	})

	structor.SetCreateExpr(create)
	structor.SetAddExpr(add)
	structor.SetModifyExpr(modify)
	structor.SetPrimaryExpr(primary)
	structor.SetIndexExpr(index)
	structor.SetUniqueExpr(unique)
	structor.SetDropIndexExpr(dropIndex)
}

// Dorm dorm
type Dorm struct {
	db *mongo.Database
	C  *mongo.Collection

	builder *MONGO
	opt     *options.FindOptions
}

type option struct {
	offset int64
	limit  int64
	sort   bson.D
}

func New() (*Dorm, error) {
	conf := &mgc.Config{
		Hosts:  strings.Split(host, ","),
		Direct: direct,
	}

	conf.Credential.AuthMechanism = authMechanism
	conf.Credential.AuthSource = authSource
	conf.Credential.Username = username
	conf.Credential.Password = password
	conf.Credential.PasswordSet = passwordSet

	client, err := mgc.New(conf)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		return nil, err
	}

	return &Dorm{
		db: client.Database(database),
	}, nil
}

func (d *Dorm) Table(tablename string) dorm.Dorm {
	return &Dorm{
		C:       d.db.Collection(tablename),
		builder: new(MONGO),
		opt:     new(options.FindOptions),
	}
}

func (d *Dorm) Where(expr clause.Expression) dorm.Dorm {
	expr.Build(d.builder)
	return d
}

func (d *Dorm) Select(exprs ...clause.Expression) dorm.Dorm {
	bsons := make([]bson.M, 0)
	if vars := d.builder.Vars; len(vars) != 0 {
		bsons = append(bsons, bson.M{"$match": vars})
	}

	for _, expr := range exprs {
		expr.Build(d.builder)
	}

	bsons = append(bsons, bson.M{
		"$group": d.builder.Agg,
	}, bson.M{
		"$project": bson.M{"_id": 0},
	})
	return d
}

func (d *Dorm) Limit(limit int64) dorm.Dorm {
	d.opt = d.opt.SetLimit(limit)
	return d
}

func (d *Dorm) Offset(offset int64) dorm.Dorm {
	d.opt = d.opt.SetSkip(offset)
	return d
}

func (d *Dorm) Order(arr ...string) dorm.Dorm {
	sort := make(bson.D, 0, len(arr))
	for _, elem := range arr {
		if strings.HasPrefix(elem, "-") {
			sort = append(sort, bson.E{Key: elem[1:], Value: -1})
			continue
		}
		sort = append(sort, bson.E{Key: elem, Value: 1})
	}

	d.opt = d.opt.SetSort(sort)
	return d
}

func (d *Dorm) find(ctx context.Context) ([]map[string]interface{}, error) {
	cursor, err := d.C.Find(ctx, d.builder.Vars, d.opt)
	if err != nil {
		return nil, err
	}

	result := make([]map[string]interface{}, 0)
	err = cursor.All(ctx, &result)
	if err == mongo.ErrNoDocuments || err == mongo.ErrNilDocument {
		return nil, nil
	}
	d.unmarshal(result)
	return result, err
}

func (d *Dorm) agg(ctx context.Context) ([]map[string]interface{}, error) {
	bsons := make([]bson.M, 0)
	if vars := d.builder.Vars; len(vars) != 0 {
		bsons = append(bsons, bson.M{"$match": vars})
	}
	_, ok := d.builder.Agg["$count"]
	if ok {
		bsons = append(bsons, d.builder.Agg)
	} else {
		bsons = append(bsons, bson.M{
			"$group": d.builder.Agg,
		}, bson.M{
			"$project": bson.M{"_id": 0},
		})
	}
	cursor, err := d.C.Aggregate(ctx, bsons, nil)
	if err != nil {
		return nil, err
	}
	result := make([]map[string]interface{}, 0)
	err = cursor.All(ctx, &result)

	if err == mongo.ErrNoDocuments || err == mongo.ErrNilDocument {
		return nil, nil
	}

	return result, err
}

func (d *Dorm) FindOne(ctx context.Context) (map[string]interface{}, error) {
	singleResult := d.C.FindOne(ctx, d.builder.Vars)

	result := make(map[string]interface{})
	err := singleResult.Decode(&result)
	if err == mongo.ErrNoDocuments || err == mongo.ErrNilDocument {
		return nil, nil
	}
	d.unmarshal(result)
	return result, err
}

func (d *Dorm) Find(ctx context.Context) ([]map[string]interface{}, error) {
	// FIXME Should configure `group by` syntax implementation
	if len(d.builder.Agg) != 0 {
		return d.agg(ctx)
	}

	return d.find(ctx)
}

func (d *Dorm) Count(ctx context.Context) (int64, error) {
	return d.C.CountDocuments(ctx, d.builder.Vars)
}

func (d *Dorm) Insert(ctx context.Context, entities ...interface{}) (int64, error) {
	ret, err := d.C.InsertMany(ctx, entities)
	return int64(len(ret.InsertedIDs)), err
}

func (d *Dorm) Update(ctx context.Context, entity interface{}) (int64, error) {
	result, err := d.C.UpdateMany(ctx, d.builder.Vars,
		bson.M{
			"$set": entity,
		},
	)
	return result.ModifiedCount, err
}

func (d *Dorm) Delete(ctx context.Context) (int64, error) {
	result, err := d.C.DeleteMany(ctx, d.builder.Vars)
	return result.DeletedCount, err
}

// ************************************************************************************************************

func (d *Dorm) Create(ctx context.Context, c structor.Constructor) error {
	builder := &MONGO{}
	c.Build(builder)
	opts := options.CreateCollection().SetValidator(builder.Schema)
	err := d.db.CreateCollection(ctx, c.GetTable(), opts)
	return err
}

func (d *Dorm) Add(ctx context.Context, c structor.Constructor) error {
	return nil
}

func (d *Dorm) Modify(ctx context.Context, c structor.Constructor) error {
	return nil
}

func (d *Dorm) Primary(ctx context.Context, c structor.Constructor) error {
	// BUG: duplicate index in collection, its automatically created.
	// builder := &MONGO{}
	// c.Build(builder)
	// col := d.db.Collection(c.GetTable())
	// _, err := col.Indexes().CreateOne(ctx, mongo.IndexModel{
	// 	Keys:    builder.Keys,
	// 	Options: options.Index().SetUnique(true),
	// })
	return nil
}

func (d *Dorm) Index(ctx context.Context, c structor.Constructor) error {
	builder := &MONGO{}
	c.Build(builder)
	col := d.db.Collection(c.GetTable())
	_, err := col.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    builder.Keys,
		Options: options.Index().SetName(c.GetIndex()).SetUnique(false),
	})

	return err
}

func (d *Dorm) Unique(ctx context.Context, c structor.Constructor) error {
	builder := &MONGO{}
	c.Build(builder)
	col := d.db.Collection(c.GetTable())
	_, err := col.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    builder.Keys,
		Options: options.Index().SetName(c.GetIndex()).SetUnique(true),
	})

	return err
}

func (d *Dorm) DropIndex(ctx context.Context, c structor.Constructor) error {
	builder := &MONGO{}
	c.Build(builder)
	col := d.db.Collection(c.GetTable())

	_, err := col.Indexes().DropOne(ctx, c.GetIndex())
	return err
}

func (d *Dorm) marshal(entities interface{}) error {
	// nothing to do
	return nil
}

func (d *Dorm) unmarshal(entities interface{}) error {
	doUnmarshal := func(entity map[string]interface{}) error {
		for key, value := range entity {
			// TODO: Whether to judge the type of interface {}
			pa, ok := value.(primitive.A)
			if !ok {
				continue
			}
			reflect.ValueOf(entity).
				SetMapIndex(reflect.ValueOf(key), reflect.ValueOf([]interface{}(pa)))
		}
		return nil
	}

	switch v := entities.(type) {
	case []map[string]interface{}:
		for _, entity := range v {
			err := doUnmarshal(entity)
			if err != nil {
				return err
			}
		}
	case map[string]interface{}:
		err := doUnmarshal(v)
		if err != nil {
			return err
		}
	}

	return nil
}

// MONGO mongo
type MONGO struct {
	Vars bson.M
	Agg  bson.M

	Schema bson.M
	Keys   bson.D
}

// WriteString write string
func (m *MONGO) WriteString(str string) (int, error) {
	m.WriteQuoted(str)
	return len(str), nil
}

// WriteByte write byte
func (m *MONGO) WriteByte(c byte) error {
	_, err := m.WriteString(string(c))
	return err
}

// WriteQuoted write quoted
func (m *MONGO) WriteQuoted(field string) {
	m.Vars = bson.M{
		field: m.Vars,
	}
}

// AddVar add var
func (m *MONGO) AddVar(value interface{}) {
	for key := range m.Vars {
		m.Vars[key] = value
		return
	}
}

func (m *MONGO) GetVar() interface{} {
	return m.Vars
}

// WriteQuotedAgg WriteQuotedAgg
func (m *MONGO) WriteQuotedAgg(field string) {
	m.Agg = bson.M{
		field: m.Agg,
	}
}

// AddAggVar AddAggVar
func (m *MONGO) AddAggVar(key string, value interface{}) {
	if m.Agg == nil {
		m.Agg = bson.M{}
		if key != "$count" {
			m.Agg["_id"] = "null"
		}
	}
	m.Agg[key] = value
}

func (m *MONGO) WriteRaw(field string) {
	m.Schema = bson.M{
		field: nil,
	}
}

func (m *MONGO) AddRawVal(content interface{}) {
	for k := range m.Schema {
		m.Schema[k] = content
		return
	}
}

func (m *MONGO) AddIndex(field string) {
	m.Keys = append(m.Keys, bson.E{field, 1})
}
