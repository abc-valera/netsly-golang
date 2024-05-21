package domain

import (
	"github.com/abc-valera/netsly-api-golang/internal/domain/entity"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence"
)

type Entities struct {
	User        entity.IUser
	Joke        entity.IJoke
	Like        entity.ILike
	Comment     entity.IComment
	Room        entity.IRoom
	RoomMember  entity.IRoomMember
	RoomMessage entity.IRoomMessage
}

func NewEntities(
	commands persistence.Commands,
	queries persistence.Queries,
	services Services,
) Entities {
	return Entities{
		User:        entity.NewUser(commands.User, queries.User, services.PasswordMaker),
		Joke:        entity.NewJoke(commands.Joke, queries.Joke),
		Like:        entity.NewLike(commands.Like, queries.Like),
		Comment:     entity.NewComment(commands.Comment, queries.Comment),
		Room:        entity.NewRoom(commands.Room, queries.Room),
		RoomMember:  entity.NewRoomMember(commands.RoomMember, queries.RoomMember),
		RoomMessage: entity.NewRoomMessage(commands.RoomMessage, queries.RoomMessage),
	}
}
