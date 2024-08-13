package commandsAndQueries

import (
	"github.com/abc-valera/netsly-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/boilerSqlite/boilerSqliteCommand"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/gormSqlite/gormSqliteCommand"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/localFileSaver/localFileSaverCommand"
)

func newCommands(deps Dependencies) (persistence.Commands, error) {
	var commands persistence.Commands

	if deps.GormSqlite != nil {
		commands.User = gormSqliteCommand.NewUser(deps.GormSqlite)
		commands.Joke = gormSqliteCommand.NewJoke(deps.GormSqlite)
		commands.Like = gormSqliteCommand.NewLike(deps.GormSqlite)
		commands.Comment = gormSqliteCommand.NewComment(deps.GormSqlite)
		commands.Room = gormSqliteCommand.NewRoom(deps.GormSqlite)
		commands.RoomMember = gormSqliteCommand.NewRoomMember(deps.GormSqlite)
		commands.RoomMessage = gormSqliteCommand.NewRoomMessage(deps.GormSqlite)
		commands.FileInfo = gormSqliteCommand.NewFileInfo(deps.GormSqlite)
	}

	if deps.BoilerSqlite != nil {
		commands.User = boilerSqliteCommand.NewUser(deps.BoilerSqlite)
		commands.Joke = boilerSqliteCommand.NewJoke(deps.BoilerSqlite)
		commands.Like = boilerSqliteCommand.NewLike(deps.BoilerSqlite)
		commands.Comment = boilerSqliteCommand.NewComment(deps.BoilerSqlite)
		commands.Room = boilerSqliteCommand.NewRoom(deps.BoilerSqlite)
		commands.RoomMember = boilerSqliteCommand.NewRoomMember(deps.BoilerSqlite)
		commands.RoomMessage = boilerSqliteCommand.NewRoomMessage(deps.BoilerSqlite)
		commands.FileInfo = boilerSqliteCommand.NewFileInfo(deps.BoilerSqlite)
	}

	if deps.LocalFileSaver != "" {
		commands.FileContent = localFileSaverCommand.New(deps.LocalFileSaver)
	}

	// Check if all commands are initialized
	if err := coderr.CheckIfStructHasEmptyFields(commands); err != nil {
		return persistence.Commands{}, coderr.NewInternalErr(err)
	}

	return commands, nil
}
