//+build mongo

package db

import (
	"context"
	"flag"
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
)

func init() {
	flag.StringVar(&host, "mongo-host", "127.0.0.1:27017", "Mongo host. default 127.0.0.1:27017")
	flag.BoolVar(&direct, "mongo-direct", true, "Direct connections cannot be made if multiple hosts are specified or an SRV URI is used.")
	flag.StringVar(&authMechanism, "mongo-auth-mechanism", "SCRAM-SHA-1", `The mechanism to use for authentication. Supported values include "SCRAM-SHA-256", "SCRAM-SHA-1", "MONGODB-CR", "PLAIN", "GSSAPI", "MONGODB-X509", and "MONGODB-AWS". This can also be set through the "authMechanism" URI option.`)
	flag.StringVar(&authSource, "mongo-auth-source", "admin", `The name of the database to use for authentication. This defaults to "$external" for MONGODB-X509, GSSAPI, and PLAIN and "admin" for all other mechanisms. This can also be set through the "authSource" URI option.`)
	flag.StringVar(&username, "mongo-username", "", "The username for authentication. This can also be set through the URI as a username:password pair before the first @ character.")
	flag.StringVar(&password, "mongo-password", "", "The password for authentication. This must not be specified for X509 and is optional for GSSAPI authentication.")
	flag.BoolVar(&passwordSet, "mongo-password-set", false, "For GSSAPI, this must be true if a password is specified, even if the password is the empty string, and false if no password is specified, indicating that the password should be taken from the context of the running process.")
	flag.StringVar(&database, "mongo-database", "", "Database name.")

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

		(&Sum{}).GetTag(): sum,
		(&Avg{}).GetTag(): avg,
		(&Min{}).GetTag(): min,
		(&Max{}).GetTag(): max,
	})

	structor.SetDdlConstructors(map[string]structor.Expr{
		(&Create{}).GetTag():      create,
		(&Add{}).GetTag():         add,
		(&Modify{}).GetTag():      modify,
		(&Index{}).GetTag():       index,
		(&Unique{}).GetTag():      unique,
		(&DropIndexes{}).GetTag(): dropIndexes,
	})
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
	bsons = append(bsons, bson.M{
		"$group": d.builder.Agg,
	}, bson.M{
		"$project": bson.M{"_id": 0},
	})

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

func (d *Dorm) Build(table string, expr structor.Constructor) dorm.Dept {
	dorm := &Dorm{
		db:      d.db,
		C:       d.db.Collection(table),
		builder: new(MONGO),
	}
	expr.Build(table, dorm.builder)
	return dorm
}

func (d *Dorm) Exec(c context.Context) error {
	if d.builder.IsCreate {
		jsonSchema := bson.M{
			"bsonType": "object",
			"required": []string{"_id"},
			"properties": bson.M{
				"_id": bson.M{
					"bsonType": "string",
				},
			},
		}
		validator := bson.M{
			"$jsonSchema": jsonSchema,
		}
		opts := options.CreateCollection().SetValidator(validator)

		err := d.db.CreateCollection(c, d.builder.ColName, opts)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *Dorm) Index(ctx context.Context, name string) error {
	_, err := d.C.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    d.builder.Keys,
		Options: options.Index().SetName(name).SetUnique(d.builder.IsUnique),
	})

	return err
}

func (d *Dorm) DropIndexes(ctx context.Context) error {
	for _, name := range d.builder.Indexes {
		_, err := d.C.Indexes().DropOne(ctx, name)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *Dorm) marshal(entities interface{}) error {
	// nothing to do
	return nil
}

func (d *Dorm) unmarshal(entities interface{}) error {
	var doUnmarshal = func(entity map[string]interface{}) error {
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
	Vars     bson.M
	Agg      bson.M
	Keys     bson.D
	IsUnique bool
	Indexes  []string
	IsCreate bool
	ColName  string
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
	m.Keys = append(m.Keys, bson.E{field, 1})
}

func (m *MONGO) Unique(unique bool) {
	m.IsUnique = unique
}

func (m *MONGO) IndexName(names []string) {
	m.Indexes = names
}

func (m *MONGO) Create(f bool, name ...string) {
	m.IsCreate = f
	if len(name) > 0 {
		m.ColName = name[0]
	}
}
