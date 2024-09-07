package persistences

import (
	"context"
	"os"
	"strings"

	"github.com/abc-valera/netsly-golang/internal/domain/persistence"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-golang/internal/domain/util/coderr"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/commands"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/dependencies/bunSqlite"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/dependencies/gormSqlite"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/dependencies/localFileSaver"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/queries"
	"github.com/uptrace/bun"
	"gorm.io/gorm"
)

// DB contains all the dependencies needed to start a transaction on every datasource.
type DB struct {
	gormSqlite     *gorm.DB
	bunSqlite      *bun.DB
	localFileSaver string
}

// NewDB creates a new DB struct with all the dependencies.
// Note, that the dependencies are optional and any of them can be nil.
func NewDB() persistence.IDB {
	var db DB

	if gormSqliteEnv := strings.TrimSpace(os.Getenv("GORM_SQLITE_FOLDER_PATH")); gormSqliteEnv != "" {
		db.gormSqlite = coderr.Must(gormSqlite.New(gormSqliteEnv))
	}

	if bunSqliteEnv := strings.TrimSpace(os.Getenv("BUN_SQLITE_FOLDER_PATH")); bunSqliteEnv != "" {
		db.bunSqlite = coderr.Must(bunSqlite.New(bunSqliteEnv))
	}

	if localFileSaverEnv := strings.TrimSpace(os.Getenv("LOCAL_FILE_SAVER_FOLDER_PATH")); localFileSaverEnv != "" {
		db.localFileSaver = coderr.Must(localFileSaver.New(localFileSaverEnv))
	}

	return db
}

func (db DB) RunInTX(
	ctx context.Context,
	fn func(context.Context, command.Commands, query.Queries) error,
) error {
	tx, err := newTX(ctx, db)
	if err != nil {
		return err
	}

	if err := fn(ctx, tx.Commands(), tx.Queries()); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (db DB) BeginTX(ctx context.Context) (persistence.ITX, error) {
	return newTX(ctx, db)
}

func (db DB) Commands() command.Commands {
	return commands.New(commands.Dependency{
		GormSqlite:     db.gormSqlite,
		BunSqlite:      db.bunSqlite,
		LocalFileSaver: db.localFileSaver,
	})
}

func (db DB) Queries() query.Queries {
	return queries.New(queries.Dependency{
		GormSqlite:     db.gormSqlite,
		BunSqlite:      db.bunSqlite,
		LocalFileSaver: db.localFileSaver,
	})
}
