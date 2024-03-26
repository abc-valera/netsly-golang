package ws

import (
	"github.com/abc-valera/netsly-api-golang/pkg/core/coderr"
	"github.com/abc-valera/netsly-api-golang/pkg/presentation/jsonApi/ws/event"
	"github.com/abc-valera/netsly-api-golang/pkg/presentation/jsonApi/ws/handler"
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
