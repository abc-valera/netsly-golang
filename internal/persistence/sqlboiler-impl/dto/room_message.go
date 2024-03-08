package dto

import (
	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model/common"
	"github.com/abc-valera/netsly-api-golang/internal/persistence/sqlboiler-impl/errors"
)

func ToDomainRoomMessage(roomMessage *sqlboiler.RoomMessage) model.RoomMessage {
	if roomMessage == nil {
		return model.RoomMessage{}
	}

	return model.RoomMessage{
		BaseEntity: common.BaseEntity{
			ID:        roomMessage.ID,
			CreatedAt: roomMessage.CreatedAt,
		},
		Text:   roomMessage.Text,
		UserID: roomMessage.UserID,
		RoomID: roomMessage.RoomID,
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
