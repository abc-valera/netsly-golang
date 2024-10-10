package bunSqliteSelector

import (
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/selector"
	"github.com/uptrace/bun"
)

func NewSelect(db bun.IDB, s selector.Selector) *bun.SelectQuery {
	return db.NewSelect().Limit(int(s.Limit)).Offset(int(s.Offset))
}
