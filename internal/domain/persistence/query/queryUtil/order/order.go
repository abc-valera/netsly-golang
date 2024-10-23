package order

const (
	AscString  = "asc"
	DescString = "desc"
	AscInt     = 1
	DescInt    = -1
	AscFloat   = 1.0
	DescFloat  = -1.0
)

type Order[DomainModel any] struct {
	// By must be a struct, more precisely a domain/model,
	// with all the fields unitialized,
	// except the one that will be used as an ordering field.
	//
	// This field must be a string, int, or float,
	// and should be initialized with one of the package Order* constants.
	//
	// Note, that by default, all the ordering is done
	// by the CreatedAt field in descending order.
	By DomainModel
}
