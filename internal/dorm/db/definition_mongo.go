//+build mongo

package db

import (
	"github.com/quanxiang-cloud/structor/internal/dorm/structor"
	"go.mongodb.org/mongo-driver/bson"
)

type Create struct {
	structor.Create
}

func create() structor.Constructor {
	return &Create{}
}

const (
	jsonSchema = "$jsonSchema"
	id         = "_id"
)

var defaultID = bson.M{
	"bsonType": "object",
	"required": []string{"_id"},
	"properties": bson.M{
		"_id": bson.M{
			"bsonType": "string",
		},
	},
}

func (c *Create) Build(builder structor.Builder) {
	builder.WriteRaw(jsonSchema)
	builder.AddRawVal(defaultID)
}

type Drop struct {
	structor.Drop
}

func drop() structor.Constructor {
	return &Drop{}
}

func (d *Drop) Build(builder structor.Builder) {
	// do nothing
}

type Add struct {
	structor.Add
}

func add() structor.Constructor {
	return &Add{}
}

func (a *Add) Build(builder structor.Builder) {
	// do nothing
}

type Del struct {
	structor.Del
}

func del() structor.Constructor {
	return &Del{}
}

func (d *Del) Build(builder structor.Builder) {
	// do nothing
}

type Modify struct {
	structor.Modify
}

func modify() structor.Constructor {
	return &Modify{}
}

func (m *Modify) Build(builder structor.Builder) {
	// do nothing
}

type Primary struct {
	structor.Primary
}

func primary() structor.Constructor {
	return &Primary{}
}

func (p *Primary) Build(builder structor.Builder) {
	builder.AddIndex(id)
}

type Index struct {
	structor.Index
}

func index() structor.Constructor {
	return &Index{}
}

func (i *Index) Build(builder structor.Builder) {
	for _, value := range i.Fields {
		builder.AddIndex(value.Title)
	}
}

type Unique struct {
	structor.Unique
}

func unique() structor.Constructor {
	return &Unique{}
}

func (u *Unique) Build(builder structor.Builder) {
	for _, value := range u.Fields {
		builder.AddIndex(value.Title)
	}
}

type DropIndex struct {
	structor.DropIndex
}

func dropIndex() structor.Constructor {
	return &DropIndex{}
}

func (d *DropIndex) Build(builder structor.Builder) {
	return
}
