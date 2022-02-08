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

func (fs Fields) Convert(dialector Dialector) []string {
	fields := make([]string, 0, len(fs))
	dialectMgr.Register(dialector)
	for _, f := range fs {
		builder := bytes.Buffer{}
		ds := dialectMgr.Transform(f)
		builder.WriteString(fmt.Sprintf(" `%s` %s ", f.Title, ds(f)))
		if f.NotNull {
			builder.WriteString(" NOT NULL ")
		}
		if f.Comment != "" {
			builder.WriteString(fmt.Sprintf(" COMMENT '%s' ", f.Comment))
		}
		fields = append(fields, builder.String())
	}
	return fields
}

func (fs Fields) ConvertIndex() string {
	builder := bytes.Buffer{}
	for index, f := range fs {
		builder.WriteString(fmt.Sprintf(" `%s` ", f.Title))
		if index != len(fs)-1 {
			builder.WriteString(",")
		}
	}
	return builder.String()
}

func (fs Fields) GenIndexName(indexType string) string {
	var builder bytes.Buffer
	builder.WriteString(fmt.Sprintf("%s_", indexType))
	for _, f := range fs {
		builder.WriteString(f.Title[:1])
	}

	return builder.String()
}

type Create struct {
	Table  string
	Fields Fields
}

func (c *Create) GetTag() string {
	return "create"
}

func (c *Create) Set(Table string, Fields ...*Field) {
	c.Table = Table
	c.Fields = Fields
}

type Drop struct {
	Table string
}

func (d *Drop) GetTag() string {
	return "drop"
}

func (d *Drop) Set(Table string, Fields ...*Field) {
	d.Table = Table
}

type Add struct {
	Table  string
	Fields Fields
}

func (a *Add) GetTag() string {
	return "add"
}

func (a *Add) Set(Table string, Fields ...*Field) {
	a.Table = Table
	a.Fields = Fields
}

type Del struct {
	Table  string
	Fields Fields
}

func (d *Del) GetTag() string {
	return "del"
}

func (d *Del) Set(Table string, Fields ...*Field) {
	d.Table = Table
	d.Fields = Fields
}

type Modify struct {
	Table  string
	Fields Fields
}

func (u *Modify) GetTag() string {
	return "modify"
}

func (u *Modify) Set(Table string, Fields ...*Field) {
	u.Table = Table
	u.Fields = Fields
}

type Primary struct {
	Table  string
	Fields Fields
}

func (p *Primary) GetTag() string {
	return "primary"
}

func (p *Primary) Set(Table string, fields ...*Field) {
	p.Table = Table
	p.Fields = fields
}

type Index struct {
	Table    string
	Fields   Fields
	IsUnique bool
}

func (i *Index) GetTag() string {
	return "index"
}

func (i *Index) Set(Table string, Fields ...*Field) {
	i.Table = Table
	i.Fields = Fields
}

type Unique struct {
	Table    string
	Fields   Fields
	IsUnique bool
}

func (u *Unique) GetTag() string {
	return "unique"
}

func (u *Unique) Set(Table string, Fields ...*Field) {
	u.Table = Table
	u.Fields = Fields
	u.IsUnique = true
}

const DropIndexesOpt = "drop_indexes"

type DropIndexes struct {
	Table  string
	Fields Fields
}

func (d *DropIndexes) GetTag() string {
	return "drop_indexes"
}

func (d *DropIndexes) Set(Table string, Fields ...*Field) {
	d.Table = Table
	d.Fields = Fields
}
