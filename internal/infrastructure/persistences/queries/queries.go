package queries

import (
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-golang/internal/domain/util/coderr"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/dependencies/bunSqlite/bunSqliteDto"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/queries/bunSqliteQuery"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/queries/bunSqliteQuery/bunSqliteQueryGeneric"
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
		coderr.Fatal("GORM SQLITE QUERIES NOT IMPLEMENTED")
	}

	if dep.BunSqlite != nil {
		queries.User = bunSqliteQueryGeneric.New(dep.BunSqlite, bunSqliteDto.NewUser)
		queries.Joke = bunSqliteQueryGeneric.New(dep.BunSqlite, bunSqliteDto.NewJoke)
		queries.Like = bunSqliteQuery.NewLike(dep.BunSqlite)
		queries.Comment = bunSqliteQueryGeneric.New(dep.BunSqlite, bunSqliteDto.NewComment)
		queries.Room = bunSqliteQueryGeneric.New(dep.BunSqlite, bunSqliteDto.NewRoom)
		queries.RoomMember = bunSqliteQueryGeneric.New(dep.BunSqlite, bunSqliteDto.NewRoomMember)
		queries.RoomMessage = bunSqliteQueryGeneric.New(dep.BunSqlite, bunSqliteDto.NewRoomMessage)
		queries.FileInfo = bunSqliteQueryGeneric.New(dep.BunSqlite, bunSqliteDto.NewFileInfo)
	}

	if dep.LocalFileSaver != "" {
		queries.FileContent = localFileSaverQuery.New(dep.LocalFileSaver)
	}

	return queries
}
