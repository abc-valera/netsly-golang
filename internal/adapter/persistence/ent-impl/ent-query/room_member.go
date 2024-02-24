package entquery

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/ent"
	"github.com/abc-valera/netsly-api-golang/gen/ent/room"
	"github.com/abc-valera/netsly-api-golang/gen/ent/roommember"
	"github.com/abc-valera/netsly-api-golang/gen/ent/user"
	"github.com/abc-valera/netsly-api-golang/internal/adapter/persistence/ent-impl/dto"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query"
)

type roomMemberQuery struct {
	*ent.Client
}

func NewRoomMemberQuery(client *ent.Client) query.IRoomMember {
	return &roomMemberQuery{
		Client: client,
	}
}

func (cq roomMemberQuery) GetByIDs(ctx context.Context, userID string, roomID string) (model.RoomMember, error) {
	return dto.FromEntRoomMemberWithErrHandle(cq.RoomMember.Query().Where(
		roommember.HasUserWith(user.ID(userID)),
		roommember.HasRoomWith(room.ID(roomID)),
	).Only(ctx))
}
