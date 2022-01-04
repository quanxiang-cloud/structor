package clause

type logical struct {
	Vars []interface{}
}

// LSet set value
func (l *logical) LSet(values ...interface{}) {
	l.Vars = values
}

type MUST struct {
	logical
}

// GetTag get tag
func (m *MUST) GetTag() string {
	return "must"
}

// Set set value
func (m *MUST) Set(column string, values ...interface{}) {
	m.LSet(values...)
}

type SHOULD struct {
	logical
}

// GetTag get tag
func (s *SHOULD) GetTag() string {
	return "should"
}

// Set set value
func (s *SHOULD) Set(column string, values ...interface{}) {
	s.LSet(values...)
}

type MUSTNOT struct {
	logical
}

// GetTag get tag
func (m *MUSTNOT) GetTag() string {
	return "must_not"
}

// Set set value
func (m *MUSTNOT) Set(column string, values ...interface{}) {
	m.LSet(values...)
}

type RANGE struct {
	logical
}

func (r *RANGE) GetTag() string {
	return "range"
}

func (r *RANGE) Set(column string, values ...interface{}) {
	r.LSet(values...)
}
