package structor

import (
	"bytes"
	"fmt"
)

type Field struct {
	Title   string
	Type    string
	Max     int64
	Comment string
	NotNull bool
}

type Fields []*Field

func (fs Fields) Convert(dialector Dialector) string {
	dialectMgr.Register(dialector)
	builder := bytes.Buffer{}
	for index, f := range fs {
		ds := dialectMgr.Transform(f)
		builder.WriteString(fmt.Sprintf(" `%s` %s ", f.Title, ds(f)))
		if f.NotNull {
			builder.WriteString(" NOT NULL ")
		}
		if f.Comment != "" {
			builder.WriteString(fmt.Sprintf(" COMMENT '%s' ", f.Comment))
		}
		if index != len(fs)-1 {
			builder.WriteString(",")
		}
	}
	return builder.String()
}

type Create struct {
	Column string
	Values Fields
}

func (c *Create) GetTag() string {
	return "create"
}

func (c *Create) Set(column string, values ...*Field) {
	c.Column = column
	c.Values = values
}

type Drop struct {
	Column string
}

func (d *Drop) GetTag() string {
	return "drop"
}

func (d *Drop) Set(column string, values ...*Field) {
	d.Column = column
}

type Add struct {
	Column string
	Values Fields
}

func (a *Add) GetTag() string {
	return "add"
}

func (a *Add) Set(column string, values ...*Field) {
	a.Column = column
	a.Values = values
}

type Del struct {
	Column string
	Values Fields
}

func (d *Del) GetTag() string {
	return "del"
}

func (d *Del) Set(column string, values ...*Field) {
	d.Column = column
	d.Values = values
}

type Modify struct {
	Column string
	Values Fields
}

func (u *Modify) GetTag() string {
	return "modify"
}

func (u *Modify) Set(column string, values ...*Field) {
	u.Column = column
	u.Values = values
}
