package bunSqliteQueryGeneric

import (
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/queryUtil/selector"
	"github.com/uptrace/bun"
)

func ApplyPaging(query *bun.SelectQuery, p selector.Paging) {
	if p.Limit > 0 {
		query = query.Limit(int(p.Limit))
	}

	if p.Offset > 0 {
		query = query.Offset(int(p.Offset))
	}
}
