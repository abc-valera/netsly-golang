package dto

import (
	"github.com/abc-valera/netsly-api-golang/gen/ent"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model/common"
)

func FromEntRoom(entRoom *ent.Room) model.Room {
	if entRoom == nil {
		return model.Room{}
	}
	return model.Room{
		BaseEntity: common.BaseEntity{
			ID:        entRoom.ID,
			CreatedAt: entRoom.CreatedAt,
		},
		Name:        entRoom.Name,
		CreatorID:   entRoom.CreatorID,
		Description: entRoom.Description,
	}
}

func FromEntRoomWithErrHandle(entRoom *ent.Room, err error) (model.Room, error) {
	return FromEntRoom(entRoom), err
}

func FromEntRooms(entRooms []*ent.Room) model.Rooms {
	rooms := make(model.Rooms, len(entRooms))
	for i, entRoom := range entRooms {
		rooms[i] = FromEntRoom(entRoom)
	}
	return rooms
}

func FromEntRoomsWithErrHandle(entRooms []*ent.Room, err error) (model.Rooms, error) {
	return FromEntRooms(entRooms), err
}
