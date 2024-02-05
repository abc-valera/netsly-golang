package spec

import "github.com/abc-valera/netsly-api-golang/internal/domain/coderr"

var (
	ErrInvalidOrder  = coderr.NewMessage(coderr.CodeInvalidArgument, "Order must be 'asc' or 'desc'")
	ErrInvalidLimit  = coderr.NewMessage(coderr.CodeInvalidArgument, "Limit must be greater than 0")
	ErrInvalidOffset = coderr.NewMessage(coderr.CodeInvalidArgument, "Offset must be greater than 0")
)

// SelectParams represents query data for specifying select details.
type SelectParams struct {
	Order  string // Order is order of sorting ('acs' or 'desc')
	Limit  int    // Limit limits number of returned units
	Offset int    // Offset sets an offset for returned units
}

// NewSelectParams validates input and creates new SelectParams instance.
func NewSelectParams(
	order string,
	limit int,
	offset int,
) (SelectParams, error) {
	if order != "asc" && order != "desc" {
		return SelectParams{}, ErrInvalidOrder
	}
	if limit < 0 {
		return SelectParams{}, ErrInvalidLimit
	}
	if offset < 0 {
		return SelectParams{}, ErrInvalidOffset
	}
	return SelectParams{
		Order:  order,
		Limit:  limit,
		Offset: offset,
	}, nil
}
