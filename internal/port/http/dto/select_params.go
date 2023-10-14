package dto

import (
	"github.com/abc-valera/flugo-api-golang/gen/ogen"
	"github.com/abc-valera/flugo-api-golang/internal/domain/repository/spec"
)

type ISelectParams interface {
	GetOrderBy() ogen.OptString
	GetOrder() ogen.OptOrder
	GetLimit() int
	GetOffset() int
}

func NewDomainSelectParams(selectParams ISelectParams) (spec.SelectParams, error) {
	return spec.NewSelectParams(
		selectParams.GetOrderBy().Value,
		string(selectParams.GetOrder().Value),
		uint(selectParams.GetLimit()),
		uint(selectParams.GetOffset()),
	)
}
