package bunSqliteDto

import (
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
)

type Room struct {
	ID          string    `bun:",pk,type:uuid"`
	Name        string    `bun:",unique,notnull"`
	Description string    `bun:",notnull"`
	CreatedAt   time.Time `bun:",notnull"`
	UpdatedAt   time.Time `bun:",notnull"`
	DeletedAt   time.Time `bun:",notnull"`

	CreatorUserID string       `bun:",notnull"`
	Members       RoomMembers  `bun:"rel:has-many,join:id=room_id"`
	Messages      RoomMessages `bun:"rel:has-many,join:id=room_id"`
}

func (r Room) ToDomain() model.Room {
	return model.Room{
		ID:          r.ID,
		Name:        r.Name,
		Description: r.Description,
		CreatedAt:   r.CreatedAt,
		UpdatedAt:   r.UpdatedAt,
		DeletedAt:   r.DeletedAt,
	}
}

type Rooms []Room

func (r Rooms) ToDomain() model.Rooms {
	rooms := make(model.Rooms, 0, len(r))
	for _, room := range r {
		rooms = append(rooms, room.ToDomain())
	}
	return rooms
}
