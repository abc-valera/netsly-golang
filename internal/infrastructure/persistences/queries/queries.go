package queries

import (
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-golang/internal/domain/util/coderr"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/dependencies/bunSqlite/bunSqliteDto"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/queries/bunSqliteQuery"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/queries/bunSqliteQuery/bunQueryGeneric"
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

	// TODO: FINISH GORM SQLITE QUERIES
	if dep.GormSqlite != nil {
		// queries.User = gormSqliteQuery.NewUser(dep.GormSqlite)
		// queries.Joke = gormSqliteQuery.NewJoke(dep.GormSqlite)
		// queries.Like = gormSqliteQuery.NewLike(dep.GormSqlite)
		// queries.Comment = gormSqliteQuery.NewComment(dep.GormSqlite)
		// queries.Room = gormSqliteQuery.NewRoom(dep.GormSqlite)
		// queries.RoomMember = gormSqliteQuery.NewRoomMember(dep.GormSqlite)
		// queries.RoomMessage = gormSqliteQuery.NewRoomMessage(dep.GormSqlite)
		// queries.FileInfo = gormSqliteQuery.NewFileInfo(dep.GormSqlite)

		coderr.Fatal("GORM SQLITE QUERIES NOT IMPLEMENTED")
	}

	if dep.BunSqlite != nil {
		queries.User = bunQueryGeneric.NewGetOneGetMany(dep.BunSqlite, bunSqliteDto.NewUser)
		queries.Joke = bunQueryGeneric.NewGetOneGetMany(dep.BunSqlite, bunSqliteDto.NewJoke)
		queries.Like = bunSqliteQuery.NewLike(dep.BunSqlite)
		queries.Comment = bunQueryGeneric.NewGetOneGetMany(dep.BunSqlite, bunSqliteDto.NewComment)
		queries.Room = bunQueryGeneric.NewGetOneGetMany(dep.BunSqlite, bunSqliteDto.NewRoom)
		queries.RoomMember = bunQueryGeneric.NewGetOneGetMany(dep.BunSqlite, bunSqliteDto.NewRoomMember)
		queries.RoomMessage = bunQueryGeneric.NewGetOneGetMany(dep.BunSqlite, bunSqliteDto.NewRoomMessage)
		queries.FileInfo = bunQueryGeneric.NewGetOneGetMany(dep.BunSqlite, bunSqliteDto.NewFileInfo)
	}

	if dep.LocalFileSaver != "" {
		queries.FileContent = localFileSaverQuery.New(dep.LocalFileSaver)
	}

	return queries
}
