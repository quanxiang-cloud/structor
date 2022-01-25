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
	if value, ok := s.Value.(string); ok {
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
	if value, ok := a.Value.(string); ok {
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
	if value, ok := m.Value.(string); ok {
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
	if value, ok := m.Value.(string); ok {
		builder.AddAggVar(m.Alias, fmt.Sprintf("max(%s)", value))
	}
}
