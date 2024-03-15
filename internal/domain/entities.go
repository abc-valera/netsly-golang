package domain

import "github.com/abc-valera/netsly-api-golang/internal/domain/entity"

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
	return Entities{
		User:        entity.NewUser(commands.User, queries.User, services.PasswordMaker),
		Joke:        entity.NewJoke(commands.Joke),
		Like:        entity.NewLike(commands.Like),
		Comment:     entity.NewComment(commands.Comment),
		Room:        entity.NewRoom(commands.Room),
		RoomMember:  entity.NewRoomMember(commands.RoomMember),
		RoomMessage: entity.NewRoomMessage(commands.RoomMessage),
	}
}
