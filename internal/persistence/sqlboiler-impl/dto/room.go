package dto

import (
	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model/common"
	"github.com/abc-valera/netsly-api-golang/internal/persistence/sqlboiler-impl/errors"
)

func ToDomainRoom(room *sqlboiler.Room) model.Room {
	if room == nil {
		return model.Room{}
	}

	return model.Room{
		BaseModel: common.BaseModel{
			ID:        room.ID,
			CreatedAt: room.CreatedAt,
		},
		Name:        room.Name,
		Description: room.Description.String,
		CreatorID:   room.CreatorID,
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
