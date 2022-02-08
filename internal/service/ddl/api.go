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
	// Index(ctx context.Context, req *IndexReq) (*IndexResp, error)
	// DropIndexes(ctx context.Context, req *DropIndexesReq) (*DropIndexesResp, error)
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

// type IndexReq struct {
// 	Option string
// 	Table  string
// 	Fields []*Field
// }

// type IndexResp struct {
// 	Index string
// }

// func (d *ddl) Index(ctx context.Context, req *IndexReq) (*IndexResp, error) {
// 	c, err := convert(req.Option, req.Table, req.Fields...)
// 	if err != nil {
// 		return nil, err
// 	}

// 	sc := d.db.Build(req.Table, c)
// 	indexName := genIndexName(req.Option, req.Fields...)

// 	err = sc.Index(ctx, indexName)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &IndexResp{
// 		Index: indexName,
// 	}, nil
// }

// type DropIndexesReq struct {
// 	Option string
// 	Table  string
// 	Fields []*Field
// }

// type DropIndexesResp struct {
// }

// func (d *ddl) DropIndexes(ctx context.Context, req *DropIndexesReq) (*DropIndexesResp, error) {
// 	c, err := convert(req.Option, req.Table, req.Fields...)
// 	if err != nil {
// 		return nil, err
// 	}

// 	sc := d.db.Build(req.Table, c)

// 	err = sc.DropIndexes(ctx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &DropIndexesResp{}, nil

// }

// func convert(op string, table string, values ...*Field) (structor.Constructor, error) {
// 	return structor.GetDdlConstructor(op, table, values...)
// }

// func genIndexName(op string, values ...*Field) string {
// 	var builder strings.Builder
// 	builder.WriteString(fmt.Sprintf("%s_", op))
// 	for _, f := range values {
// 		builder.WriteString(f.Title[:1])
// 	}
// 	return builder.String()
// }
