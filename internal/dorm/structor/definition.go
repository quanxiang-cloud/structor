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

type Create struct {
	Table  string
	Fields Fields
}

func (c *Create) GetTag() string {
	return "create"
}

func (c *Create) Set(table string, index string, fields ...*Field) {
	c.Table = table
	c.Fields = fields
}

func (c *Create) GetTable() string {
	return c.Table
}

func (c *Create) GetIndex() string {
	return ""
}

type Drop struct {
	Table string
}

func (d *Drop) GetTag() string {
	return "drop"
}

func (d *Drop) Set(table string, index string, fields ...*Field) {
	d.Table = table
}

func (d *Drop) GetTable() string {
	return d.Table
}

func (d *Drop) GetIndex() string {
	return ""
}

type Add struct {
	Table  string
	Fields Fields
}

func (a *Add) GetTag() string {
	return "add"
}

func (a *Add) Set(table string, index string, fields ...*Field) {
	a.Table = table
	a.Fields = fields
}

func (a *Add) GetTable() string {
	return a.Table
}

func (a *Add) GetIndex() string {
	return ""
}

type Del struct {
	Table  string
	Fields Fields
}

func (d *Del) GetTag() string {
	return "del"
}

func (d *Del) Set(table string, index string, fields ...*Field) {
	d.Table = table
	d.Fields = fields
}

func (d *Del) GetTable() string {
	return d.Table
}

func (d *Del) GetIndex() string {
	return ""
}

type Modify struct {
	Table  string
	Fields Fields
}

func (m *Modify) GetTag() string {
	return "modify"
}

func (m *Modify) Set(table string, index string, fields ...*Field) {
	m.Table = table
	m.Fields = fields
}

func (m *Modify) GetTable() string {
	return m.Table
}

func (m *Modify) GetIndex() string {
	return ""
}

type Primary struct {
	Table  string
	Fields Fields
}

func (p *Primary) GetTag() string {
	return "primary"
}

func (p *Primary) Set(table string, index string, fields ...*Field) {
	p.Table = table
	p.Fields = fields
}

func (p *Primary) GetTable() string {
	return p.Table
}

func (p *Primary) GetIndex() string {
	return ""
}

type Index struct {
	Table  string
	Name   string
	Fields Fields
}

func (i *Index) GetTag() string {
	return "index"
}

func (i *Index) Set(table string, index string, fields ...*Field) {
	i.Table = table
	i.Name = index
	i.Fields = fields
}

func (i *Index) GetTable() string {
	return i.Table
}

func (i *Index) GetIndex() string {
	return i.Name
}

type Unique struct {
	Table  string
	Name   string
	Fields Fields
}

func (u *Unique) GetTag() string {
	return "unique"
}

func (u *Unique) Set(table string, index string, fields ...*Field) {
	u.Table = table
	u.Name = index
	u.Fields = fields
}

func (u *Unique) GetTable() string {
	return u.Table
}

func (u *Unique) GetIndex() string {
	return u.Name
}

type DropIndex struct {
	Table  string
	Name   string
	Fields Fields
}

func (d *DropIndex) GetTag() string {
	return "drop_index"
}

func (d *DropIndex) Set(table string, index string, fields ...*Field) {
	d.Table = table
	d.Name = index
	d.Fields = fields
}

func (d *DropIndex) GetTable() string {
	return d.Table
}

func (d *DropIndex) GetIndex() string {
	return d.Name
}
