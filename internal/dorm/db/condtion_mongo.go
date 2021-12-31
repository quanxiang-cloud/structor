//+build mongo

package db

import (
	"github.com/quanxiang-cloud/structor/internal/dorm/clause"
)

// IN Whether a value is within a set of values
type IN struct {
	clause.IN
}

func in() clause.Expression {
	return &IN{}
}

// Build build mongo bson
func (in *IN) Build(builder clause.Builder) {
	builder.WriteString("$in")
	builder.AddVar(in.Values)
	builder.WriteQuoted(in.Column)
}

// LIKE fuzzy
type LIKE struct {
	clause.LIKE
}

func like() clause.Expression {
	return &LIKE{}
}

// Build build mongo bson
func (like *LIKE) Build(builder clause.Builder) {
	builder.WriteString("$regex")
	builder.AddVar(like.Values)
	builder.WriteQuoted(like.Column)
}

// EQUAL equal
type EQUAL struct {
	clause.EQUAL
}

func equal() clause.Expression {
	return &EQUAL{}
}

// Build build mongo bson
func (equal *EQUAL) Build(builder clause.Builder) {
	builder.WriteQuoted(equal.Column)
	switch len(equal.Values) {
	case 0:
		equal.Values = append(equal.Values, "NULL")
	case 1:
		builder.AddVar(equal.Values[0])
	default:
		builder.AddVar(equal.Values)
	}
}

// LT less than
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
