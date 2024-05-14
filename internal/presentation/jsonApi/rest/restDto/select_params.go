package restDto

import (
	"github.com/abc-valera/netsly-api-golang/gen/ogen"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query/selector"
)

type ISelectorHelper interface {
	GetOrder() ogen.OptOrder
	GetLimit() ogen.OptInt
	GetOffset() ogen.OptInt
}

func NewDomainSelector(params ISelectorHelper) selector.Selector {
	order := selector.OrderDesc
	if params.GetOrder().Value == ogen.Order(selector.OrderAsc) {
		order = selector.OrderAsc
	}
	return selector.Selector{
		Order:  order,
		Limit:  uint(params.GetLimit().Value),
		Offset: uint(params.GetOffset().Value),
	}
}
