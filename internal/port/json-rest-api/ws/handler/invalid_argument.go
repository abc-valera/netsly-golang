package handler

import (
	"encoding/json"
	"fmt"

	"github.com/abc-valera/netsly-api-golang/internal/port/json-rest-api/ws/event"
	"github.com/gorilla/websocket"
)

const (
	EventTypeInvalidArgument event.Type = "invalid_argument"
)

type InvalidArgumentPayload struct {
	Message string `json:"message"`
}

func InvalidArgumentHandler(e event.Event, conn *websocket.Conn) error {
	var payload InvalidArgumentPayload
	if err := json.Unmarshal(e.Payload, &payload); err != nil {
		return err
	}

	fmt.Println(payload.Message)

	return nil
}
