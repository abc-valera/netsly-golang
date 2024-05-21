package persistence

import (
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/boiler/boilerQuery"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func NewQueries(executor boil.ContextExecutor) persistence.Queries {
	return persistence.Queries{
		User:        boilerQuery.NewUser(executor),
		Joke:        boilerQuery.NewJoke(executor),
		Like:        boilerQuery.NewLike(executor),
		Comment:     boilerQuery.NewComment(executor),
		Room:        boilerQuery.NewRoom(executor),
		RoomMember:  boilerQuery.NewRoomMember(executor),
		RoomMessage: boilerQuery.NewRoomMessage(executor),
	}
}
