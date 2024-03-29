package dto

import (
	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/model"
	"github.com/abc-valera/netsly-api-golang/pkg/infrastructure/persistence/boiler/errors"
)

func ToDomainRoom(room *sqlboiler.Room) model.Room {
	if room == nil {
		return model.Room{}
	}

	return model.Room{
		ID:            room.ID,
		Name:          room.Name,
		Description:   room.Description.String,
		CreatedAt:     room.CreatedAt,
		CreatorUserID: room.CreatorUserID,
	}
}

func ToDomainRoomWithErrHandle(room *sqlboiler.Room, err error) (model.Room, error) {
	return ToDomainRoom(room), errors.HandleErr(err)
}

func ToDomainRooms(rooms sqlboiler.RoomSlice) model.Rooms {
	var domainRooms model.Rooms
	for _, room := range rooms {
		domainRooms = append(domainRooms, ToDomainRoom(room))
	}
	return domainRooms
}

func ToDomainRoomsWithErrHandle(rooms sqlboiler.RoomSlice, err error) (model.Rooms, error) {
	return ToDomainRooms(rooms), errors.HandleErr(err)
}
