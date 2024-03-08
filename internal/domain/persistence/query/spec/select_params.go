package spec

type Order string

const (
	OrderAsc  Order = "asc"
	OrderDesc Order = "desc"
)

// SelectParams represents query data for specifying select details.
type SelectParams interface {
	Order() Order
	Limit() int
	Offset() int
}

// NewSelectParams creates a new SelectParams instance.
//
//   - If the order is not "asc" or "desc", it will be set to "desc".
//   - If the limit is less than 0, it will be set to 5.
//   - If the offset is less than 0, it will be set to 0.
func NewSelectParams(order Order, limit, offset int) SelectParams {
	if order != OrderAsc && order != OrderDesc {
		order = OrderDesc
	}
	if limit < 0 {
		limit = 5
	}
	if offset < 0 {
		offset = 0
	}
	return &selectParams{
		order:  order,
		limit:  limit,
		offset: offset,
	}
}

type selectParams struct {
	order  Order
	limit  int
	offset int
}

func (s selectParams) Order() Order {
	return s.order
}

func (s selectParams) Limit() int {
	return s.limit
}

func (s selectParams) Offset() int {
	return s.offset
}
