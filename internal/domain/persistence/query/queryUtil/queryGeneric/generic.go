// queryGeneric package contains generic interfaces for the datasource queries.
//
// These queries are simple and can't be used as abstractions for all possible SELECT clauses.
// Instead, they are used to provide a common interface for the most common SELECT queries.
//
// WHERE queries with pagination, multiple filters and ordering fields can be performed.
package queryGeneric

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/queryUtil/selector"
)

type IGetOneGetMany[Model any] interface {
	IGet[Model]
	IGetMany[Model]
}

type IGet[Model any] interface {
	// Get retrieves a single record from the datasource.
	// The model parameter is used a filter.
	Get(context.Context, Model) (Model, error)
}

type IGetMany[Model any] interface {
	GetMany(context.Context, ...selector.Option[Model]) ([]Model, error)
}
