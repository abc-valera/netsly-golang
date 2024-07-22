package boilerSqliteDto

import (
	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
)

func NewDomainRoomMessage(roomMessage *sqlboiler.RoomMessage) model.RoomMessage {
	if roomMessage == nil {
		return model.RoomMessage{}
	}

	return model.RoomMessage{
		ID:        roomMessage.ID,
		Text:      roomMessage.Text,
		CreatedAt: roomMessage.CreatedAt,
		UpdatedAt: roomMessage.UpdatedAt,
		DeletedAt: roomMessage.DeletedAt,
	}
}

func NewDomainRoomMessages(roomMessages sqlboiler.RoomMessageSlice) model.RoomMessages {
	var domainRoomMessages model.RoomMessages
	for _, roomMessage := range roomMessages {
		domainRoomMessages = append(domainRoomMessages, NewDomainRoomMessage(roomMessage))
	}
	return domainRoomMessages
}
