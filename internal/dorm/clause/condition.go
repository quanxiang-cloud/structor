package clause

// IN Whether a value is within a set of values
type IN struct {
	Column string
	Values []interface{}
}

// GetTag get tag
func (in *IN) GetTag() string {
	return "in"
}

// Set set value
func (in *IN) Set(column string, values ...interface{}) {
	in.Column = column
	in.Values = values
}

// LIKE fuzzy query
type LIKE struct {
	Column string
	Values interface{}
}

// GetTag get tag
func (like *LIKE) GetTag() string {
	return "like"
}

// Set set value
func (like *LIKE) Set(column string, values ...interface{}) {
	var checkString = func(value interface{}) bool {
		_, ok := value.(string)
		return ok
	}
	like.Column = column
	if len(values) != 1 || !checkString(values[0]) {
		like.Values = "NULL"
	} else {
		like.Values = values[0]
	}
}

// EQUAL equal
type EQUAL struct {
	Column string
	Values []interface{}
}

// GetTag get tag
func (equal *EQUAL) GetTag() string {
	return "eq"
}

// Set set value
func (equal *EQUAL) Set(column string, values ...interface{}) {
	equal.Column = column
	equal.Values = values
}

type conditionOP struct {
	Column string
	Values interface{}
}

// CSet set value
func (c *conditionOP) CSet(column string, values interface{}) {
	c.Column = column
	if values == nil {
		values = "NULL"
	}
	c.Values = values
}

// LT less than
type LT struct {
	conditionOP
}

// GetTag get tag
func (lt *LT) GetTag() string {
	return "lt"
}

// Set set value
func (lt *LT) Set(column string, values ...interface{}) {
	var value interface{}
	if len(values) > 0 {
		value = values[0]
	}
	lt.CSet(column, value)
}

// LTE less than or equal
type LTE struct {
	conditionOP
}

// GetTag get tag
func (lte *LTE) GetTag() string {
	return "lte"
}

// Set set value
func (lte *LTE) Set(column string, values ...interface{}) {
	var value interface{}
	if len(values) > 0 {
		value = values[0]
	}
	lte.CSet(column, value)
}

// GT greater than
type GT struct {
	conditionOP
}

// GetTag get tag
func (gt *GT) GetTag() string {
	return "gt"
}

// Set set value
func (gt *GT) Set(column string, values ...interface{}) {
	var value interface{}
	if len(values) > 0 {
		value = values[0]
	}
	gt.CSet(column, value)
}

// GTE greater than or equal
type GTE struct {
	conditionOP
}

// GetTag get tag
func (gte *GTE) GetTag() string {
	return "gte"
}

// Set set value
func (gte *GTE) Set(column string, values ...interface{}) {
	var value interface{}
	if len(values) > 0 {
		value = values[0]
	}
	gte.CSet(column, value)
}
