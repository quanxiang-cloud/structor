package clause

import (
	"errors"
)

var (
	// ErrNoExpression no expression
	ErrNoExpression = errors.New("no expression like this")
)

type Expr func() Expression

var dmlExprs map[string]Expr

func SetDmlExpressions(es map[string]Expr) {
	dmlExprs = es
}

func getDmlExpressions() map[string]Expr {
	return dmlExprs
}

// GetDmlExpression get dml expression with op
func GetDmlExpression(op string, column string, values ...interface{}) (Expression, error) {
	exprs := getDmlExpressions()
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

var ddlExprs map[string]Expr
