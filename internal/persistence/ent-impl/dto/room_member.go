package dto

import (
	"github.com/abc-valera/netsly-api-golang/gen/ent"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
	errhandler "github.com/abc-valera/netsly-api-golang/internal/persistence/ent-impl/errors"
)

func FromEntRoomMember(entRoomMember *ent.RoomMember) model.RoomMember {
	if entRoomMember == nil {
		return model.RoomMember{}
	}
	return model.RoomMember{
		UserID:    entRoomMember.Edges.User.ID,
		RoomID:    entRoomMember.Edges.Room.ID,
		CreatedAt: entRoomMember.CreatedAt,
	}
}

func FromEntRoomMemberWithErrHandle(entRoomMember *ent.RoomMember, err error) (model.RoomMember, error) {
	return FromEntRoomMember(entRoomMember), errhandler.HandleErr(err)
}

func FromEntRoomMembers(entRoomMembers []*ent.RoomMember) model.RoomMembers {
	roomMembers := make(model.RoomMembers, len(entRoomMembers))
	for i, entRoomMember := range entRoomMembers {
		roomMembers[i] = FromEntRoomMember(entRoomMember)
	}
	return roomMembers
}

func FromEntRoomMembersWithErrHandle(entRoomMembers []*ent.RoomMember, err error) (model.RoomMembers, error) {
	return FromEntRoomMembers(entRoomMembers), errhandler.HandleErr(err)
}
