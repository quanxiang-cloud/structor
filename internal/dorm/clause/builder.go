package clause

// Writer write
type Writer interface {
	WriteByte(byte) error
	WriteString(string) (int, error)
}

// QueryBuilder db condition builder
type QueryBuilder interface {
	Writer
	WriteQuoted(field string)
	AddVar(value interface{})
	GetVar() interface{}
}

// AggBuilder AggBuilder
type AggBuilder interface {
	WriteQuotedAgg(field string)
	AddAggVar(key string, value interface{})
}

//Builder Builder
type Builder interface {
	QueryBuilder
	AggBuilder
}
