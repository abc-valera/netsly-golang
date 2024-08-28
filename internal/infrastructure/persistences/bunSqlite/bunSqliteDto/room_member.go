package bunSqliteDto

import (
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
)

// RoomMember is a struct that represents a chat member entity.
// Technically, it's a many-to-many relationship between Room and User entities.
type RoomMember struct {
	CreatedAt time.Time `bun:",notnull"`
	DeletedAt time.Time `bun:",notnull"`

	UserID   string       `bun:",pk,notnull"`
	RoomID   string       `bun:",pk,notnull"`
	Messages RoomMessages `bun:"rel:has-many,join:user_id=user_id,join:room_id=room_id"`
}

func (m RoomMember) ToDomain() model.RoomMember {
	return model.RoomMember{
		CreatedAt: m.CreatedAt,
		DeletedAt: m.DeletedAt,
	}
}

type RoomMembers []RoomMember

func (m RoomMembers) ToDomain() model.RoomMembers {
	members := make(model.RoomMembers, 0, len(m))
	for _, member := range m {
		members = append(members, member.ToDomain())
	}
	return members
}
