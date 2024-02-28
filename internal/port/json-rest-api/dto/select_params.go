package dto

import (
	"github.com/abc-valera/netsly-api-golang/gen/ogen"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query/spec"
)

type ISelectParamsHelper interface {
	GetOrder() ogen.OptOrder
	GetLimit() ogen.OptInt
	GetOffset() ogen.OptInt
}

func NewDomainSelectParams(params ISelectParamsHelper) spec.SelectParams {
	return spec.NewSelectParams(
		string(params.GetOrder().Value),
		params.GetLimit().Value,
		params.GetOffset().Value,
	)
}
