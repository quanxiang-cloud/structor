package structor

import (
	"fmt"
	"sort"
)

type Type string

const (
	INT      Type = "int"
	FLOAT    Type = "float"
	BOOL     Type = "bool"
	STRING   Type = "string"
	DATETIME Type = "datetime"
	OBJECT   Type = "object"
)

var typeEnums = []Type{
	INT, FLOAT, BOOL, STRING, DATETIME, OBJECT,
}

func init() {
	sort.Slice(typeEnums, func(i, j int) bool {
		return typeEnums[i] < typeEnums[j]
	})
}

func CheckType(t string) error {
	low, high := 0, len(typeEnums)-1
	for low <= high {
		mid := (low + high) / 2
		switch {
		case string(typeEnums[mid]) < t:
			low = mid + 1
		case string(typeEnums[mid]) > t:
			high = mid - 1
		case string(typeEnums[mid]) == t:
			return nil
		}
	}
	return fmt.Errorf("invalid type (%s)", t)
}

type DialSpeaker func(*Field) string

type Dialector interface {
	Int(*Field) string
	Float(*Field) string
	Bool(*Field) string
	String(*Field) string
	Datetime(*Field) string
	Object(*Field) string
}

var dialectMgr = &Dialect{}

type Dialect struct {
	dialector Dialector
}

func (d *Dialect) Transform(f *Field) DialSpeaker {
	switch Type(f.Type) {
	case INT:
		return d.dialector.Int
	case FLOAT:
		return d.dialector.Float
	case BOOL:
		return d.dialector.Bool
	case STRING:
		return d.dialector.String
	case DATETIME:
		return d.dialector.Datetime
	case OBJECT:
		return d.dialector.Object
	default:
		return nil
	}
}

func (d *Dialect) Register(dialector Dialector) {
	d.dialector = dialector
}
