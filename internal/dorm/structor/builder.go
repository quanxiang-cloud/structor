package structor

type StructBuilder interface {
	WriteRaw(string)
	AddRawVal(interface{})
	AddIndex(string)
	Unique(bool)
	IndexName(names []string)
}

type Builder interface {
	StructBuilder
}
