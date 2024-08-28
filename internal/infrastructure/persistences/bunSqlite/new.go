package bunSqlite

// Note, that we don't import sqlite driver here
// because it is already imported in the gormSqlite/new.go file.

import (
	"context"
	"database/sql"
	"os"

	"github.com/abc-valera/netsly-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/bunSqlite/bunSqliteDto"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
)

// New creates a new sqlite database connection and migrates all models
func New(bunSqliteFolderPath string) (*bun.DB, error) {
	// Create the folder if not exists
	if err := os.MkdirAll(bunSqliteFolderPath, 0o755); err != nil {
		if !os.IsExist(err) {
			return nil, coderr.NewInternalErr(err)
		}
	}

	sqlDB, err := sql.Open("sqlite", bunSqliteFolderPath+"/sqlite.db")
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
