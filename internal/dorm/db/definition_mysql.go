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

func (c *Create) Build(builder structor.Builder) {
	builder.WriteRaw(fmt.Sprintf(" CREATE TABLE %s ( ", c.Column))
	builder.WriteRaw(c.Values.Convert())
	builder.WriteRaw(fmt.Sprintf(") ENGINE=%s DEFAULT CHARSET=%s COLLATE=%s;", engine, charset, collate))
}

type Add struct {
	structor.Add
}

func add() structor.Constructor {
	return &Add{}
}

func (a *Add) Build(builder structor.Builder) {
	builder.WriteRaw(fmt.Sprintf("ALTER TABLE %s ADD COLUMN ", a.Column))
	builder.WriteRaw(a.Values.Convert())
	builder.WriteRaw(";")
}

type Modify struct {
	structor.Modify
}

func modify() structor.Constructor {
	return &Modify{}
}

func (m *Modify) Build(builder structor.Builder) {
	builder.WriteRaw(fmt.Sprintf("ALTER TABLE %s MODIFY COLUMN ", m.Column))
	builder.WriteRaw(m.Values.Convert())
	builder.WriteRaw(";")
}
