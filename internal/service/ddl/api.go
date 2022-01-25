package ddl

import (
	"context"

	"github.com/quanxiang-cloud/structor/internal/dorm"
	"github.com/quanxiang-cloud/structor/internal/dorm/structor"
)

type DDLService interface {
	Execute(context.Context, *ExecuteReq) (*ExecuteResp, error)
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
	Fields []Field
}

type ExecuteResp struct {
	Table string
}

func (d *ddl) Execute(ctx context.Context, req *ExecuteReq) (*ExecuteResp, error) {
	c, err := convert(req.Option, req.Table, req.Fields...)
	if err != nil {
		return nil, err
	}

	sc := d.db.Build(c)
	err = sc.Exec(ctx)
	if err != nil {
		return nil, err
	}
	return &ExecuteResp{
		Table: req.Table,
	}, nil
}

func convert(op string, table string, values ...Field) (structor.Constructor, error) {
	return structor.GetDdlConstructor(op, table, values...)
}
