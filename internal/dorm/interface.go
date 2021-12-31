package dorm

import (
	"context"

	"github.com/quanxiang-cloud/structor/internal/dorm/clause"
	"github.com/quanxiang-cloud/structor/internal/dorm/db"
)

// Dorm dorm
type Dorm interface {
	Table(tablename string) Dorm
	FindOne(ctx context.Context, expr clause.Expression) (clause.Data, error)
	Find(ctx context.Context, expr clause.Expression, findOpt clause.FindOptions) ([]clause.Data, error)
	Count(ctx context.Context, expr clause.Expression) (int64, error)
	Insert(ctx context.Context, entity ...interface{}) error
	Update(ctx context.Context, expr clause.Expression, entity interface{}) (int64, error)
	Delete(ctx context.Context, expr clause.Expression) (int64, error)
}

var _ Dorm = dorm{}

func New() (Dorm, error) {
	d, err := db.New()
	if err != nil {
		return nil, err
	}
	return &dorm{
		Dorm: d,
	}, nil
}

type dorm struct {
	*db.Dorm
}

func (d dorm) Table(tablename string) Dorm {
	dm := d.Dorm.Table(tablename)
	return &dorm{
		Dorm: dm,
	}
}

func NewBuilder() clause.Builder {
	return db.NewBuilder()
}
