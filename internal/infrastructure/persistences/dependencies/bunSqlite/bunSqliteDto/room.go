package bunSqliteDto

import (
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/uptrace/bun"
)

type Room struct {
	bun.BaseModel

	ID          string    `bun:"id,pk,type:uuid"`
	Name        string    `bun:",unique,notnull"`
	Description string    `bun:",notnull"`
	CreatedAt   time.Time `bun:",notnull"`
	UpdatedAt   time.Time `bun:",notnull"`
	DeletedAt   time.Time `bun:",notnull"`

	CreatorUserID string       `bun:",notnull"`
	Members       RoomMembers  `bun:"rel:has-many,join:id=room_id"`
	Messages      RoomMessages `bun:"rel:has-many,join:id=room_id"`
}

func NewRoom(room model.Room) Room {
	return Room{
		ID:          room.ID,
		Name:        room.Name,
		Description: room.Description,
		CreatedAt:   room.CreatedAt,
		UpdatedAt:   room.UpdatedAt,
		DeletedAt:   room.DeletedAt,

		CreatorUserID: room.CreatorUserID,
	}
}

func (dto Room) ToDomain() model.Room {
	return model.Room{
		ID:          dto.ID,
		Name:        dto.Name,
		Description: dto.Description,
		CreatedAt:   dto.CreatedAt,
		UpdatedAt:   dto.UpdatedAt,
		DeletedAt:   dto.DeletedAt,

		CreatorUserID: dto.CreatorUserID,
	}
}

type Rooms []Room

func NewRooms(rooms model.Rooms) Rooms {
	dtos := make(Rooms, 0, len(rooms))
	for _, room := range rooms {
		dtos = append(dtos, NewRoom(room))
	}
	return dtos
}

func (dtos Rooms) ToDomain() model.Rooms {
	rooms := make(model.Rooms, 0, len(dtos))
	for _, room := range dtos {
		rooms = append(rooms, room.ToDomain())
	}
	return rooms
}
