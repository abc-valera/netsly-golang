package bunSqlite

// Note, that we don't import sqlite driver here
// because it is already imported in the gormSqlite/new.go file.

import (
	"context"
	"database/sql"
	"os"

	"github.com/abc-valera/netsly-golang/internal/domain/util/coderr"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/dependencies/bunSqlite/bunSqliteDto"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"

	_ "github.com/mattn/go-sqlite3"
)

// New creates a new sqlite database connection and migrates all models
func New(bunSqliteFolderPath string) (*bun.DB, error) {
	// Create the folder if not exists
	if err := os.MkdirAll(bunSqliteFolderPath, 0o755); err != nil {
		if !os.IsExist(err) {
			return nil, coderr.NewInternalErr(err)
		}
	}

	// Initialize the database
	//
	// Consider the following:
	// - Set WAL mode, so readers and writers can access the database concurrently
	// - Set busy timeout, so concurrent writers wait on each other instead of erroring immediately
	// - Enable foreign key checks
	sqlDB, err := sql.Open(
		"sqlite3",
		bunSqliteFolderPath+"/sqlite.db"+"?_journal=WAL&_timeout=5000&_foreign_keys=true",
	)
	if err != nil {
		return nil, coderr.NewInternalErr(err)
	}

	bunDB := bun.NewDB(sqlDB, sqlitedialect.New())

	coderr.Must(bunDB.NewCreateTable().Model((*bunSqliteDto.User)(nil)).IfNotExists().Exec(context.Background()))
	coderr.Must(bunDB.NewCreateTable().Model((*bunSqliteDto.Joke)(nil)).IfNotExists().Exec(context.Background()))
	coderr.Must(bunDB.NewCreateTable().Model((*bunSqliteDto.Like)(nil)).IfNotExists().Exec(context.Background()))
	coderr.Must(bunDB.NewCreateTable().Model((*bunSqliteDto.Comment)(nil)).IfNotExists().Exec(context.Background()))
	coderr.Must(bunDB.NewCreateTable().Model((*bunSqliteDto.Room)(nil)).IfNotExists().Exec(context.Background()))
	coderr.Must(bunDB.NewCreateTable().Model((*bunSqliteDto.RoomMember)(nil)).IfNotExists().Exec(context.Background()))
	coderr.Must(bunDB.NewCreateTable().Model((*bunSqliteDto.RoomMessage)(nil)).IfNotExists().Exec(context.Background()))
	coderr.Must(bunDB.NewCreateTable().Model((*bunSqliteDto.FileInfo)(nil)).IfNotExists().Exec(context.Background()))

	return bunDB, nil
}
