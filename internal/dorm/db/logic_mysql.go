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
	// vars := make([]string, 0, len(m.Vars))
	// for _, val := range m.Vars {
	// 	val.(clause.Expression).Build(builder)
	// 	each, ok := builder.GetVar().(string)
	// 	if ok {
	// 		vars = append(vars, each)
	// 	}
	// }
	// expc := strings.Join(vars, "and")
	// builder.AddVar(expc)
}

type SHOULD struct {
	clause.SHOULD
}

func should() clause.Expression {
	return &SHOULD{}
}

func (s *SHOULD) Build(builder clause.Builder) {
	// vars := make([]string, 0, len(s.Vars))
	// for _, val := range s.Vars {
	// 	val.(clause.Expression).Build(builder)
	// 	each, ok := builder.GetVar().(string)
	// 	if ok {
	// 		vars = append(vars, each)
	// 	}
	// }
	// expc := strings.Join(vars, "or")
	// builder.AddVar(expc)
}

type MUSTNOT struct {
	clause.MUSTNOT
}

func mustNot() clause.Expression {
	return &MUSTNOT{}
}

func (m *MUSTNOT) Build(builder clause.Builder) {
	// TODO:
}

type RANGE struct {
	clause.RANGE
}

func range1() clause.Expression {
	return &RANGE{}
}

func (r *RANGE) Build(builder clause.Builder) {
	// TODO:
}
