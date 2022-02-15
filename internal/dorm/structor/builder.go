package structor

type StructBuilder interface {
	WriteRaw(string)
	AddRawVal(interface{})
	AddIndex(string)
}

type Builder interface {
	StructBuilder
}
