package gormSelector

import (
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/selector"
	"gorm.io/gorm"
)

func WithSelector(db *gorm.DB, sel selector.Selector) *gorm.DB {
	query := db.Limit(int(sel.Limit)).Offset(int(sel.Offset))
	if sel.Order == selector.OrderAsc {
		query = query.Order("created_at asc")
	}
	if sel.Order == selector.OrderDesc {
		query = query.Order("created_at desc")
	}
	return query
}
