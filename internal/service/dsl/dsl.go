package dsl

import (
	"reflect"
)

type Bool map[string][]Query

type Aggs map[string]Agg

type Agg map[string]struct {
	Field string `json:"field,omitempty"`
}

type Query map[string]Field

type DSL struct {
	Query Query `json:"query,omitempty"`
	Bool  Bool  `json:"bool,omitempty"`
	Aggs  Aggs  `json:"aggs,omitempty"`
}

type Field map[string]Value

type Value interface{}

func Disintegration(v Value) []interface{} {
	switch reflect.TypeOf(v).Kind() {
	case reflect.Ptr:
		return Disintegration(reflect.ValueOf(v).Elem())
	case reflect.Array, reflect.Slice:
		valueOfValue := reflect.ValueOf(v)
		l := valueOfValue.Len()
		result := make([]interface{}, 0, l)
		for index := 0; index < l; index++ {
			if valueOfValue.Index(index).CanInterface() {
				result = append(result, valueOfValue.Index(index).Interface())
			}
		}
		return result
	default:
		return []interface{}{v}
	}
}

func NewValue(values ...interface{}) Value {
	return values
}
