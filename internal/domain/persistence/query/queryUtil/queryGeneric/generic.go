package queryGeneric

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/queryUtil/filter"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/queryUtil/selector"
)

type IGetOneGetMany[DomainModel any] interface {
	IGetOne[DomainModel]
	IGetMany[DomainModel]
}

type IGetOne[DomainModel any] interface {
	GetOne(ctx context.Context, fitlerOptions ...filter.Option[DomainModel]) (DomainModel, error)
}

type IGetMany[DomainModel any] interface {
	GetMany(ctx context.Context, selectorOptions ...selector.Option[DomainModel]) ([]DomainModel, error)
}
