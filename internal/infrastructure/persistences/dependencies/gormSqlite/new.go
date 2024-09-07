package gormSqlite

import (
	"os"

	"github.com/abc-valera/netsly-golang/internal/domain/util/coderr"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/dependencies/gormSqlite/gormSqliteDto"
	"gorm.io/gorm"

	"gorm.io/driver/sqlite"
)

// New creates a new sqlite database connection and migrates all models
func New(gormSqliteFolderPath string) (*gorm.DB, error) {
	// Create the folder
	if err := os.MkdirAll(gormSqliteFolderPath, 0o755); err != nil {
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
	db, err := gorm.Open(
		sqlite.Open(gormSqliteFolderPath+"/sqlite.db"+"?_journal=WAL&_timeout=5000&_foreign_keys=true"),
		&gorm.Config{},
	)
	if err != nil {
		return nil, coderr.NewInternalErr(err)
	}

	if err := db.AutoMigrate(
		&gormSqliteDto.User{},
		&gormSqliteDto.Joke{},
		&gormSqliteDto.Like{},
		&gormSqliteDto.Comment{},
		&gormSqliteDto.Room{},
		&gormSqliteDto.RoomMember{},
		&gormSqliteDto.RoomMessage{},
		&gormSqliteDto.FileInfo{},
	); err != nil {
		return nil, coderr.NewInternalErr(err)
	}

	return db, nil
}
