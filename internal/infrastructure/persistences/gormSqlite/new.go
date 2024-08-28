package gormSqlite

import (
	"os"

	"github.com/abc-valera/netsly-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/gormSqlite/gormSqliteDto"
	"gorm.io/gorm"

	// Note, that we're importing sqlite driver only here
	// (this lib calls init() function of modernc.org/sqlite package).
	// It would be nice to import it in another place, but it's not possible
	// because of the way gorm works (a shitty ORM).
	"github.com/glebarez/sqlite"
)

// New creates a new sqlite database connection and migrates all models
func New(gormSqliteFolderPath string) (*gorm.DB, error) {
	// Create the folder
	if err := os.MkdirAll(gormSqliteFolderPath, 0o755); err != nil {
		if !os.IsExist(err) {
			return nil, coderr.NewInternalErr(err)
		}
	}

	db, err := gorm.Open(sqlite.Open(gormSqliteFolderPath+"/sqlite.db"), &gorm.Config{})
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
