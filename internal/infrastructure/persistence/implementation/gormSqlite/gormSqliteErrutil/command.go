package gormSqliteErrutil

import (
	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
	"gorm.io/gorm"
)

func HandleCommandResult(result *gorm.DB) error {
	if result.Error != nil {
		// Handler errors here
		return coderr.NewInternalErr(result.Error)
	}

	return nil
}
