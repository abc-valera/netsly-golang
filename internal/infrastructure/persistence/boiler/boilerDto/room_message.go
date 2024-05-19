package boilerDto

import (
	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/boiler/errors"
)

func NewDomainRoomMessage(roomMessage *sqlboiler.RoomMessage) model.RoomMessage {
	if roomMessage == nil {
		return model.RoomMessage{}
	}

	return model.RoomMessage{
		ID:        roomMessage.ID,
		Text:      roomMessage.Text,
		CreatedAt: roomMessage.CreatedAt,
		UserID:    roomMessage.UserID,
		RoomID:    roomMessage.RoomID,
	}
}

func NewDomainRoomMessageWithErrHandle(roomMessage *sqlboiler.RoomMessage, err error) (model.RoomMessage, error) {
	return NewDomainRoomMessage(roomMessage), errors.HandleErr(err)
}

func NewDomainRoomMessages(roomMessages sqlboiler.RoomMessageSlice) model.RoomMessages {
	var domainRoomMessages model.RoomMessages
	for _, roomMessage := range roomMessages {
		domainRoomMessages = append(domainRoomMessages, NewDomainRoomMessage(roomMessage))
	}
	return domainRoomMessages
}

func NewDomainRoomMessagesWithErrHandle(roomMessages sqlboiler.RoomMessageSlice, err error) (model.RoomMessages, error) {
	return NewDomainRoomMessages(roomMessages), errors.HandleErr(err)
}
