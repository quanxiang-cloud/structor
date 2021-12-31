//+build mongo

package db

import "github.com/quanxiang-cloud/structor/internal/dorm/clause"

type AND struct {
	clause.AND
}

func and() clause.Expression {
	return &AND{}
}

// Build build mongo bson
func (and *AND) Build(builder clause.Builder) {
	vars := make([]interface{}, 0, len(and.Vars))
	for _, value := range and.Vars {
		value.(clause.Expression).Build(builder)
		vars = append(vars, builder.(*MONGO).Vars)
	}
	builder.WriteString("$and")
	builder.AddVar(vars)
}

// OR or
type OR struct {
	clause.OR
}

func or() clause.Expression {
	return &OR{}
}

// GetTag get tag
func (or OR) GetTag() string {
	return "or"
}

// Build build mongo bson
func (or *OR) Build(builder clause.Builder) {
	vars := make([]interface{}, 0, len(or.Vars))
	for _, value := range or.Vars {
		value.(clause.Expression).Build(builder)
		vars = append(vars, builder.(*MONGO).Vars)
	}
	builder.WriteString("$or")
	builder.AddVar(vars)
}

// NOR nor
type NOR struct {
	clause.NOR
}

func nor() clause.Expression {
	return &NOR{}
}

// GetTag get tag
func (nor NOR) GetTag() string {
	return "nor"
}

// Build build mongo bson
func (nor *NOR) Build(builder clause.Builder) {
	vars := make([]interface{}, 0, len(nor.Vars))
	for _, value := range nor.Vars {
		value.(clause.Expression).Build(builder)
		vars = append(vars, builder.(*MONGO).Vars)
	}
	builder.WriteString("$nor")
	builder.AddVar(vars)
}
