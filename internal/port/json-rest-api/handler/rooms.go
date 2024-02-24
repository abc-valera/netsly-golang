package handler

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/ogen"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-api-golang/internal/port/json-rest-api/dto"
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
	params, err := dto.NewDomainSelectParams(&ogenParams.SelectParams)
	if err != nil {
		return nil, err
	}

	domainMessages, err := h.roomMessageQuery.GetAllByRoomID(ctx, ogenParams.RoomID, params)
	return dto.NewRoomMessagesResponse(domainMessages), err
}
