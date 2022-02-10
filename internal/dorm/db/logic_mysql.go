//+build mysql

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

func (m *MUST) Build(builder clause.Builder) {
	for i, val := range m.Vars {
		if i != 0 {
			builder.WriteQuoted(" and ")
		}
		val.(clause.Expression).Build(builder)
	}
}

type SHOULD struct {
	clause.SHOULD
}

func should() clause.Expression {
	return &SHOULD{}
}

func (s *SHOULD) Build(builder clause.Builder) {
	for i, val := range s.Vars {
		if i != 0 {
			builder.WriteQuoted(" or ")
		}
		val.(clause.Expression).Build(builder)
	}
}

type MUSTNOT struct {
	clause.MUSTNOT
}

func mustNot() clause.Expression {
	return &MUSTNOT{}
}

func (m *MUSTNOT) Build(builder clause.Builder) {
	builder.WriteString(" not ( ")
	for i, val := range m.Vars {
		if i != 0 {
			builder.WriteQuoted(" and ")
		}
		val.(clause.Expression).Build(builder)
	}
	builder.WriteString(" ) ")
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
			count := 0
			for k, v := range val {
				if count != 0 {
					builder.WriteQuoted(" and ")
				}

				subExpr, err := clause.GetDmlExpression(k, r.Column, v)
				if err != nil {
					continue
				}
				subExpr.Build(builder)
				count++
			}
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

	builder.WriteQuoted(" ( ")
	where.Build(builder)
	builder.WriteQuoted(" ) ")
}
