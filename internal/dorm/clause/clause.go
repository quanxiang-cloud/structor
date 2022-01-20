package clause

import (
	"errors"
)

// Writer write
type Writer interface {
	WriteByte(byte) error
	WriteString(string) (int, error)
}

// QueryBuilder db condition builder
type QueryBuilder interface {
	Writer
	WriteQuoted(field string)
	AddVar(value interface{})
	GetVar() interface{}
}

// AggBuilder AggBuilder
type AggBuilder interface {
	WriteQuotedAgg(field string)
	AddAggVar(key string, value interface{})
}

//Builder Builder
type Builder interface {
	QueryBuilder
	AggBuilder
}

var (
	// ErrNoExpression no expression
	ErrNoExpression = errors.New("no expression like this")
)

type Expr func() Expression

var expressions map[string]Expr

var ddlExprs map[string]Expr

func SetExpressions(es map[string]Expr) {
	expressions = es
}

func getExpressions() map[string]Expr {
	return expressions
}

// Clause expressions set
type Clause struct {
}

// New new a clause
func New() *Clause {
	return &Clause{}
}

// GetExpression get expression with op
func GetExpression(op string, column string, values ...interface{}) (Expression, error) {
	exprs := getExpressions()
	if exprs == nil {
		return nil, ErrNoExpression
	}

	expr, ok := exprs[op]
	if !ok {
		return nil, ErrNoExpression
	}

	expression := expr()
	expression.Set(column, values...)
	return expression, nil
}
