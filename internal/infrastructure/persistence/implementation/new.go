package implementation

import (
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/boilerSqlite/boilerSqliteCommand"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/boilerSqlite/boilerSqliteQuery"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/gormSqlite/gormSqliteCommand"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/gormSqlite/gormSqliteQuery"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/localFileSaver/localFileSaverCommand"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/localFileSaver/localFileSaverQuery"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"gorm.io/gorm"
)

type CommandsAndQueriesDependencies struct {
	GormSqlite     *gorm.DB
	BoilerSqlite   boil.ContextExecutor
	LocalFileSaver string
}

func NewCommandsAndQueries(deps CommandsAndQueriesDependencies) (persistence.Commands, persistence.Queries, error) {
	// Init the commands and queries per available dependencies

	var commands persistence.Commands
	var queries persistence.Queries

	if deps.GormSqlite != nil {
		commands.User = gormSqliteCommand.NewUser(deps.GormSqlite)
		commands.Joke = gormSqliteCommand.NewJoke(deps.GormSqlite)
		commands.Like = gormSqliteCommand.NewLike(deps.GormSqlite)
		commands.Comment = gormSqliteCommand.NewComment(deps.GormSqlite)
		commands.Room = gormSqliteCommand.NewRoom(deps.GormSqlite)
		commands.RoomMember = gormSqliteCommand.NewRoomMember(deps.GormSqlite)
		commands.RoomMessage = gormSqliteCommand.NewRoomMessage(deps.GormSqlite)
		commands.FileInfo = gormSqliteCommand.NewFileInfo(deps.GormSqlite)

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
		commands.User = boilerSqliteCommand.NewUser(deps.BoilerSqlite)
		commands.Joke = boilerSqliteCommand.NewJoke(deps.BoilerSqlite)
		commands.Like = boilerSqliteCommand.NewLike(deps.BoilerSqlite)
		commands.Comment = boilerSqliteCommand.NewComment(deps.BoilerSqlite)
		commands.Room = boilerSqliteCommand.NewRoom(deps.BoilerSqlite)
		commands.RoomMember = boilerSqliteCommand.NewRoomMember(deps.BoilerSqlite)
		commands.RoomMessage = boilerSqliteCommand.NewRoomMessage(deps.BoilerSqlite)
		commands.FileInfo = boilerSqliteCommand.NewFileInfo(deps.BoilerSqlite)

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
		commands.FileContent = localFileSaverCommand.New(deps.LocalFileSaver)
		queries.FileContent = localFileSaverQuery.New(deps.LocalFileSaver)
	}

	// Check if all commands and queries were initialized via reflection

	if err := persistence.ValidateCommands(commands); err != nil {
		return persistence.Commands{}, persistence.Queries{}, err
	}

	if err := persistence.ValidateQueries(queries); err != nil {
		return persistence.Commands{}, persistence.Queries{}, err
	}

	return commands, queries, nil
}
