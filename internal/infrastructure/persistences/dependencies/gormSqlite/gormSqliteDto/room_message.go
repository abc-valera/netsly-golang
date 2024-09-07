package gormSqliteDto

import (
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/command"
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

func NewRoomMessageUpdate(roomMessage RoomMessage, req command.RoomMessageUpdateRequest) RoomMessage {
	roomMessage.UpdatedAt = req.UpdatedAt

	if req.Text != nil {
		roomMessage.Text = *req.Text
	}

	return roomMessage
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

func NewRoomMessages(domainRoomMessages model.RoomMessages) RoomMessages {
	var roomMessages RoomMessages
	for _, domainRoomMessage := range domainRoomMessages {
		roomMessages = append(roomMessages, NewRoomMessage(domainRoomMessage))
	}
	return roomMessages
}

func (dtos RoomMessages) ToDomain() model.RoomMessages {
	var domainRoomMessages model.RoomMessages
	for _, dto := range dtos {
		domainRoomMessages = append(domainRoomMessages, dto.ToDomain())
	}
	return domainRoomMessages
}
