package handler

import (
	"encoding/json"
	"fmt"

	"github.com/abc-valera/netsly-api-golang/internal/presentation/json-rest-api/ws/client"
	"github.com/abc-valera/netsly-api-golang/internal/presentation/json-rest-api/ws/event"
)

const (
	EventTypeInvalidArgument event.Type = "invalid_argument"
)

type Error struct {
	userID  string
	clients *client.Clients
}

func NewError(
	userID string,
	clients *client.Clients,
) Error {
	return Error{
		userID:  userID,
		clients: clients,
	}
}

type invalidArgumentPayload struct {
	Message string `json:"message"`
}

func (h Error) InvalidArgumentHandler(e event.Event) error {
	var payload invalidArgumentPayload
	if err := json.Unmarshal(e.Payload, &payload); err != nil {
		return err
	}

	fmt.Println(payload.Message)

	return nil
}
