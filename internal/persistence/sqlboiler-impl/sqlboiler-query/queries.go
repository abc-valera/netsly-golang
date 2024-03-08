package sqlboilerquery

import (
	"github.com/abc-valera/netsly-api-golang/internal/domain"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func NewQueries(executor boil.ContextExecutor) domain.Queries {
	return domain.NewQueries(
		newUser(executor),
		newJoke(executor),
		newLike(executor),
		newComment(executor),
		newRoom(executor),
		newRoomMember(executor),
		newRoomMessage(executor),
	)
}
