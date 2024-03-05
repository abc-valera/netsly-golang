package ws

import (
	"github.com/abc-valera/netsly-api-golang/internal/domain/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/presentation/json-api/ws/event"
	"github.com/abc-valera/netsly-api-golang/internal/presentation/json-api/ws/handler"
)

func routeEvent(
	e event.Event,
	roomHandler handler.RoomMessage,
) error {
	switch e.Type {
	case handler.EventTypeRoomMessage:
		return roomHandler.RoomMessageHandler(e)
	default:
		return coderr.NewCodeMessage(coderr.CodeInvalidArgument, "Invalid event type")
	}
}
