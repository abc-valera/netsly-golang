package boilerDto

import (
	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/boiler/errutil"
)

func NewDomainRoom(room *sqlboiler.Room) model.Room {
	if room == nil {
		return model.Room{}
	}

	return model.Room{
		ID:            room.ID,
		Name:          room.Name,
		Description:   room.Description,
		CreatorUserID: room.CreatorUserID,
		CreatedAt:     room.CreatedAt,
	}
}

func NewDomainRoomWithErrHandle(room *sqlboiler.Room, err error) (model.Room, error) {
	return NewDomainRoom(room), errutil.HandleErr(err)
}

func NewDomainRooms(rooms sqlboiler.RoomSlice) model.Rooms {
	var domainRooms model.Rooms
	for _, room := range rooms {
		domainRooms = append(domainRooms, NewDomainRoom(room))
	}
	return domainRooms
}

func NewDomainRoomsWithErrHandle(rooms sqlboiler.RoomSlice, err error) (model.Rooms, error) {
	return NewDomainRooms(rooms), errutil.HandleErr(err)
}
