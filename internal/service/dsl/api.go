package dsl

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/quanxiang-cloud/structor/internal/dorm/clause"
	"github.com/quanxiang-cloud/structor/internal/dorm/db"
)

type DSLService interface {
	FindOne(ctx context.Context, req *FindOneReq, apiOpts ...APIOption) (*FindOneResp, error)
	Find(ctx context.Context, req *FindReq, apiOpts ...APIOption) (*FindResp, error)
	Count(ctx context.Context, req *CountReq, apiOpts ...APIOption) (*CountResp, error)
	Insert(ctx context.Context, req *InsertReq, apiOpts ...APIOption) (*InsertResp, error)
	Update(ctx context.Context, req *UpdateReq, apiOpts ...APIOption) (*UpdateResp, error)
	Delete(ctx context.Context, req *DeleteReq, apiOpts ...APIOption) (*DeleteResp, error)
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

type APIOption func(...interface{}) error

func WithInsert(suffix string) APIOption {
	return func(entities ...interface{}) error {
		for _, entity := range entities {
			ek := reflect.TypeOf(entity).Kind()
			if ek != reflect.Map {
				return fmt.Errorf("invalid entity type: %s", ek)
			}
			iter := reflect.ValueOf(entity).MapRange()
			for iter.Next() {
				if !iter.Value().CanInterface() {
					continue
				}

				if !strings.HasSuffix(iter.Key().String(), suffix) {
					continue
				}

				buf, err := json.Marshal(iter.Value().Interface())
				if err != nil {
					return err
				}
				reflect.ValueOf(entity).SetMapIndex(iter.Key(), reflect.ValueOf(string(buf)))

			}
		}
		return nil
	}
}

func WithSearch(suffix string) APIOption {
	return func(datas ...interface{}) error {
		for _, data := range datas {
			dk := reflect.TypeOf(data).Kind()
			if dk != reflect.Map {
				return fmt.Errorf("invalid entity type: %s", dk)
			}
			iter := reflect.ValueOf(data).MapRange()
			for iter.Next() {
				if !iter.Value().CanInterface() {
					continue
				}

				if !strings.HasSuffix(iter.Key().String(), suffix) {
					continue
				}

				var value interface{}
				err := json.Unmarshal([]byte(iter.Value().Elem().String()), &value)
				if err != nil {
					return err
				}
				reflect.ValueOf(data).SetMapIndex(iter.Key(), reflect.ValueOf(value))

			}
		}
		return nil
	}
}

type FindOneReq struct {
	TableName string
	DSL       DSL
}

type FindOneResp struct {
	Data interface{}
}

func (d *dsl) FindOne(ctx context.Context, req *FindOneReq, apiOpts ...APIOption) (*FindOneResp, error) {
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

	for _, opt := range apiOpts {
		opt(data)
	}

	return &FindOneResp{
		Data: data,
	}, nil
}

type FindReq struct {
	TableName string
	Page      int64
	Size      int64
	Sort      []string
	DSL       DSL
}

type FindResp struct {
	Data []interface{}
}

func (d *dsl) Find(ctx context.Context, req *FindReq, apiOpts ...APIOption) (*FindResp, error) {
	where, aggs, err := d.convert(req.DSL)
	if err != nil {
		return &FindResp{}, err
	}

	ql := d.db.Table(req.TableName)
	if where != nil {
		ql = ql.Where(where)
	}

	if aggs != nil {
		ql = ql.Select(aggs...)
	}

	ql = ql.Offset((req.Page - 1) * req.Size).Limit(req.Size)
	ql = ql.Order(req.Sort...)

	data, err := ql.Find(ctx)
	if err != nil {
		return &FindResp{}, err
	}

	dl := make([]interface{}, 0, len(data))
	for _, v := range data {
		dl = append(dl, v)
	}

	for _, opt := range apiOpts {
		opt(dl...)
	}

	return &FindResp{
		Data: dl,
	}, nil
}

type CountReq struct {
	TableName string
	DSL       DSL
}

type CountResp struct {
	Data int64
}

func (d *dsl) Count(ctx context.Context, req *CountReq, apiOpts ...APIOption) (*CountResp, error) {
	where, aggs, err := d.convert(req.DSL)
	if err != nil {
		return &CountResp{}, err
	}

	ql := d.db.Table(req.TableName)
	if where != nil {
		ql = ql.Where(where)
	}

	if aggs != nil {
		ql = ql.Select(aggs...)
	}

	data, err := ql.Count(ctx)
	if err != nil {
		return &CountResp{}, err
	}

	return &CountResp{
		Data: data,
	}, nil
}

type UpdateReq struct {
	TableName string
	DSL       DSL
	Entity    interface{}
}

type UpdateResp struct {
	Count int64
}

func (d *dsl) Update(ctx context.Context, req *UpdateReq, apiOpts ...APIOption) (*UpdateResp, error) {
	where, _, err := d.convert(req.DSL)
	if err != nil {
		return &UpdateResp{}, err
	}

	ql := d.db.Table(req.TableName)
	if where != nil {
		ql = ql.Where(where)
	}

	count, err := ql.Update(ctx, req.Entity)
	return &UpdateResp{
		Count: count,
	}, err
}

type InsertReq struct {
	TableName string
	Entities  []interface{}
}

type InsertResp struct {
	Count int64
}

func (d *dsl) Insert(ctx context.Context, req *InsertReq, apiOpts ...APIOption) (*InsertResp, error) {
	ql := d.db.Table(req.TableName)

	for _, opt := range apiOpts {
		opt(req.Entities...)
	}

	count, err := ql.Insert(ctx, req.Entities...)
	if err != nil {
		return &InsertResp{}, err
	}

	return &InsertResp{
		Count: count,
	}, nil
}

type DeleteReq struct {
	TableName string
	DSL       DSL
}

type DeleteResp struct {
	Count int64
}

func (d *dsl) Delete(ctx context.Context, req *DeleteReq, apiOpts ...APIOption) (*DeleteResp, error) {
	where, _, err := d.convert(req.DSL)
	if err != nil {
		return &DeleteResp{}, err
	}

	ql := d.db.Table(req.TableName)
	if where != nil {
		ql = ql.Where(where)
	}

	count, err := ql.Delete(ctx)
	if err != nil {
		return nil, err
	}

	return &DeleteResp{
		Count: count,
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
