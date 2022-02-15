package errors

import "errors"

var (
	ErrNoExpression  = errors.New("no expression")
	ErrNoConstructor = errors.New("no constructor")
)
