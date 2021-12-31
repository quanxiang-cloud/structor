package clause

// Sum Sum
type Sum struct {
	Alias string
	Value interface{}
}

// GetTag GetTag
func (sum *Sum) GetTag() string {
	return "sum"
}

// Set
func (sum *Sum) Set(alias string, values ...interface{}) {
	var value interface{}
	if len(values) > 0 {
		value = values[0]
	}

	sum.Alias = alias
	sum.Value = value
}
