package bunSqliteDto

import (
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/uptrace/bun"
)

type RoomMessage struct {
	bun.BaseModel

	ID        string    `bun:"id,pk,type:uuid"`
	Text      string    `bun:",notnull"`
	CreatedAt time.Time `bun:",notnull"`
	UpdatedAt time.Time `bun:",notnull"`
	DeletedAt time.Time `bun:",notnull"`

	UserID string `bun:",notnull"`
	RoomID string `bun:",notnull"`
}

func NewRoomMessage(roomMessage model.RoomMessage) RoomMessage {
	return RoomMessage{
		ID:        roomMessage.ID,
		Text:      roomMessage.Text,
		CreatedAt: roomMessage.CreatedAt,
		UpdatedAt: roomMessage.UpdatedAt,
		DeletedAt: roomMessage.DeletedAt,

		UserID: roomMessage.UserID,
		RoomID: roomMessage.RoomID,
	}
}

func (dto RoomMessage) ToDomain() model.RoomMessage {
	return model.RoomMessage{
		ID:        dto.ID,
		Text:      dto.Text,
		CreatedAt: dto.CreatedAt,
		UpdatedAt: dto.UpdatedAt,
		DeletedAt: dto.DeletedAt,

		UserID: dto.UserID,
		RoomID: dto.RoomID,
	}
}

type RoomMessages []RoomMessage

func NewRoomMessages(roomMessages model.RoomMessages) RoomMessages {
	dtos := make(RoomMessages, 0, len(roomMessages))
	for _, roomMessage := range roomMessages {
		dtos = append(dtos, NewRoomMessage(roomMessage))
	}
	return dtos
}

func (dtos RoomMessages) ToDomain() model.RoomMessages {
	messages := make(model.RoomMessages, 0, len(dtos))
	for _, message := range dtos {
		messages = append(messages, message.ToDomain())
	}
	return messages
}
