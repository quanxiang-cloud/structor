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

// Avg Avg
type Avg struct {
	Alias string
	Value interface{}
}

// GetTag GetTag
func (avg *Avg) GetTag() string {
	return "avg"
}

// Set
func (avg *Avg) Set(alias string, values ...interface{}) {
	var value interface{}
	if len(values) > 0 {
		value = values[0]
	}

	avg.Alias = alias
	avg.Value = value
}

// Min Min
type Min struct {
	Alias string
	Value interface{}
}

// GetTag GetTag
func (min *Min) GetTag() string {
	return "min"
}

// Set
func (min *Min) Set(alias string, values ...interface{}) {
	var value interface{}
	if len(values) > 0 {
		value = values[0]
	}

	min.Alias = alias
	min.Value = value
}

// Max Max
type Max struct {
	Alias string
	Value interface{}
}

// GetTag GetTag
func (max *Max) GetTag() string {
	return "max"
}

// Set
func (max *Max) Set(alias string, values ...interface{}) {
	var value interface{}
	if len(values) > 0 {
		value = values[0]
	}

	max.Alias = alias
	max.Value = value
}

// Count Max
type Count struct {
	Alias string
	Value interface{}
}

// GetTag GetTag
func (count *Count) GetTag() string {
	return "count"
}

// Set
func (count *Count) Set(alias string, values ...interface{}) {
	var value interface{}
	if len(values) > 0 {
		value = values[0]
	}

	count.Alias = alias
	count.Value = value
}
