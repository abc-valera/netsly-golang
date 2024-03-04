package ws

import (
	"github.com/abc-valera/netsly-api-golang/internal/domain/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/presentation/json-api/ws/event"
	"github.com/abc-valera/netsly-api-golang/internal/presentation/json-api/ws/handler"
)

func routeEvent(
	e event.Event,
	errorHandler handler.Error,
	roomHandler handler.Room,
) error {
	switch e.Type {
	case handler.EventTypeInvalidArgument:
		return errorHandler.InvalidArgumentHandler(e)
	case handler.EventTypeRoomMessage:
		return roomHandler.SendRoomMessageHandler(e)
	default:
		return coderr.NewMessage(coderr.CodeInvalidArgument, "Invalid event type")
	}
}
