package transactors

import (
	"database/sql"

	"github.com/uptrace/bun"
	"gorm.io/gorm"
)

type Dependencies struct {
	GormSqlite     *gorm.DB
	BoilerSqlite   *sql.DB
	BunSqlite      *bun.DB
	LocalFileSaver string
}
