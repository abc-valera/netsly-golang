package bunQueryGeneric

import (
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/queryUtil/paging"
	"github.com/uptrace/bun"
)

func ApplyPaing(query *bun.SelectQuery, p paging.Paging) {
	if p.Limit > 0 {
		query = query.Limit(int(p.Limit))
	}

	if p.Offset > 0 {
		query = query.Offset(int(p.Offset))
	}
}
