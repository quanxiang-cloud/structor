//+build mongo

package db

import (
	"github.com/quanxiang-cloud/structor/internal/dorm/clause"
)

type Terms struct {
	clause.Terms
}

func terms() clause.Expression {
	return &Terms{}
}

// Build build mongo bson
func (t *Terms) Build(builder clause.Builder) {
	builder.WriteString("$in")
	builder.AddVar(t.Values)
	builder.WriteQuoted(t.Column)
}

type Match struct {
	clause.Match
}

func match() clause.Expression {
	return &Match{}
}

// Build build mongo bson
func (m *Match) Build(builder clause.Builder) {
	builder.WriteString("$regex")
	builder.AddVar(m.Values)
	builder.WriteQuoted(m.Column)
}

type Term struct {
	clause.Term
}

func term() clause.Expression {
	return &Term{}
}

// Build build mongo bson
func (t *Term) Build(builder clause.Builder) {
	builder.WriteQuoted(t.Column)
	switch len(t.Values) {
	case 0:
		t.Values = append(t.Values, "NULL")
	case 1:
		builder.AddVar(t.Values[0])
	default:
		builder.AddVar(t.Values)
	}
}

type LT struct {
	clause.LT
}

func lt() clause.Expression {
	return &LT{}
}

// Build build mongo bson
func (lt *LT) Build(builder clause.Builder) {
	builder.WriteString("$lt")
	builder.AddVar(lt.Values)
	builder.WriteQuoted(lt.Column)
}

// LTE less than or equal
type LTE struct {
	clause.LTE
}

func lte() clause.Expression {
	return &LTE{}
}

// Build build mongo bson
func (lte *LTE) Build(builder clause.Builder) {
	builder.WriteString("$lte")
	builder.AddVar(lte.Values)
	builder.WriteQuoted(lte.Column)
}

type GT struct {
	clause.GT
}

func gt() clause.Expression {
	return &GT{}
}

// Build build mongo bson
func (gt *GT) Build(builder clause.Builder) {
	builder.WriteString("$gt")
	builder.AddVar(gt.Values)
	builder.WriteQuoted(gt.Column)
}

type GTE struct {
	clause.GTE
}

func gte() clause.Expression {
	return &GTE{}
}

// Build build mongo bson
func (gte *GTE) Build(builder clause.Builder) {
	builder.WriteString("$gte")
	builder.AddVar(gte.Values)
	builder.WriteQuoted(gte.Column)
}
