package gormSqliteDto

import (
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
)

type RoomMessage struct {
	ID        string    `gorm:"primaryKey;not null"`
	Text      string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
	DeletedAt time.Time `gorm:"not null"`
	UserID    string    `gorm:"not null"`
	RoomID    string    `gorm:"not null"`
}

func NewRoomMessage(roomMessage model.RoomMessage) RoomMessage {
	return RoomMessage{
		ID:        roomMessage.ID,
		Text:      roomMessage.Text,
		CreatedAt: roomMessage.CreatedAt,
		UpdatedAt: roomMessage.UpdatedAt,
		DeletedAt: roomMessage.DeletedAt,
		UserID:    roomMessage.UserID,
		RoomID:    roomMessage.RoomID,
	}
}

func (dto RoomMessage) ToDomain() model.RoomMessage {
	return model.RoomMessage{
		ID:        dto.ID,
		Text:      dto.Text,
		CreatedAt: dto.CreatedAt,
		UpdatedAt: dto.UpdatedAt,
		DeletedAt: dto.DeletedAt,
		UserID:    dto.UserID,
		RoomID:    dto.RoomID,
	}
}

type RoomMessages []RoomMessage

func NewRoomMessages(domainRoomMessages []model.RoomMessage) RoomMessages {
	var roomMessages RoomMessages
	for _, domainRoomMessage := range domainRoomMessages {
		roomMessages = append(roomMessages, NewRoomMessage(domainRoomMessage))
	}
	return roomMessages
}

func (dtos RoomMessages) ToDomain() []model.RoomMessage {
	var domainRoomMessages []model.RoomMessage
	for _, dto := range dtos {
		domainRoomMessages = append(domainRoomMessages, dto.ToDomain())
	}
	return domainRoomMessages
}
