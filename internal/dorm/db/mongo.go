//+build mongo

package db

import (
	"context"
	"flag"
	"fmt"
	"strings"

	mgc "github.com/quanxiang-cloud/cabin/tailormade/db/mongo"
	"github.com/quanxiang-cloud/structor/internal/dorm/clause"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	_ clause.Builder = &MONGO{}
)

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
	flag.StringVar(&host, "host", "127.0.0.1:27017", "Mongo host. default 127.0.0.1:27017")
	flag.BoolVar(&direct, "direct", true, "Direct connections cannot be made if multiple hosts are specified or an SRV URI is used.")
	flag.StringVar(&authMechanism, "auth-mechanism", "SCRAM-SHA-1", `The mechanism to use for authentication. Supported values include "SCRAM-SHA-256", "SCRAM-SHA-1", "MONGODB-CR", "PLAIN", "GSSAPI", "MONGODB-X509", and "MONGODB-AWS". This can also be set through the "authMechanism" URI option.`)
	flag.StringVar(&authSource, "auth-source", "admin", `The name of the database to use for authentication. This defaults to "$external" for MONGODB-X509, GSSAPI, and PLAIN and "admin" for all other mechanisms. This can also be set through the "authSource" URI option.`)
	flag.StringVar(&username, "username", "", "The username for authentication. This can also be set through the URI as a username:password pair before the first @ character.")
	flag.StringVar(&password, "password", "", "The password for authentication. This must not be specified for X509 and is optional for GSSAPI authentication.")
	flag.BoolVar(&passwordSet, "password-set", false, "For GSSAPI, this must be true if a password is specified, even if the password is the empty string, and false if no password is specified, indicating that the password should be taken from the context of the running process.")
	flag.StringVar(&database, "database", "", "Database name.")

	// expressions
	clause.SetExpressions(map[string]clause.Expr{
		(&IN{}).GetTag():    in,
		(&LIKE{}).GetTag():  like,
		(&EQUAL{}).GetTag(): equal,
		(&LT{}).GetTag():    lt,
		(&LTE{}).GetTag():   lte,
		(&GT{}).GetTag():    gt,
		(&GTE{}).GetTag():   gte,

		(&AND{}).GetTag(): and,
		(&OR{}).GetTag():  or,
		(&NOR{}).GetTag(): nor,

		(&Sum{}).GetTag(): sum,
		(&Avg{}).GetTag(): avg,
		(&Min{}).GetTag(): min,
		(&Max{}).GetTag(): max,
	})
}

// Dorm dorm
type Dorm struct {
	db *mongo.Database
	C  *mongo.Collection
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

func (d Dorm) Table(tablename string) *Dorm {
	return &Dorm{
		C: d.db.Collection(tablename),
	}
}

// FindOne find one entity
func (d Dorm) FindOne(ctx context.Context, expr clause.Expression) (clause.Data, error) {
	builder := &MONGO{}
	expr.Build(builder)

	singleResult := d.C.FindOne(ctx, builder.Vars)
	var result = make(clause.Data)
	err := singleResult.Decode(&result)
	if err == mongo.ErrNoDocuments || err == mongo.ErrNilDocument {
		return nil, nil
	}
	return result, err
}

// Find find entities
func (d Dorm) Find(ctx context.Context, expr clause.Expression, findOpt clause.FindOptions) ([]clause.Data, error) {
	builder := &MONGO{}
	expr.Build(builder)

	if builder.Agg != nil {
		return d.aggregation(ctx, builder)
	}

	fmt.Println(builder.Vars)
	opt := &options.FindOptions{}
	opt = opt.SetLimit(findOpt.Size)
	opt = opt.SetSkip((findOpt.Page - 1) * findOpt.Size)
	opt = opt.SetSort(sort(findOpt.Sort...))
	cursor, err := d.C.Find(ctx, builder.Vars, opt)
	if err != nil {
		return nil, err
	}
	result := make([]clause.Data, 0)
	err = cursor.All(ctx, &result)

	return result, err
}

// Count count entities
func (d Dorm) Count(ctx context.Context, expr clause.Expression) (int64, error) {
	builder := &MONGO{}
	expr.Build(builder)

	return d.C.CountDocuments(ctx, builder.Vars)
}

// Insert insert entities
func (d Dorm) Insert(ctx context.Context, entity ...interface{}) error {
	_, err := d.C.InsertMany(ctx, entity)
	return err
}

// Update update entities
func (d Dorm) Update(ctx context.Context, expr clause.Expression, entity interface{}) (int64, error) {
	builder := &MONGO{}
	expr.Build(builder)

	result, err := d.C.UpdateMany(ctx, builder.Vars, bson.M{"$set": entity})
	if err != nil {
		return 0, err
	}
	return result.ModifiedCount, nil
}

// Delete delete entities with condition
func (d Dorm) Delete(ctx context.Context, expr clause.Expression) (int64, error) {
	builder := &MONGO{}
	expr.Build(builder)

	result, err := d.C.DeleteMany(ctx, builder.Vars)
	if err != nil {
		return 0, err
	}
	return result.DeletedCount, nil
}

func (d Dorm) aggregation(ctx context.Context, b *MONGO) ([]clause.Data, error) {
	bsons := make([]bson.M, 0)
	if b.Vars != nil || len(b.Vars) != 0 {
		bsons = append(bsons, bson.M{"$match": b.Vars})
	}
	bsons = append(bsons, bson.M{
		"$group": b.Agg,
	}, bson.M{
		"$project": bson.M{"_id": 0},
	})
	fmt.Println(bsons)
	cursor, err := d.C.Aggregate(ctx, bsons, nil)
	if err != nil {
		return nil, err
	}
	result := make([]clause.Data, 0)
	err = cursor.All(ctx, &result)
	return result, err
}

func sort(array ...string) bson.D {
	sort := make(bson.D, 0, len(array))
	for _, elem := range array {
		if strings.HasPrefix(elem, "-") {
			sort = append(sort, bson.E{Key: elem[1:], Value: -1})
			continue
		}
		sort = append(sort, bson.E{Key: elem, Value: 1})
	}
	return sort
}

// MONGO mongo
type MONGO struct {
	Vars bson.M
	Agg  bson.M
}

func NewBuilder() clause.Builder {
	return &MONGO{}
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
