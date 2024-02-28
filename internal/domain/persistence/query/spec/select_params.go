package spec

// SelectParams represents query data for specifying select details.
type SelectParams interface {
	Order() string
	Limit() int
	Offset() int
}

// NewSelectParams creates a new SelectParams instance.
//
//   - If the order is not "asc" or "desc", it will be set to "desc".
//   - If the limit is less than 0, it will be set to 5.
//   - If the offset is less than 0, it will be set to 0.
func NewSelectParams(order string, limit, offset int) SelectParams {
	if order != "asc" && order != "desc" {
		order = "desc"
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
	order  string
	limit  int
	offset int
}

func (s selectParams) Order() string {
	return s.order
}

func (s selectParams) Limit() int {
	return s.limit
}

func (s selectParams) Offset() int {
	return s.offset
}
