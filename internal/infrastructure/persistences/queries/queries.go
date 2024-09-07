package queries

import (
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/queries/bunSqliteQuery"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/queries/gormSqliteQuery"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/queries/localFileSaverQuery"
	"github.com/uptrace/bun"
	"gorm.io/gorm"
)

type Dependency struct {
	GormSqlite     *gorm.DB
	BunSqlite      bun.IDB
	LocalFileSaver string
}

func New(dep Dependency) query.Queries {
	var queries query.Queries

	if dep.GormSqlite != nil {
		queries.User = gormSqliteQuery.NewUser(dep.GormSqlite)
		queries.Joke = gormSqliteQuery.NewJoke(dep.GormSqlite)
		queries.Like = gormSqliteQuery.NewLike(dep.GormSqlite)
		queries.Comment = gormSqliteQuery.NewComment(dep.GormSqlite)
		queries.Room = gormSqliteQuery.NewRoom(dep.GormSqlite)
		queries.RoomMember = gormSqliteQuery.NewRoomMember(dep.GormSqlite)
		queries.RoomMessage = gormSqliteQuery.NewRoomMessage(dep.GormSqlite)
		queries.FileInfo = gormSqliteQuery.NewFileInfo(dep.GormSqlite)
	}

	if dep.BunSqlite != nil {
		queries.User = bunSqliteQuery.NewUser(dep.BunSqlite)
		queries.Joke = bunSqliteQuery.NewJoke(dep.BunSqlite)
		queries.Like = bunSqliteQuery.NewLike(dep.BunSqlite)
		queries.Comment = bunSqliteQuery.NewComment(dep.BunSqlite)
		queries.Room = bunSqliteQuery.NewRoom(dep.BunSqlite)
		queries.RoomMember = bunSqliteQuery.NewRoomMember(dep.BunSqlite)
		queries.RoomMessage = bunSqliteQuery.NewRoomMessage(dep.BunSqlite)
		queries.FileInfo = bunSqliteQuery.NewFileInfo(dep.BunSqlite)
	}

	if dep.LocalFileSaver != "" {
		queries.FileContent = localFileSaverQuery.New(dep.LocalFileSaver)
	}

	return queries
}
