//+build mongo

package db

import "github.com/quanxiang-cloud/structor/internal/dorm/clause"

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
