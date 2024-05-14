package restDto

import (
	"github.com/abc-valera/netsly-api-golang/gen/ogen"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
)

func NewRoomMessageResponse(roomMessage model.RoomMessage) *ogen.RoomMessage {
	return &ogen.RoomMessage{
		ID:        roomMessage.ID,
		RoomID:    roomMessage.RoomID,
		UserID:    roomMessage.UserID,
		Text:      roomMessage.Text,
		CreatedAt: roomMessage.CreatedAt,
	}
}

func NewRoomMessagesResponse(roomMessages []model.RoomMessage) *ogen.RoomMessages {
	var ogenRoomMessages []ogen.RoomMessage
	for _, roomMessage := range roomMessages {
		ogenRoomMessages = append(ogenRoomMessages, *NewRoomMessageResponse(roomMessage))
	}
	return &ogen.RoomMessages{RoomMessages: ogenRoomMessages}
}
