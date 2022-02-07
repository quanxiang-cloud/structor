//+build mysql

package db

import (
	"fmt"

	"github.com/quanxiang-cloud/structor/internal/dorm/structor"
)

var mySQLDialector = &MySQLDialector{}

type MySQLDialector struct {
}

var (
	SMALLINT_MAX  int64 = 1 << 7  // 8 bit
	MEDIUMINT_MAX int64 = 1 << 15 // 16 bit
	INT_MAX       int64 = 1 << 23 // 24 bit
	BIGINT_MAX    int64 = 1 << 31 // 32 bit
)

func (m *MySQLDialector) Int(f *structor.Field) string {
	if f.Max != 0 {
		switch {
		case f.Max < SMALLINT_MAX:
			return "TINYINT"
		case f.Max < MEDIUMINT_MAX:
			return "SMALLINT"
		case f.Max < INT_MAX:
			return "MEDIUMINT"
		case f.Max < BIGINT_MAX:
			return "INT"
		default:
			return "BIGINT"
		}
	}
	return "INT"
}

func (m *MySQLDialector) Float(f *structor.Field) string {
	return "DOUBLE"
}

func (m *MySQLDialector) Bool(f *structor.Field) string {
	return "BOOL"
}

func (m *MySQLDialector) String(f *structor.Field) string {
	if f.Max > 255 {
		return "TEXT"
	}
	if f.Max == 0 {
		f.Max = 64
	}
	return fmt.Sprintf("VARCHAR(%d)", f.Max)
}

func (m *MySQLDialector) Datetime(f *structor.Field) string {
	return "VARCHAR(64)"
}

func (m *MySQLDialector) Object(f *structor.Field) string {
	return "TEXT"
}
