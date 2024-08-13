package boilerSqliteDto

import (
	"github.com/abc-valera/netsly-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-golang/internal/domain/model"
)

func NewDomainRoom(room *sqlboiler.Room) model.Room {
	if room == nil {
		return model.Room{}
	}

	return model.Room{
		ID:          room.ID,
		Name:        room.Name,
		Description: room.Description,
		CreatedAt:   room.CreatedAt,
		UpdatedAt:   room.UpdatedAt,
		DeletedAt:   room.DeletedAt,
	}
}

func NewDomainRooms(rooms sqlboiler.RoomSlice) model.Rooms {
	var domainRooms model.Rooms
	for _, room := range rooms {
		domainRooms = append(domainRooms, NewDomainRoom(room))
	}
	return domainRooms
}
