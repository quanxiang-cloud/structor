package clause

import "github.com/quanxiang-cloud/structor/pkg/errors"

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
		return nil, errors.ErrNoExpression
	}

	expr, ok := exprs[op]
	if !ok {
		return nil, errors.ErrNoExpression
	}

	expression := expr()
	expression.Set(column, values...)
	return expression, nil
}
