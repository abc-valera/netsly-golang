package domain

import (
	"github.com/abc-valera/netsly-api-golang/pkg/core/validator"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/entity"
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
	commands Commands,
	queries Queries,
	services Services,
) Entities {
	validator := validator.NewValidator()

	return Entities{
		User:        entity.NewUser(commands.User, queries.User, validator, services.PasswordMaker),
		Joke:        entity.NewJoke(commands.Joke, queries.Joke, validator),
		Like:        entity.NewLike(commands.Like, queries.Like, validator),
		Comment:     entity.NewComment(commands.Comment, queries.Comment, validator),
		Room:        entity.NewRoom(commands.Room, queries.Room, validator),
		RoomMember:  entity.NewRoomMember(commands.RoomMember, queries.RoomMember, validator),
		RoomMessage: entity.NewRoomMessage(commands.RoomMessage, queries.RoomMessage, validator),
	}
}
