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

func (c *Create) Build(builder structor.Builder) {
	// do nothing
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
