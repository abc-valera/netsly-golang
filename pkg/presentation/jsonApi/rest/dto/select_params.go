package dto

import (
	"github.com/abc-valera/netsly-api-golang/gen/ogen"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/query/spec"
)

type ISelectParamsHelper interface {
	GetOrder() ogen.OptOrder
	GetLimit() ogen.OptInt
	GetOffset() ogen.OptInt
}

func NewDomainSelectParams(params ISelectParamsHelper) spec.SelectParams {
	order := spec.OrderDesc
	if params.GetOrder().Value == ogen.Order(spec.OrderAsc) {
		order = spec.OrderAsc
	}
	return spec.NewSelectParams(
		order,
		params.GetLimit().Value,
		params.GetOffset().Value,
	)
}
