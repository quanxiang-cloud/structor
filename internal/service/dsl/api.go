package dsl

import (
	"context"
	"fmt"

	"github.com/quanxiang-cloud/structor/internal/dorm/clause"
	"github.com/quanxiang-cloud/structor/internal/dorm/db"
)

type DSLService interface {
	FindOne(ctx context.Context, req *FindOneReq) (*FindOneResp, error)
}

type dsl struct {
	db *db.Dorm

	clause *clause.Clause
}

func New(ctx context.Context, opts ...Option) DSLService {
	d := &dsl{
		clause: clause.New(),
	}

	for _, opt := range opts {
		opt(d)
	}
	return d
}

type Option func(*dsl)

func WithDB(db *db.Dorm) Option {
	return func(d *dsl) {
		d.db = db
	}
}

type FindOneReq struct {
	TableName string
	DSL       DSL
}

type FindOneResp struct {
	Data interface{}
}

func (d *dsl) FindOne(ctx context.Context, req *FindOneReq) (*FindOneResp, error) {
	where, aggs, err := d.convert(req.DSL)
	if err != nil {
		return &FindOneResp{}, err
	}

	ql := d.db.Table(req.TableName)
	if where != nil {
		ql = ql.Where(where)
	}

	if aggs != nil {
		ql = ql.Select(aggs...)
	}

	data, err := ql.FindOne(ctx)
	if err != nil {
		return &FindOneResp{}, err
	}

	return &FindOneResp{
		Data: data,
	}, nil
}

func (d *dsl) convert(dsl DSL) (where clause.Expression, aggs []clause.Expression, err error) {
	aggs = make([]clause.Expression, 0)
	for alias, agg := range dsl.Aggs {
		for op, field := range agg {
			expr, err := d.clause.GetExpression(op, alias, field.Field)
			if err != nil {
				return nil, nil, err
			}
			aggs = append(aggs, expr)
			break
		}
	}

	for op, queries := range dsl.Bool {
		subExpr := make([]interface{}, 0, len(queries))
		for _, query := range queries {
			expr, err := d.query(query)
			if err != nil {
				return nil, nil, err
			}
			if expr != nil {
				subExpr = append(subExpr, expr)
			}
		}

		where, err = d.clause.GetExpression(op, "", subExpr...)
		if err != nil {
			return nil, nil, err
		}
		return
	}

	where, err = d.query(dsl.Query)
	if err != nil {
		return nil, nil, err
	}
	return
}

func (d *dsl) query(query Query) (clause.Expression, error) {
	if len(query) == 0 {
		return nil, nil
	}

	for op, field := range query {
		for name, value := range field {
			return d.clause.GetExpression(op, name, Disintegration(value)...)
		}
	}
	return nil, fmt.Errorf("query must have one")
}
