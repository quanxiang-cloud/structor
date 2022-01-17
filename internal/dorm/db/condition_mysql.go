//+build mysql

package db

import (
	"fmt"

	"github.com/quanxiang-cloud/structor/internal/dorm/clause"
)

type Terms struct {
	clause.Terms
}

func terms() clause.Expression {
	return &Terms{}
}

func (t *Terms) Build(builder clause.Builder) {
	// TODO:
	builder.WriteString(fmt.Sprintf(" %s in ? ", t.Column))
	builder.AddVar(t.Values)
}

type Match struct {
	clause.Match
}

func match() clause.Expression {
	return &Match{}
}

func (m *Match) Build(builder clause.Builder) {
	builder.WriteString(fmt.Sprintf(" %s like ? ", m.Column))
	builder.AddVar(fmt.Sprintf("%%%v%%", m.Values))
}

type Term struct {
	clause.Term
}

func term() clause.Expression {
	return &Term{}
}

func (t *Term) Build(builder clause.Builder) {
	builder.WriteString(fmt.Sprintf(" %s = ? ", t.Column))
	builder.AddVar(t.Values)
}

type LT struct {
	clause.LT
}

func lt() clause.Expression {
	return &LT{}
}

func (lt *LT) Build(builder clause.Builder) {
	builder.WriteString(fmt.Sprintf(" %s < ? ", lt.Column))
	builder.AddVar(lt.Values)
}

type LTE struct {
	clause.LTE
}

func lte() clause.Expression {
	return &LTE{}
}

func (lte *LTE) Build(builder clause.Builder) {
	builder.WriteString(fmt.Sprintf(" %s <= ? ", lte.Column))
	builder.AddVar(lte.Values)
}

type GT struct {
	clause.GT
}

func gt() clause.Expression {
	return &GT{}
}

func (gt *GT) Build(builder clause.Builder) {
	builder.WriteString(fmt.Sprintf(" %s > ? ", gt.Column))
	builder.AddVar(gt.Values)
}

type GTE struct {
	clause.GTE
}

func gte() clause.Expression {
	return &GTE{}
}

func (gte *GTE) Build(builder clause.Builder) {
	builder.WriteString(fmt.Sprintf(" %s >= ? ", gte.Column))
	builder.AddVar(gte.Values)
}
