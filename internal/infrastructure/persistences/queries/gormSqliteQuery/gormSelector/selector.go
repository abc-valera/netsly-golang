package gormSelector

import (
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/selector"
	"gorm.io/gorm"
)

func WithSelector(db *gorm.DB, sel selector.Selector) *gorm.DB {
	return db.Limit(int(sel.Limit)).Offset(int(sel.Offset))
}
