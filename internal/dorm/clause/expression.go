package clause

type Expression interface {
	GetTag() string
	Set(column string, values ...interface{})

	Build(builder Builder)
}
