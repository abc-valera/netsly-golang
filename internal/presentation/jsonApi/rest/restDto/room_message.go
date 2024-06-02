package restDto

import (
	"github.com/abc-valera/netsly-api-golang/gen/ogen"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
)

func NewRoomMessage(roomMessage model.RoomMessage) *ogen.RoomMessage {
	return &ogen.RoomMessage{
		ID:        roomMessage.ID,
		Text:      roomMessage.Text,
		CreatedAt: roomMessage.CreatedAt,
	}
}

func NewRoomMessages(roomMessages []model.RoomMessage) *ogen.RoomMessages {
	var ogenRoomMessages []ogen.RoomMessage
	for _, roomMessage := range roomMessages {
		ogenRoomMessages = append(ogenRoomMessages, *NewRoomMessage(roomMessage))
	}
	return &ogen.RoomMessages{RoomMessages: ogenRoomMessages}
}
