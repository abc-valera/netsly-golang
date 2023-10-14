package spec

import "github.com/abc-valera/flugo-api-golang/internal/domain/codeerr"

var (
	ErrInvalidOrder = codeerr.NewMsgErr(codeerr.CodeInvalidArgument, "Order must be 'asc' or 'desc'")
)

// SelectParams represents query data for specifying select details.
type SelectParams struct {
	OrderBy string // OrderField is field by which sorting will be performed (usually is 'created_at')
	Order   string // Order is order of sorting ('acs' or 'desc')
	Limit   uint   // Limit limits number of returned units
	Offset  uint   // Offset sets an offset for returned units
}

func NewSelectParams(orderBy, order string, limit, offset uint) (SelectParams, error) {
	if orderBy == "" {
		orderBy = "created_at"
	}
	if order == "" {
		order = "desc"
	}
	if order != "asc" && order != "desc" {
		return SelectParams{}, ErrInvalidOrder
	}
	return SelectParams{
		OrderBy: orderBy,
		Order:   order,
		Limit:   limit,
		Offset:  offset,
	}, nil
}
