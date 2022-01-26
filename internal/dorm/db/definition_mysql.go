//+build mysql

package db

import (
	"fmt"

	"github.com/quanxiang-cloud/structor/internal/dorm/structor"
)

type Create struct {
	structor.Create
}

func create() structor.Constructor {
	return &Create{}
}

func (c *Create) Build(table string, builder structor.Builder) {
	builder.WriteRaw(fmt.Sprintf(" CREATE TABLE `%s` ( ", c.Column))
	builder.WriteRaw(c.Values.Convert(mySQLDialector))
	builder.WriteRaw(fmt.Sprintf(") ENGINE=%s DEFAULT CHARSET=%s COLLATE=%s;", engine, charset, collate))
}

type Add struct {
	structor.Add
}

func add() structor.Constructor {
	return &Add{}
}

func (a *Add) Build(table string, builder structor.Builder) {
	builder.WriteRaw(fmt.Sprintf(" ALTER TABLE `%s` ADD COLUMN ", a.Column))
	builder.WriteRaw(a.Values.Convert(mySQLDialector))
	builder.WriteRaw(";")
}

type Modify struct {
	structor.Modify
}

func modify() structor.Constructor {
	return &Modify{}
}

func (m *Modify) Build(table string, builder structor.Builder) {
	builder.WriteRaw(fmt.Sprintf(" ALTER TABLE `%s` MODIFY COLUMN ", m.Column))
	builder.WriteRaw(m.Values.Convert(mySQLDialector))
	builder.WriteRaw(";")
}

type Index struct {
	structor.Index
}

func index() structor.Constructor {
	return &Index{}
}

func (i *Index) Build(table string, builder structor.Builder) {
	builder.WriteRaw(fmt.Sprintf(" ALTER TABLE `%s` ADD INDEX `%s` (", i.Column, i.Values.GenIndexName(i.GetTag())))
	builder.WriteRaw(i.Values.ConvertIndex())
	builder.WriteRaw(");")
}

type Unique struct {
	structor.Unique
}

func unique() structor.Constructor {
	return &Unique{}
}

func (u *Unique) Build(table string, builder structor.Builder) {
	builder.WriteRaw(fmt.Sprintf(" ALTER TABLE `%s` ADD UNIQUE `%s` ( ", u.Column, u.Values.GenIndexName(u.GetTag())))
	builder.WriteRaw(u.Values.ConvertIndex())
	builder.WriteRaw(");")
}
