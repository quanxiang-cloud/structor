//+build mongo

package db

import (
	"github.com/quanxiang-cloud/structor/internal/dorm/structor"
)

type Create struct {
	structor.Create
}

func create() structor.Constructor {
	return &Create{}
}

func (c *Create) Build(table string, builder structor.Builder) {
	builder.Create(true, c.Column)
}

type Drop struct {
	structor.Drop
}

func drop() structor.Constructor {
	return &Drop{}
}

func (d *Drop) Build(table string, builder structor.Builder) {
	// do nothing
}

type Add struct {
	structor.Add
}

func add() structor.Constructor {
	return &Add{}
}

func (a *Add) Build(table string, builder structor.Builder) {
	// do nothing
}

type Del struct {
	structor.Del
}

func del() structor.Constructor {
	return &Del{}
}

func (d *Del) Build(table string, builder structor.Builder) {
	// do nothing
}

type Modify struct {
	structor.Modify
}

func modify() structor.Constructor {
	return &Modify{}
}

func (m *Modify) Build(table string, builder structor.Builder) {
	// do nothing
}

type Index struct {
	structor.Index
}

func index() structor.Constructor {
	return &Index{}
}

func (i *Index) Build(table string, builder structor.Builder) {
	for _, value := range i.Values {
		builder.WriteRaw(value.Title)
	}
}

type Unique struct {
	structor.Unique
}

func unique() structor.Constructor {
	return &Unique{}
}

func (u *Unique) Build(table string, builder structor.Builder) {
	for _, value := range u.Values {
		builder.WriteRaw(value.Title)
	}
	builder.Unique(u.IsUnique)
}

type DropIndexes struct {
	structor.DropIndexes
}

func dropIndexes() structor.Constructor {
	return &DropIndexes{}
}

func (d *DropIndexes) Build(table string, builder structor.Builder) {
	indexes := make([]string, 0, len(d.Values))
	for _, value := range d.Values {
		indexes = append(indexes, value.Title)
	}
	builder.IndexName(indexes)
}
