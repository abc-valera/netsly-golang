package handler

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/ogen"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/query"
	"github.com/abc-valera/netsly-api-golang/pkg/presentation/jsonApi/rest/restDto"
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
	domainMessages, err := h.roomMessageQuery.GetAllByRoomID(ctx, ogenParams.RoomID, restDto.NewDomainSelectParams(&ogenParams.SelectParams))
	return restDto.NewRoomMessagesResponse(domainMessages), err
}
