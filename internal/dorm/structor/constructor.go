package structor

type Constructor interface {
	Set(table string, index string, fields ...*Field)
	Build(builder Builder)
	GetTable() string
	GetIndex() string
}
