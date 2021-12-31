package clause

type logical struct {
	Vars []interface{}
}

// LSet set value
func (l *logical) LSet(values ...interface{}) {
	l.Vars = values
}

// AND and
type AND struct {
	logical
}

// GetTag get tag
func (and *AND) GetTag() string {
	return "and"
}

// Set set value
func (and *AND) Set(column string, values ...interface{}) {
	and.LSet(values...)
}

// OR or
type OR struct {
	logical
}

// GetTag get tag
func (or *OR) GetTag() string {
	return "or"
}

// Set set value
func (or *OR) Set(column string, values ...interface{}) {
	or.LSet(values...)
}

// NOR nor
type NOR struct {
	logical
}

// GetTag get tag
func (nor *NOR) GetTag() string {
	return "nor"
}

// Set set value
func (nor *NOR) Set(column string, values ...interface{}) {
	nor.LSet(values...)
}
