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

func NewDomainSelectParams(params ISelectParamsHelper) (spec.SelectParams, error) {
	order := string(params.GetOrder().Value)
	if !params.GetOrder().IsSet() {
		order = "desc"
	}
	limit := params.GetLimit().Value
	if !params.GetLimit().IsSet() {
		limit = 10
	}
	offset := params.GetOffset().Value
	if !params.GetOffset().IsSet() {
		offset = 0
	}
	return spec.NewSelectParams(
		order,
		limit,
		offset,
	)
}
