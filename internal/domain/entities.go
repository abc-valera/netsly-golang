package domain

import "github.com/abc-valera/netsly-api-golang/internal/domain/entity"

type Entities struct {
	User        entity.User
	Joke        entity.Joke
	Like        entity.Like
	Comment     entity.Comment
	Room        entity.Room
	RoomMember  entity.RoomMember
	RoomMessage entity.RoomMessage
}

func NewEntities(
	commands Commands,
	queries Queries,
	services Services,
) Entities {
	return Entities{
		User:        entity.NewUser(commands.User, queries.User, services.UUUIDMaker, services.Time, services.PasswordMaker),
		Joke:        entity.NewJoke(commands.Joke, services.UUUIDMaker, services.Time),
		Like:        entity.NewLike(commands.Like, services.Time),
		Comment:     entity.NewComment(commands.Comment, services.UUUIDMaker, services.Time),
		Room:        entity.NewRoom(commands.Room, services.UUUIDMaker, services.Time),
		RoomMember:  entity.NewRoomMember(commands.RoomMember, services.Time),
		RoomMessage: entity.NewRoomMessage(commands.RoomMessage, services.UUUIDMaker, services.Time),
	}
}
