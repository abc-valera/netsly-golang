package bunSqliteDto

import (
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
)

type RoomMessage struct {
	ID        string    `bun:",pk,type:uuid"`
	Text      string    `bun:",notnull"`
	CreatedAt time.Time `bun:",notnull"`
	UpdatedAt time.Time `bun:",notnull"`
	DeletedAt time.Time `bun:",notnull"`

	UserID string `bun:",notnull"`
	RoomID string `bun:",notnull"`
}

func (m RoomMessage) ToDomain() model.RoomMessage {
	return model.RoomMessage{
		ID:        m.ID,
		Text:      m.Text,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		DeletedAt: m.DeletedAt,
	}
}

type RoomMessages []RoomMessage

func (m RoomMessages) ToDomain() model.RoomMessages {
	messages := make(model.RoomMessages, 0, len(m))
	for _, message := range m {
		messages = append(messages, message.ToDomain())
	}
	return messages
}
