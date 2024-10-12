package selector

// TODO: add support for multiple filters and ordering fields.

const (
	StringValue = "Not empty"
	IntValue    = 1
	FloatValue  = 1.0
)

// Selector is a struct that holds the information needed to perform:
//   - pagination,
//   - filtering,
//   - sorting.
type Selector struct {
	Limit  uint
	Offset uint

	// FilterBy must be a struct, more precisely a domain/model,
	// with all the fields unitialized,
	// except the one that will be used as a filter.
	//
	// The initialized field must be a string,
	// and should be initialized with a selector.StringValue constant.
	FilterBy any
	Filter   string

	// OrderBy must be a struct, more precisely a domain/model,
	// with all the fields unitialized,
	// except the one that will be used as an ordering field.
	//
	// This field must be a string, int, or float,
	// and should be initialized with one of the package constants.
	//
	// Note, that by default, all the ordering is done
	// by the CreatedAt field.
	OrderBy any
	// Order accepts true for ascending order,
	// and false for descending order.
	Order bool
}
