package commands

import (
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-golang/internal/domain/util/coderr"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/commands/bunSqliteCommand"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/commands/localFileSaverCommand"
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
		coderr.Fatal("GORM SQLITE QUERIES NOT IMPLEMENTED")
	}

	if dep := dep.BunSqlite; dep != nil {
		commands.User = bunSqliteCommand.NewUser(dep)
		commands.Joke = bunSqliteCommand.NewJoke(dep)
		commands.Like = bunSqliteCommand.NewLike(dep)
		commands.Comment = bunSqliteCommand.NewComment(dep)
		commands.Room = bunSqliteCommand.NewRoom(dep)
		commands.RoomMember = bunSqliteCommand.NewRoomMember(dep)
		commands.RoomMessage = bunSqliteCommand.NewRoomMessage(dep)
		commands.FileInfo = bunSqliteCommand.NewFileInfo(dep)
	}

	if dep.LocalFileSaver != "" {
		commands.FileContent = localFileSaverCommand.New(dep.LocalFileSaver)
	}

	return commands
}
