package restDto

import (
	"github.com/abc-valera/netsly-golang/gen/ogen"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/selector"
)

type ISelectorHelper interface {
	GetOrder() ogen.OptOrder
	GetLimit() ogen.OptInt
	GetOffset() ogen.OptInt
}

func NewDomainSelector(params ISelectorHelper) selector.Selector {
	return selector.Selector{
		Limit:  uint(params.GetLimit().Value),
		Offset: uint(params.GetOffset().Value),
	}
}
