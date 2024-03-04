package entquery

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/ent"
	"github.com/abc-valera/netsly-api-golang/gen/ent/room"
	"github.com/abc-valera/netsly-api-golang/gen/ent/roommessage"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query/spec"
	"github.com/abc-valera/netsly-api-golang/internal/persistence/ent-impl/dto"
)

type roomMessageQuery struct {
	*ent.Client
}

func NewRoomMessageQuery(client *ent.Client) query.IRoomMessage {
	return &roomMessageQuery{
		Client: client,
	}
}

func (cm roomMessageQuery) GetByID(ctx context.Context, id string) (model.RoomMessage, error) {
	return dto.FromEntRoomMessageWithErrHandle(cm.RoomMessage.Get(ctx, id))
}

func (cm roomMessageQuery) GetAllByRoomID(ctx context.Context, roomID string, params spec.SelectParams) (model.RoomMessages, error) {
	query := cm.RoomMessage.
		Query().
		Where(roommessage.HasRoomWith(room.ID(roomID)))

	if params.Order() == "asc" {
		query = query.Order(ent.Asc("created_at"))
	} else {
		query = query.Order(ent.Desc("created_at"))
	}

	query.Limit(params.Limit())
	query.Offset(params.Offset())

	return dto.FromEntRoomMessagesWithErrHandle(query.All(ctx))
}

func (cm roomMessageQuery) SearchAllByText(ctx context.Context, keyword string, params spec.SelectParams) (model.RoomMessages, error) {
	query := cm.RoomMessage.
		Query().
		Where(roommessage.TextContains(keyword))

	if params.Order() == "asc" {
		query = query.Order(ent.Asc("created_at"))
	} else {
		query = query.Order(ent.Desc("created_at"))
	}

	query.Limit(params.Limit())
	query.Offset(params.Offset())

	return dto.FromEntRoomMessagesWithErrHandle(query.All(ctx))
}