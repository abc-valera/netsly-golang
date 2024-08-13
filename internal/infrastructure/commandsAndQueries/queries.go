package commandsAndQueries

import (
	"github.com/abc-valera/netsly-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/boilerSqlite/boilerSqliteQuery"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/gormSqlite/gormSqliteQuery"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/localFileSaver/localFileSaverQuery"
)

func newQueries(deps Dependencies) (persistence.Queries, error) {
	var queries persistence.Queries

	if deps.GormSqlite != nil {
		queries.User = gormSqliteQuery.NewUser(deps.GormSqlite)
		queries.Joke = gormSqliteQuery.NewJoke(deps.GormSqlite)
		queries.Like = gormSqliteQuery.NewLike(deps.GormSqlite)
		queries.Comment = gormSqliteQuery.NewComment(deps.GormSqlite)
		queries.Room = gormSqliteQuery.NewRoom(deps.GormSqlite)
		queries.RoomMember = gormSqliteQuery.NewRoomMember(deps.GormSqlite)
		queries.RoomMessage = gormSqliteQuery.NewRoomMessage(deps.GormSqlite)
		queries.FileInfo = gormSqliteQuery.NewFileInfo(deps.GormSqlite)
	}

	if deps.BoilerSqlite != nil {
		queries.User = boilerSqliteQuery.NewUser(deps.BoilerSqlite)
		queries.Joke = boilerSqliteQuery.NewJoke(deps.BoilerSqlite)
		queries.Like = boilerSqliteQuery.NewLike(deps.BoilerSqlite)
		queries.Comment = boilerSqliteQuery.NewComment(deps.BoilerSqlite)
		queries.Room = boilerSqliteQuery.NewRoom(deps.BoilerSqlite)
		queries.RoomMember = boilerSqliteQuery.NewRoomMember(deps.BoilerSqlite)
		queries.RoomMessage = boilerSqliteQuery.NewRoomMessage(deps.BoilerSqlite)
		queries.FileInfo = boilerSqliteQuery.NewFileInfo(deps.BoilerSqlite)
	}

	if deps.LocalFileSaver != "" {
		queries.FileContent = localFileSaverQuery.New(deps.LocalFileSaver)
	}

	// Check if all queries are initialized
	if err := coderr.CheckIfStructHasEmptyFields(queries); err != nil {
		return persistence.Queries{}, coderr.NewInternalErr(err)
	}

	return queries, nil
}
