package dto

import (
	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/model"
	"github.com/abc-valera/netsly-api-golang/pkg/infrastructure/persistence/boiler/errors"
)

func ToDomainRoomMessage(roomMessage *sqlboiler.RoomMessage) model.RoomMessage {
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

func ToDomainRoomMessageWithErrHandle(roomMessage *sqlboiler.RoomMessage, err error) (model.RoomMessage, error) {
	return ToDomainRoomMessage(roomMessage), errors.HandleErr(err)
}

func ToDomainRoomMessages(roomMessages sqlboiler.RoomMessageSlice) model.RoomMessages {
	var domainRoomMessages model.RoomMessages
	for _, roomMessage := range roomMessages {
		domainRoomMessages = append(domainRoomMessages, ToDomainRoomMessage(roomMessage))
	}
	return domainRoomMessages
}

func ToDomainRoomMessagesWithErrHandle(roomMessages sqlboiler.RoomMessageSlice, err error) (model.RoomMessages, error) {
	return ToDomainRoomMessages(roomMessages), errors.HandleErr(err)
}
