package handler

import (
	"encoding/json"

	"github.com/abc-valera/netsly-api-golang/internal/presentation/jsonApi/ws/client"
	"github.com/abc-valera/netsly-api-golang/internal/presentation/jsonApi/ws/event"
)

const (
	EventTypeError event.Type = "error"
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

type sendErrorPayload struct {
	ErrorMessage string `json:"error_message"`
}

func (h Error) HandleError(err error, currentClient client.Client) {
	send, _ := json.Marshal(sendErrorPayload{
		ErrorMessage: err.Error(),
	})

	currentClient.Write() <- event.Event{
		Type:    EventTypeError,
		Payload: send,
	}
}
