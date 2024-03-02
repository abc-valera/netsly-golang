package handler

import (
	"encoding/json"
	"time"

	"github.com/abc-valera/netsly-api-golang/internal/domain/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-api-golang/internal/port/json-rest-api/ws/client"
	"github.com/abc-valera/netsly-api-golang/internal/port/json-rest-api/ws/event"
)

const (
	EventTypeRoomMessage event.Type = "send_room_message"
)

type Room struct {
	userID  string
	clients *client.Clients

	roomMemberQuery query.IRoomMember
}

func NewRoom(
	userID string,
	clients *client.Clients,

	roomMemberQuery query.IRoomMember,
) Room {
	return Room{
		userID:  userID,
		clients: clients,

		roomMemberQuery: roomMemberQuery,
	}
}

type sendRoomMessagePayload struct {
	FromID  string    `json:"from_id"`
	RoomID  string    `json:"room_id"`
	Content string    `json:"content"`
	SentAt  time.Time `json:"sent_at"`
}

func (h Room) SendRoomMessageHandler(e event.Event) error {
	var payload sendRoomMessagePayload
	if err := json.Unmarshal(e.Payload, &payload); err != nil {
		return coderr.NewInternal(err)
	}
	payload.FromID = h.userID

	// Send message to all clients in the clients map
	for _, client := range h.clients.GetAllByUserID(payload.FromID) {
		client.Write() <- event.Event{
			Type:    EventTypeRoomMessage,
			Payload: json.RawMessage(e.Payload),
		}
	}

	// roomMembers, err := h.roomMemberQuery.GetByRoomID(context.Background(), payload.RoomID)
	// if err != nil {
	// 	return err
	// }

	// for _, member := range roomMembers {
	// 	for _, client := range h.clients.GetAllByUserID(member.UserID) {
	// 		client.Write() <- event.Event{
	// 			Type:    EventTypeSendRoomMessage,
	// 			Payload: json.RawMessage(e.Payload),
	// 		}
	// 	}
	// }

	return nil
}
