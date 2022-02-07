package ddl

import (
	"context"
	"fmt"
	"strings"

	"github.com/quanxiang-cloud/structor/internal/dorm"
	"github.com/quanxiang-cloud/structor/internal/dorm/structor"
)

type DDLService interface {
	Execute(context.Context, *ExecuteReq) (*ExecuteResp, error)
	Index(ctx context.Context, req *IndexReq) (*IndexResp, error)
	DropIndexes(ctx context.Context, req *DropIndexesReq) (*DropIndexesResp, error)
}

type Field = structor.Field

type ddl struct {
	db dorm.Dept
}

func NewDDL(ctx context.Context, opts ...Option) DDLService {
	d := &ddl{}

	for _, opt := range opts {
		opt(d)
	}
	return d
}

type Option func(*ddl)

func WithDB(db dorm.Dept) Option {
	return func(d *ddl) {
		d.db = db
	}
}

type ExecuteReq struct {
	Option string
	Table  string
	Fields []*Field
}

type ExecuteResp struct {
	Table string
}

func (d *ddl) Execute(ctx context.Context, req *ExecuteReq) (*ExecuteResp, error) {
	c, err := convert(req.Option, req.Table, req.Fields...)
	if err != nil {
		return nil, err
	}

	sc := d.db.Build(req.Table, c)
	err = sc.Exec(ctx)
	if err != nil {
		return nil, err
	}
	return &ExecuteResp{
		Table: req.Table,
	}, nil
}

type IndexReq struct {
	Option string
	Table  string
	Fields []*Field
}

type IndexResp struct {
	Index string
}

func (d *ddl) Index(ctx context.Context, req *IndexReq) (*IndexResp, error) {
	c, err := convert(req.Option, req.Table, req.Fields...)
	if err != nil {
		return nil, err
	}

	sc := d.db.Build(req.Table, c)
	indexName := genIndexName(req.Option, req.Fields...)

	err = sc.Index(ctx, indexName)
	if err != nil {
		return nil, err
	}

	return &IndexResp{
		Index: indexName,
	}, nil
}

type DropIndexesReq struct {
	Option string
	Table  string
	Fields []*Field
}

type DropIndexesResp struct {
}

func (d *ddl) DropIndexes(ctx context.Context, req *DropIndexesReq) (*DropIndexesResp, error) {
	c, err := convert(req.Option, req.Table, req.Fields...)
	if err != nil {
		return nil, err
	}

	sc := d.db.Build(req.Table, c)

	err = sc.DropIndexes(ctx)
	if err != nil {
		return nil, err
	}
	return &DropIndexesResp{}, nil

}

func convert(op string, table string, values ...*Field) (structor.Constructor, error) {
	return structor.GetDdlConstructor(op, table, values...)
}

func genIndexName(op string, values ...*Field) string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("%s_", op))
	for _, f := range values {
		builder.WriteString(f.Title[:1])
	}
	return builder.String()
}
