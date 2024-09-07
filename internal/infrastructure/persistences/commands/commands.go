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
		commands.User = gormSqliteCommand.NewCreateUpdateDelete(dep, gormSqliteDto.NewUser, gormSqliteDto.NewUserUpdate)
		commands.Joke = gormSqliteCommand.NewCreateUpdateDelete(dep, gormSqliteDto.NewJoke, gormSqliteDto.NewJokeUpdate)
		commands.Like = gormSqliteCommand.NewCreateDelete(dep, gormSqliteDto.NewLike)
		commands.Comment = gormSqliteCommand.NewCreateUpdateDelete(dep, gormSqliteDto.NewComment, gormSqliteDto.NewCommentUpdate)
		commands.Room = gormSqliteCommand.NewCreateUpdateDelete(dep, gormSqliteDto.NewRoom, gormSqliteDto.NewRoomUpdate)
		commands.RoomMember = gormSqliteCommand.NewCreateDelete(dep, gormSqliteDto.NewRoomMember)
		commands.RoomMessage = gormSqliteCommand.NewCreateUpdateDelete(dep, gormSqliteDto.NewRoomMessage, gormSqliteDto.NewRoomMessageUpdate)
		commands.FileInfo = gormSqliteCommand.NewCreateUpdateDelete(dep, gormSqliteDto.NewFileInfo, gormSqliteDto.NewFileInfoUpdate)
	}

	if dep := dep.BunSqlite; dep != nil {
		commands.User = bunSqliteCommand.NewCreateUpdateDelete(dep, bunSqliteDto.NewUser, bunSqliteDto.NewUserUpdate)
		commands.Joke = bunSqliteCommand.NewCreateUpdateDelete(dep, bunSqliteDto.NewJoke, bunSqliteDto.NewJokeUpdate)
		commands.Like = bunSqliteCommand.NewCreateDelete(dep, bunSqliteDto.NewLike)
		commands.Comment = bunSqliteCommand.NewCreateUpdateDelete(dep, bunSqliteDto.NewComment, bunSqliteDto.NewCommentUpdate)
		commands.Room = bunSqliteCommand.NewCreateUpdateDelete(dep, bunSqliteDto.NewRoom, bunSqliteDto.NewRoomUpdate)
		commands.RoomMember = bunSqliteCommand.NewCreateDelete(dep, bunSqliteDto.NewRoomMember)
		commands.RoomMessage = bunSqliteCommand.NewCreateUpdateDelete(dep, bunSqliteDto.NewRoomMessage, bunSqliteDto.NewRoomMessageUpdate)
		commands.FileInfo = bunSqliteCommand.NewCreateUpdateDelete(dep, bunSqliteDto.NewFileInfo, bunSqliteDto.NewFileInfoUpdate)
	}

	if dep.LocalFileSaver != "" {
		commands.FileContent = localFileSaverCommand.New(dep.LocalFileSaver)
	}

	return commands
}
