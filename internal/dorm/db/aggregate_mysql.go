//+build mysql

package db

import "github.com/quanxiang-cloud/structor/internal/dorm/clause"

type Sum struct {
	clause.Sum
}

func sum() clause.Expression {
	return &Sum{}
}

func (s *Sum) Build(builder clause.Builder) {
	// TODO:
}

type Avg struct {
	clause.Avg
}

func avg() clause.Expression {
	return &Avg{}
}

func (a *Avg) Build(builder clause.Builder) {
	// TODO:
}

type Min struct {
	clause.Min
}

func min() clause.Expression {
	return &Min{}
}

func (m *Min) Build(builder clause.Builder) {
	// TODO:
}

type Max struct {
	clause.Max
}

func max() clause.Expression {
	return &Max{}
}

func (m *Max) Build(builder clause.Builder) {
	// TODO:
}
