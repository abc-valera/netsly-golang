package handler

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/abc-valera/netsly-api-golang/internal/domain/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/port/json-rest-api/ws/event"
	"github.com/gorilla/websocket"
)

const (
	EventTypeSendRoomMessage event.Type = "send_message"
)

type SendRoomMessagePayload struct {
	RoomID  string    `json:"room_id"`
	FromID  string    `json:"from"`
	Message string    `json:"message"`
	SentAt  time.Time `json:"sent_at"`
}

func SendChatMsgHandler(e event.Event, conn *websocket.Conn) error {
	var payload SendRoomMessagePayload
	if err := json.Unmarshal(e.Payload, &payload); err != nil {
		return coderr.NewInternal(err)
	}

	// TODO: Do something with the payload
	fmt.Println(payload)

	return nil
}
