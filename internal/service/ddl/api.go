package ddl

import (
	"context"

	"github.com/quanxiang-cloud/structor/internal/dorm"
	"github.com/quanxiang-cloud/structor/internal/dorm/structor"
)

type DDLService interface {
	Create(context.Context, *CreateReq) (*CreateResp, error)
	Add(context.Context, *AddReq) (*AddResp, error)
	Modify(context.Context, *ModifyReq) (*ModifyResp, error)
	Index(ctx context.Context, req *IndexReq) (*IndexResp, error)
	Unique(ctx context.Context, req *UniqueReq) (*UniqueResp, error)
	DropIndex(ctx context.Context, req *DropIndexReq) (*DropIndexResp, error)
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
	Table  string
	Fields []*Field
}

type ExecuteResp struct {
	Table string
}

type CreateReq = ExecuteReq

type CreateResp = ExecuteResp

func (d *ddl) Create(ctx context.Context, req *CreateReq) (*CreateResp, error) {
	c := structor.GetCreateExpr(req.Table, req.Fields)
	if err := d.db.Create(ctx, c); err != nil {
		return nil, err
	}

	c = structor.GetAddExpr(req.Table, req.Fields)
	if err := d.db.Add(ctx, c); err != nil {
		return nil, err
	}

	c = structor.GetPriMaryExpr(req.Table, req.Fields)
	if err := d.db.Primary(ctx, c); err != nil {
		return nil, err
	}

	return &CreateResp{
		Table: req.Table,
	}, nil
}

type AddReq = ExecuteReq

type AddResp = ExecuteResp

func (d *ddl) Add(ctx context.Context, req *AddReq) (*AddResp, error) {
	c := structor.GetAddExpr(req.Table, req.Fields)
	if err := d.db.Add(ctx, c); err != nil {
		return nil, err
	}
	return &AddResp{
		Table: req.Table,
	}, nil
}

type ModifyReq = ExecuteReq

type ModifyResp = ExecuteResp

func (d *ddl) Modify(ctx context.Context, req *ModifyReq) (*ModifyResp, error) {
	c := structor.GetModifyExpr(req.Table, req.Fields)
	if err := d.db.Modify(ctx, c); err != nil {
		return nil, err
	}
	return &ModifyResp{
		Table: req.Table,
	}, nil
}

// func (d *ddl) Execute(ctx context.Context, req *ExecuteReq) (*ExecuteResp, error) {
// 	c, err := convert(req.Option, req.Table, req.Fields...)
// 	if err != nil {
// 		return nil, err
// 	}

// 	sc := d.db.Build(req.Table, c)
// 	err = sc.Exec(ctx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &ExecuteResp{
// 		Table: req.Table,
// 	}, nil
// }

type IndexReq struct {
	ExecuteReq
	IndexName string
}

type IndexResp struct {
	IndexName string
}

func (d *ddl) Index(ctx context.Context, req *IndexReq) (*IndexResp, error) {
	c := structor.GetIndexExpr(req.Table, req.IndexName, req.Fields)

	if err := d.db.Index(ctx, c); err != nil {
		return nil, err
	}

	return &IndexResp{
		IndexName: req.IndexName,
	}, nil
}

type UniqueReq struct {
	ExecuteReq
	UniqueName string
}

type UniqueResp struct {
	UniqueName string
}

func (d *ddl) Unique(ctx context.Context, req *UniqueReq) (*UniqueResp, error) {
	c := structor.GetUniqueExpr(req.Table, req.UniqueName, req.Fields)

	if err := d.db.Unique(ctx, c); err != nil {
		return nil, err
	}

	return &UniqueResp{
		UniqueName: req.UniqueName,
	}, nil
}

type DropIndexReq struct {
	ExecuteReq
	IndexName string
}

type DropIndexResp struct {
	IndexName string
}

func (d *ddl) DropIndex(ctx context.Context, req *DropIndexReq) (*DropIndexResp, error) {
	c := structor.GetDropIndexExpr(req.Table, req.IndexName, req.Fields)

	if err := d.db.DropIndex(ctx, c); err != nil {
		return nil, err
	}

	return &DropIndexResp{
		IndexName: req.IndexName,
	}, nil
}

// func convert(op string, table string, values ...*Field) (structor.Constructor, error) {
// 	return structor.GetDdlConstructor(op, table, values...)
// }
