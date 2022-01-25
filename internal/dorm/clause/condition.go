package clause

type Terms struct {
	Column string
	Values []interface{}
}

// GetTag get tag
func (t *Terms) GetTag() string {
	return "terms"
}

// Set set value
func (t *Terms) Set(column string, values ...interface{}) {
	t.Column = column
	t.Values = values
}

type Match struct {
	Column string
	Values interface{}
}

// GetTag get tag
func (m *Match) GetTag() string {
	return "match"
}

// Set set value
func (m *Match) Set(column string, values ...interface{}) {
	var checkString = func(value interface{}) bool {
		_, ok := value.(string)
		return ok
	}
	m.Column = column
	if len(values) != 1 || !checkString(values[0]) {
		m.Values = "NULL"
	} else {
		m.Values = values[0]
	}
}

type Term struct {
	Column string
	Values []interface{}
}

// GetTag get tag
func (t *Term) GetTag() string {
	return "term"
}

// Set set value
func (t *Term) Set(column string, values ...interface{}) {
	t.Column = column
	t.Values = values
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
