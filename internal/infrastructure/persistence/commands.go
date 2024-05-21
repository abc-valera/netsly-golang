package persistence

import (
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/boiler/boilerCommand"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func NewCommands(executor boil.ContextExecutor) persistence.Commands {
	return persistence.Commands{
		User:        boilerCommand.NewUser(executor),
		Joke:        boilerCommand.NewJoke(executor),
		Like:        boilerCommand.NewLike(executor),
		Comment:     boilerCommand.NewComment(executor),
		Room:        boilerCommand.NewRoom(executor),
		RoomMember:  boilerCommand.NewRoomMember(executor),
		RoomMessage: boilerCommand.NewRoomMessage(executor),
	}
}
