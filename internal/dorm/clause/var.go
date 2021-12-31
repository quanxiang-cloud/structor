package clause

import "errors"

var (
	// ErrAssertBuilder assert builder fail
	ErrAssertBuilder = errors.New("assert builder fail")
)

// PageData entities of table
type PageData struct {
	Value []Data
	Total int64
}

// FindOptions page options
type FindOptions struct {
	Page int64
	Size int64
	Sort []string
}

// Data dest of table
type Data map[string]interface{}
