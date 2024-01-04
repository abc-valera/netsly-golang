package spec

import "github.com/abc-valera/flugo-api-golang/internal/core/domain/codeerr"

var (
	ErrInvalidOrder  = codeerr.NewMessageErr(codeerr.CodeInvalidArgument, "Order must be 'asc' or 'desc'")
	ErrInvalidLimit  = codeerr.NewMessageErr(codeerr.CodeInvalidArgument, "Limit must be greater than 0")
	ErrInvalidOffset = codeerr.NewMessageErr(codeerr.CodeInvalidArgument, "Offset must be greater than 0")
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