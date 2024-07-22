package gormSqliteDto

import (
	"time"

	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
)

type RoomMessage struct {
	ID        string    `gorm:"primaryKey;not null"`
	Text      string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
	DeletedAt time.Time `gorm:"not null"`

	UserID string `gorm:"not null"`
	RoomID string `gorm:"not null"`
}

func NewDomainRoomMessage(roomMessage RoomMessage) model.RoomMessage {
	return model.RoomMessage{
		ID:        roomMessage.ID,
		Text:      roomMessage.Text,
		CreatedAt: roomMessage.CreatedAt,
		UpdatedAt: roomMessage.UpdatedAt,
		DeletedAt: roomMessage.DeletedAt,
	}
}

type RoomMessages []RoomMessage

func NewDomainRoomMessages(roomMessages RoomMessages) model.RoomMessages {
	var domainRoomMessages model.RoomMessages
	for _, roomMessage := range roomMessages {
		domainRoomMessages = append(domainRoomMessages, NewDomainRoomMessage(roomMessage))
	}
	return domainRoomMessages
}
