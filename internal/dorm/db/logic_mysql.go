//+build mysql

package db

import (
	"github.com/quanxiang-cloud/structor/internal/dorm/clause"
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
	for i, val := range r.Vars {
		if i != 0 {
			builder.WriteQuoted(" and ")
		}
		val.(clause.Expression).Build(builder)
	}
}