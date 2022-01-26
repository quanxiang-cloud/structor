package structor

type Constructor interface {
	Set(table string, values ...*Field)

	Build(builder Builder)
}
