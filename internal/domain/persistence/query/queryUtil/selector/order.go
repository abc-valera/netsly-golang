package selector

import "time"

var (
	AscString  = "asc"
	DescString = "desc"
	AscInt     = 1
	DescInt    = -1
	AscFloat   = 1.0
	DescFloat  = -1.0
	AscTime    = time.UnixMilli(1)
	DescTime   = time.UnixMilli(2)
)

// Order is used as an abstraction over the datasource ordering mechanism.
type Order[Model any] struct {
	// By must be a struct, more precisely a domain/model,
	// with all the fields unitialized,
	// except the fields that will be used as the ordering fields.
	//
	// Each of these fields must be a string, int, float or time.Time,
	// and should be initialized with one of the package constants.
	//
	// Note, that if none of the fields are initialized,
	// the ordering will be done by the CreatedAt field.
	By Model
}
