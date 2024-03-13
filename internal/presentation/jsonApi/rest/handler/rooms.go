package handler

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/ogen"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-api-golang/internal/presentation/jsonApi/rest/dto"
)

type Rooms struct {
	roomMessageQuery query.IRoomMessage
}

func NewRooms(
	roomMessageQuery query.IRoomMessage,
) Rooms {
	return Rooms{
		roomMessageQuery: roomMessageQuery,
	}
}

func (h Rooms) MeRoomsIdMessagesGet(ctx context.Context, ogenParams ogen.MeRoomsIdMessagesGetParams) (*ogen.RoomMessages, error) {
	domainMessages, err := h.roomMessageQuery.GetAllByRoomID(ctx, ogenParams.RoomID, dto.NewDomainSelectParams(&ogenParams.SelectParams))
	return dto.NewRoomMessagesResponse(domainMessages), err
}
