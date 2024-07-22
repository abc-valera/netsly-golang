package domain

import (
	"github.com/abc-valera/netsly-api-golang/internal/domain/entity"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/commandTransactor"
)

type Entities struct {
	User        entity.IUser
	Joke        entity.IJoke
	Like        entity.ILike
	Comment     entity.IComment
	Room        entity.IRoom
	RoomMember  entity.IRoomMember
	RoomMessage entity.IRoomMessage
	File        entity.IFile
}

func NewEntities(
	commands persistence.Commands,
	commandTransactor commandTransactor.ITransactor,
	queries persistence.Queries,
	services Services,
) Entities {
	return Entities{
		User: entity.NewUser(
			commands.User,
			queries.User,
			services.Passworder,
		),
		Joke: entity.NewJoke(
			commands.Joke,
			queries.Joke,
		),
		Like: entity.NewLike(
			commands.Like,
			queries.Like,
		),
		Comment: entity.NewComment(
			commands.Comment,
			queries.Comment,
		),
		Room: entity.NewRoom(
			commands.Room,
			queries.Room,
			commandTransactor,
		),
		RoomMember: entity.NewRoomMember(
			commands.RoomMember,
			queries.RoomMember,
		),
		RoomMessage: entity.NewRoomMessage(
			commands.RoomMessage,
			queries.RoomMessage,
		),
		File: entity.NewFile(
			commands.FileInfo,
			queries.FileInfo,
			commands.FileContent,
			queries.FileContent,
			commandTransactor,
		),
	}
}
