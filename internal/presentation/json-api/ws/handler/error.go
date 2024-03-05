package handler

import (
	"encoding/json"

	"github.com/abc-valera/netsly-api-golang/internal/domain/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/presentation/json-api/ws/client"
	"github.com/abc-valera/netsly-api-golang/internal/presentation/json-api/ws/event"
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

func (h Error) HandleError(err error, currentClient client.Client) error {
	send, err := json.Marshal(sendErrorPayload{
		ErrorMessage: err.Error(),
	})
	if err != nil {
		return coderr.NewInternalErr(err)
	}

	currentClient.Write() <- event.Event{
		Type:    EventTypeError,
		Payload: send,
	}

	return nil
}
