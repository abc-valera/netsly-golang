package handler

import (
	"encoding/json"
	"time"

	"github.com/abc-valera/netsly-api-golang/internal/domain/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-api-golang/internal/presentation/json-api/ws/client"
	"github.com/abc-valera/netsly-api-golang/internal/presentation/json-api/ws/event"
)

const (
	EventTypeRoomMessage event.Type = "room_message"
)

type RoomMessage struct {
	userID  string
	clients *client.Clients

	roomMemberQuery query.IRoomMember
}

func NewRoomMessage(
	userID string,
	clients *client.Clients,

	roomMemberQuery query.IRoomMember,
) RoomMessage {
	return RoomMessage{
		userID:  userID,
		clients: clients,

		roomMemberQuery: roomMemberQuery,
	}
}

// ReceiveRoomMessage is the payload which will be received from the client
type receiveRoomMessagePayload struct {
	RoomID  string `json:"room_id"`
	Content string `json:"content"`
}

// SendRoomMessage is the payload which will be sent to the client
type sendRoomMessagePayload struct {
	FromID  string    `json:"from_id"`
	RoomID  string    `json:"room_id"`
	Content string    `json:"content"`
	SentAt  time.Time `json:"sent_at"`
}

func (h RoomMessage) RoomMessageHandler(e event.Event) error {
	var receive receiveRoomMessagePayload
	if err := json.Unmarshal(e.Payload, &receive); err != nil {
		return coderr.NewCodeMessageError(coderr.CodeInvalidArgument, "Json input error:", err)
	}

	send, err := json.Marshal(sendRoomMessagePayload{
		FromID:  h.userID,
		RoomID:  receive.RoomID,
		Content: receive.Content,
		SentAt:  time.Now(),
	})
	if err != nil {
		return coderr.NewInternalErr(err)
	}

	for _, client := range h.clients.GetAll() {
		client.Write() <- event.Event{
			Type:    EventTypeRoomMessage,
			Payload: send,
		}
	}

	return nil
}
