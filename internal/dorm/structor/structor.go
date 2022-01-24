package structor

import (
	"github.com/quanxiang-cloud/structor/pkg/errors"
)

type Expr func() Constructor

var ddlExprs map[string]Expr

func SetDdlConstructors(es map[string]Expr) {
	ddlExprs = es
}

func getDdlConstructors() map[string]Expr {
	return ddlExprs
}

func GetDdlConstructor(op string, column string, values ...Field) (Constructor, error) {
	exprs := getDdlConstructors()
	if exprs == nil {
		return nil, errors.ErrNoConstructor
	}

	expr, ok := exprs[op]
	if !ok {
		return nil, errors.ErrNoConstructor
	}

	constructor := expr()
	constructor.Set(column, values...)
	return constructor, nil
}
