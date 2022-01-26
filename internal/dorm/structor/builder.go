package structor

type StructBuilder interface {
	WriteRaw(string)
	Unique(bool)
}

type Builder interface {
	StructBuilder
}
