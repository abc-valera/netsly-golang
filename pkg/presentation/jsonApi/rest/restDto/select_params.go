package restDto

import (
	"github.com/abc-valera/netsly-api-golang/gen/ogen"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/query/selectParams"
)

type ISelectParamsHelper interface {
	GetOrder() ogen.OptOrder
	GetLimit() ogen.OptInt
	GetOffset() ogen.OptInt
}

func NewDomainSelectParams(params ISelectParamsHelper) selectParams.SelectParams {
	order := selectParams.OrderDesc
	if params.GetOrder().Value == ogen.Order(selectParams.OrderAsc) {
		order = selectParams.OrderAsc
	}
	return selectParams.NewSelectParams(
		order,
		params.GetLimit().Value,
		params.GetOffset().Value,
	)
}
