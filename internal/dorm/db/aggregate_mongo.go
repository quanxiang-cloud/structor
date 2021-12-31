//+build mongo

package db

import (
	"github.com/quanxiang-cloud/structor/internal/dorm/clause"
	"go.mongodb.org/mongo-driver/bson"
)

// Sum Sum
type Sum struct {
	clause.Sum
}

func sum() clause.Expression {
	return &Sum{}
}

func (sum *Sum) Build(builder clause.Builder) {
	value, ok := sum.Value.(string)
	if ok {
		builder.AddAggVar(sum.Alias, bson.M{"$sum": "$" + value})
	}
}
