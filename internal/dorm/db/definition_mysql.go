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

var defaultID = structor.Fields{
	{
		Title:   "_id",
		Type:    "string",
		Comment: "unique id",
	},
}

func (c *Create) Build(builder structor.Builder) {
	builder.WriteRaw(fmt.Sprintf(" CREATE TABLE `%s` ( ", c.Table))
	fields := defaultID.Convert(mySQLDialector)
	for index, field := range fields {
		builder.WriteRaw(field)
		if index != len(fields)-1 {
			builder.WriteRaw(",")
		}
	}
	builder.WriteRaw(fmt.Sprintf(") ENGINE=%s DEFAULT CHARSET=%s COLLATE=%s;", engine, charset, collate))
}

type Add struct {
	structor.Add
}

func add() structor.Constructor {
	return &Add{}
}

func (a *Add) Build(builder structor.Builder) {
	builder.WriteRaw(fmt.Sprintf(" ALTER TABLE `%s` ", a.Table))
	fields := a.Fields.Convert(mySQLDialector)
	for index, field := range fields {
		builder.WriteRaw(fmt.Sprintf(" ADD COLUMN %s ", field))
		if index != len(fields)-1 {
			builder.WriteRaw(",")
		}
	}
	builder.WriteRaw(";")
}

type Modify struct {
	structor.Modify
}

func modify() structor.Constructor {
	return &Modify{}
}

func (m *Modify) Build(builder structor.Builder) {
	builder.WriteRaw(fmt.Sprintf(" ALTER TABLE `%s` ", m.Table))
	fields := m.Fields.Convert(mySQLDialector)
	for index, field := range fields {
		builder.WriteRaw(fmt.Sprintf(" MODIFY COLUMN %s ", field))
		if index != len(fields)-1 {
			builder.WriteRaw(",")
		}
	}
	builder.WriteRaw(";")
}

type Primary struct {
	structor.Primary
}

func primary() structor.Constructor {
	return &Primary{}
}

func (p *Primary) Build(builder structor.Builder) {
	builder.WriteRaw(fmt.Sprintf(" ALTER TABLE `%s` ADD PRIMARY KEY( ", p.Table))
	builder.WriteRaw(defaultID.ConvertIndex())
	builder.WriteRaw(" ); ")
}

type Index struct {
	structor.Index
}

func index() structor.Constructor {
	return &Index{}
}

func (i *Index) Build(builder structor.Builder) {
	builder.WriteRaw(fmt.Sprintf(" ALTER TABLE `%s` ADD INDEX `%s` (", i.Table, i.Name))
	builder.WriteRaw(i.Fields.ConvertIndex())
	builder.WriteRaw(");")
}

type Unique struct {
	structor.Unique
}

func unique() structor.Constructor {
	return &Unique{}
}

func (u *Unique) Build(builder structor.Builder) {
	builder.WriteRaw(fmt.Sprintf(" ALTER TABLE `%s` ADD UNIQUE `%s` ( ", u.Table, u.Name))
	builder.WriteRaw(u.Fields.ConvertIndex())
	builder.WriteRaw(");")
}

type DropIndex struct {
	structor.DropIndex
}

func dropIndex() structor.Constructor {
	return &DropIndex{}
}

func (d *DropIndex) Build(builder structor.Builder) {
	builder.WriteRaw(fmt.Sprintf(" ALTER TABLE `%s` DROP INDEX `%s` ;", d.Table, d.Name))
}
