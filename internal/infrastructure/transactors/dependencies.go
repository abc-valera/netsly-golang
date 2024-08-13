package transactors

import (
	"database/sql"

	"gorm.io/gorm"
)

type Dependencies struct {
	GormSqlite     *gorm.DB
	BoilerSqlite   *sql.DB
	LocalFileSaver string
}
