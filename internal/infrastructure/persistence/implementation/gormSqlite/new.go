package gormSqlite

import (
	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/gormSqlite/gormSqliteDto"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// New creates a new sqlite database connection and migrates all models
func New(sqlitePath string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(sqlitePath), &gorm.Config{})
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
