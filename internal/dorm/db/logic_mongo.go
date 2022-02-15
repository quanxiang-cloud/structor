//+build mongo

package db

import (
	"encoding/json"
	"fmt"

	"github.com/quanxiang-cloud/structor/internal/dorm/clause"
	"github.com/quanxiang-cloud/structor/internal/service/dsl"
)

type MUST struct {
	clause.MUST
}

func must() clause.Expression {
	return &MUST{}
}

// Build build mongo bson
func (m *MUST) Build(builder clause.Builder) {
	vars := make([]interface{}, 0, len(m.Vars))
	for _, value := range m.Vars {
		value.(clause.Expression).Build(builder)
		vars = append(vars, builder.GetVar())
	}
	builder.WriteString("$and")
	builder.AddVar(vars)
}

type SHOULD struct {
	clause.SHOULD
}

func should() clause.Expression {
	return &SHOULD{}
}

// Build build mongo bson
func (s *SHOULD) Build(builder clause.Builder) {
	vars := make([]interface{}, 0, len(s.Vars))
	for _, value := range s.Vars {
		value.(clause.Expression).Build(builder)
		vars = append(vars, builder.GetVar())
	}
	builder.WriteString("$or")
	builder.AddVar(vars)
}

type MUSTNOT struct {
	clause.MUSTNOT
}

func mustNot() clause.Expression {
	return &MUSTNOT{}
}

// Build build mongo bson
func (m *MUSTNOT) Build(builder clause.Builder) {
	vars := make([]interface{}, 0, len(m.Vars))
	for _, value := range m.Vars {
		value.(clause.Expression).Build(builder)
		vars = append(vars, builder.GetVar())
	}
	builder.WriteString("$nor")
	builder.AddVar(vars)
}

type RANGE struct {
	clause.RANGE
}

func range1() clause.Expression {
	return &RANGE{}
}

func (r *RANGE) Build(builder clause.Builder) {
	if len(r.Vars) != 0 {
		if val, ok := r.Vars[0].(map[string]interface{}); ok {
			vars := make([]interface{}, 0, len(r.Vars))
			for k, v := range val {
				subExpr, err := clause.GetDmlExpression(k, r.Column, v)
				if err != nil {
					continue
				}
				subExpr.Build(builder)
				vars = append(vars, builder.GetVar())
			}
			builder.WriteString("$and")
			builder.AddVar(vars)
		}
	}
}

type Bool struct {
	clause.Bool
}

func bool1() clause.Expression {
	return &Bool{}
}

func (b *Bool) Build(builder clause.Builder) {
	var queryExpr = func(query dsl.Query) (clause.Expression, error) {
		if len(query) == 0 {
			return nil, nil
		}

		for op, field := range query {
			for name, value := range field {
				return clause.GetDmlExpression(op, name, dsl.Disintegration(value)...)
			}
		}
		return nil, fmt.Errorf("query must have one")
	}

	var (
		queries []dsl.Query
		subExpr = make([]interface{}, 0, len(b.Vars))
	)

	querBytes, err := json.Marshal(b.Vars)
	if err != nil {
		return
	}
	err = json.Unmarshal(querBytes, &queries)
	if err != nil {
		return
	}

	for _, query := range queries {
		expr, err := queryExpr(query)
		if err != nil {
			continue
		}
		if expr != nil {
			subExpr = append(subExpr, expr)
		}
	}

	where, err := clause.GetDmlExpression(b.Column, "", subExpr...)
	if err != nil {
		return
	}

	where.Build(builder)
}
