package commands

import (
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/commands/bunSqliteCommand"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/commands/gormSqliteCommand"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/commands/localFileSaverCommand"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/dependencies/bunSqlite/bunSqliteDto"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/dependencies/gormSqlite/gormSqliteDto"
	"github.com/uptrace/bun"
	"gorm.io/gorm"
)

type Dependency struct {
	GormSqlite     *gorm.DB
	BunSqlite      bun.IDB
	LocalFileSaver string
}

func New(dep Dependency) command.Commands {
	var commands command.Commands

	if dep := dep.GormSqlite; dep != nil {
		commands.User = gormSqliteCommand.New(dep, gormSqliteDto.NewUser)
		commands.Joke = gormSqliteCommand.New(dep, gormSqliteDto.NewJoke)
		commands.Like = gormSqliteCommand.New(dep, gormSqliteDto.NewLike)
		commands.Comment = gormSqliteCommand.New(dep, gormSqliteDto.NewComment)
		commands.Room = gormSqliteCommand.New(dep, gormSqliteDto.NewRoom)
		commands.RoomMember = gormSqliteCommand.New(dep, gormSqliteDto.NewRoomMember)
		commands.RoomMessage = gormSqliteCommand.New(dep, gormSqliteDto.NewRoomMessage)
		commands.FileInfo = gormSqliteCommand.New(dep, gormSqliteDto.NewFileInfo)
		commands.FileInfoJoke = gormSqliteCommand.New(dep, gormSqliteDto.NewFileInfoJoke)
		commands.FileInfoRoom = gormSqliteCommand.New(dep, gormSqliteDto.NewFileInfoRoom)
	}

	if dep := dep.BunSqlite; dep != nil {
		commands.User = bunSqliteCommand.New(dep, bunSqliteDto.NewUser)
		commands.Joke = bunSqliteCommand.New(dep, bunSqliteDto.NewJoke)
		commands.Like = bunSqliteCommand.New(dep, bunSqliteDto.NewLike)
		commands.Comment = bunSqliteCommand.New(dep, bunSqliteDto.NewComment)
		commands.Room = bunSqliteCommand.New(dep, bunSqliteDto.NewRoom)
		commands.RoomMember = bunSqliteCommand.New(dep, bunSqliteDto.NewRoomMember)
		commands.RoomMessage = bunSqliteCommand.New(dep, bunSqliteDto.NewRoomMessage)
		commands.FileInfo = bunSqliteCommand.New(dep, bunSqliteDto.NewFileInfo)
		commands.FileInfoJoke = bunSqliteCommand.New(dep, bunSqliteDto.NewFileInfoJoke)
		commands.FileInfoRoom = bunSqliteCommand.New(dep, bunSqliteDto.NewFileInfoRoom)
	}

	if dep.LocalFileSaver != "" {
		commands.FileContent = localFileSaverCommand.New(dep.LocalFileSaver)
	}

	return commands
}
