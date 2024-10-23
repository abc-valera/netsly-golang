package bunQueryGeneric

import (
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/queryUtil/selector"
	"github.com/uptrace/bun"
)

func ApplySelector[DomainModel, BunModel any](
	query *bun.SelectQuery,
	dto func(DomainModel) BunModel,
	genericSelector selector.Selector[DomainModel],
) error {
	ApplyPaing(query, genericSelector.Paging)

	for _, filter := range genericSelector.Filters {
		if err := ApplyFilter[DomainModel, BunModel](query, dto, filter); err != nil {
			return err
		}
	}

	return ApplyOrder[DomainModel, BunModel](query, dto, genericSelector.Order)
}
