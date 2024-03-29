package dto

import (
	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/model"
	"github.com/abc-valera/netsly-api-golang/pkg/infrastructure/persistence/boiler/errors"
)

func ToDomainRoomMember(roomMember *sqlboiler.RoomMember) model.RoomMember {
	if roomMember == nil {
		return model.RoomMember{}
	}

	return model.RoomMember{
		CreatedAt: roomMember.CreatedAt,
		UserID:    roomMember.UserID,
		RoomID:    roomMember.RoomID,
	}
}

func ToDomainRoomMemberWithErrHandle(roomMember *sqlboiler.RoomMember, err error) (model.RoomMember, error) {
	return ToDomainRoomMember(roomMember), errors.HandleErr(err)
}

func ToDomainRoomMembers(roomMembers sqlboiler.RoomMemberSlice) model.RoomMembers {
	var domainRoomMembers model.RoomMembers
	for _, roomMember := range roomMembers {
		domainRoomMembers = append(domainRoomMembers, ToDomainRoomMember(roomMember))
	}
	return domainRoomMembers
}

func ToDomainRoomMembersWithErrHandle(roomMembers sqlboiler.RoomMemberSlice, err error) (model.RoomMembers, error) {
	return ToDomainRoomMembers(roomMembers), errors.HandleErr(err)
}
