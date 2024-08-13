package restDto

import (
	"github.com/abc-valera/netsly-golang/gen/ogen"
	"github.com/abc-valera/netsly-golang/internal/domain/model"
)

func NewRoom(room model.Room) *ogen.Room {
	return &ogen.Room{
		ID:        room.ID,
		Name:      room.Name,
		CreatedAt: room.CreatedAt,
	}
}

func NewRooms(rooms []model.Room) *ogen.Rooms {
	var ogenRooms []ogen.Room
	for _, room := range rooms {
		ogenRooms = append(ogenRooms, *NewRoom(room))
	}
	return &ogen.Rooms{Rooms: ogenRooms}
}
