package sqlboilerCommand

import (
	"github.com/abc-valera/netsly-api-golang/internal/domain"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func NewCommands(executor boil.ContextExecutor) domain.Commands {
	return domain.NewCommands(
		newUser(executor),
		newJoke(executor),
		newLike(executor),
		newComment(executor),
		newRoom(executor),
		newRoomMember(executor),
		newRoomMessage(executor),
	)
}
