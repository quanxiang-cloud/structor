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

// Avg Avg
type Avg struct {
	clause.Avg
}

func avg() clause.Expression {
	return &Avg{}
}

func (avg *Avg) Build(builder clause.Builder) {
	value, ok := avg.Value.(string)
	if ok {
		builder.AddAggVar(avg.Alias, bson.M{"$avg": "$" + value})
	}
}

// Min Min
type Min struct {
	clause.Min
}

func min() clause.Expression {
	return &Min{}
}

func (min *Min) Build(builder clause.Builder) {
	value, ok := min.Value.(string)
	if ok {
		builder.AddAggVar(min.Alias, bson.M{"$min": "$" + value})
	}
}

// Max Max
type Max struct {
	clause.Max
}

func max() clause.Expression {
	return &Max{}
}

func (max *Max) Build(builder clause.Builder) {
	value, ok := max.Value.(string)
	if ok {
		builder.AddAggVar(max.Alias, bson.M{"$max": "$" + value})
	}
}

// Max Max
type Count struct {
	clause.Count
}

func count() clause.Expression {
	return &Count{}
}

func (count *Count) Build(builder clause.Builder) {
	builder.AddAggVar("$count", count.Alias)
}
