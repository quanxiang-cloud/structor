package structor

type StructBuilder interface {
	WriteRaw(string)
	Unique(bool)
	IndexName(names []string)
}

type Builder interface {
	StructBuilder
}
