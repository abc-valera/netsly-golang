package handler

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/ogen"
	"github.com/abc-valera/netsly-api-golang/internal/domain/entity"
	"github.com/abc-valera/netsly-api-golang/internal/presentation/jsonApi/rest/restDto"
)

type Rooms struct {
	roomMessage entity.IRoomMessage
}

func NewRooms(
	roomMessage entity.IRoomMessage,
) Rooms {
	return Rooms{
		roomMessage: roomMessage,
	}
}

func (h Rooms) MeRoomsIdMessagesGet(ctx context.Context, ogenParams ogen.MeRoomsIdMessagesGetParams) (*ogen.RoomMessages, error) {
	domainMessages, err := h.roomMessage.GetAllByRoomID(ctx, ogenParams.RoomID, restDto.NewDomainSelector(&ogenParams.Selector))
	return restDto.NewRoomMessages(domainMessages), err
}
