package structor

type StructBuilder interface {
	WriteRaw(string)
}

type Builder interface {
	StructBuilder
}
