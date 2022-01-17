package dorm

import (
	"context"

	"github.com/quanxiang-cloud/structor/internal/dorm/clause"
)

// Dorm dorm
type Dorm interface {
	Table(tablename string) Dorm
	Where(expr clause.Expression) Dorm
	Select(expr ...clause.Expression) Dorm
	Limit(limit int64) Dorm
	Offset(offset int64) Dorm
	Order(order ...string) Dorm
	FindOne(ctx context.Context) (map[string]interface{}, error)
	Find(ctx context.Context) ([]map[string]interface{}, error)
	Count(ctx context.Context) (int64, error)
	Insert(ctx context.Context, entities ...interface{}) (int64, error)
	Update(ctx context.Context, entity interface{}) (int64, error)
	Delete(ctx context.Context) (int64, error)
}
