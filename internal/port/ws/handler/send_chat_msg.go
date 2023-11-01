package handler

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/abc-valera/flugo-api-golang/internal/port/ws/client"
)

const (
	// EventSendMessage is the event type for the send_chat_msg event
	EventTypeSendMessage = "send_chat_msg"
)

// SendChatMsgPayload is the payload for the send_chat_msg event
type SendChatMsgPayload struct {
	ChatRoomID string    `json:"chat_room_id"`
	From       string    `json:"from"`
	Message    string    `json:"message"`
	SentAt     time.Time `json:"sent_at"`
}

// SendChatMsgHandler is the handler for the send_chat_msg event
func SendChatMsgHandler(e client.Event, c *client.Client) error {
	var payload SendChatMsgPayload
	if err := json.Unmarshal(e.Payload, &payload); err != nil {
		return err
	}

	// Do something with the payload
	fmt.Println(payload)

	return nil
}
