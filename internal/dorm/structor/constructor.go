package structor

type Constructor interface {
	Set(table string, values ...*Field)

	Build(table string, builder Builder)
}
