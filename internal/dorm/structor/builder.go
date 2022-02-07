package structor

type StructBuilder interface {
	WriteRaw(string)
	Unique(bool)
	IndexName(names []string)
	Create(bool, ...string)
}

type Builder interface {
	StructBuilder
}
