package gormSqliteErrors

import (
	"github.com/abc-valera/netsly-golang/internal/domain/util/coderr"
	"gorm.io/gorm"
)

func HandleQueryResult(result *gorm.DB) error {
	if result.Error != nil {
		// Handler errors here
		return coderr.NewInternalErr(result.Error)
	}

	return nil
}
