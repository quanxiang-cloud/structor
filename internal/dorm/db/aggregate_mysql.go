//+build mysql

package db

import (
	"fmt"

	"github.com/quanxiang-cloud/structor/internal/dorm/clause"
)

type Sum struct {
	clause.Sum
}

func sum() clause.Expression {
	return &Sum{}
}

func (s *Sum) Build(builder clause.Builder) {
	value, ok := s.Value.(string)
	if ok {
		builder.AddAggVar(s.Alias, fmt.Sprintf("sum(%s)", value))
	}
}

type Avg struct {
	clause.Avg
}

func avg() clause.Expression {
	return &Avg{}
}

func (a *Avg) Build(builder clause.Builder) {
	value, ok := a.Value.(string)
	if ok {
		builder.AddAggVar(a.Alias, fmt.Sprintf("avg(%s)", value))
	}
}

type Min struct {
	clause.Min
}

func min() clause.Expression {
	return &Min{}
}

func (m *Min) Build(builder clause.Builder) {
	value, ok := m.Value.(string)
	if ok {
		builder.AddAggVar(m.Alias, fmt.Sprintf("min(%s)", value))
	}
}

type Max struct {
	clause.Max
}

func max() clause.Expression {
	return &Max{}
}

func (m *Max) Build(builder clause.Builder) {
	value, ok := m.Value.(string)
	if ok {
		builder.AddAggVar(m.Alias, fmt.Sprintf("max(%s)", value))
	}
}
