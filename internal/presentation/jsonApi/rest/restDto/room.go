package restDto

import (
	"github.com/abc-valera/netsly-api-golang/gen/ogen"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
)

func NewRoomResponse(room model.Room) *ogen.Room {
	return &ogen.Room{
		ID:            room.ID,
		CreatorUserID: room.CreatorUserID,
		Name:          room.Name,
		CreatedAt:     room.CreatedAt,
	}
}

func NewRoomsResponse(rooms []model.Room) *ogen.Rooms {
	var ogenRooms []ogen.Room
	for _, room := range rooms {
		ogenRooms = append(ogenRooms, *NewRoomResponse(room))
	}
	return &ogen.Rooms{Rooms: ogenRooms}
}
