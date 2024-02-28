package ws

import (
	"github.com/abc-valera/netsly-api-golang/internal/domain/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/port/json-rest-api/ws/event"
	"github.com/abc-valera/netsly-api-golang/internal/port/json-rest-api/ws/handler"
	"github.com/gorilla/websocket"
)

func routeEvent(e event.Event, conn *websocket.Conn) error {
	switch e.Type {
	case handler.EventTypeSendRoomMessage:
		return handler.SendChatMsgHandler(e, conn)
	case handler.EventTypeInvalidArgument:
		return handler.InvalidArgumentHandler(e, conn)
	default:
		return coderr.NewMessage(coderr.CodeInvalidArgument, "Invalid event type")
	}
}
