package boilerDto

import (
	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/boiler/errutil"
)

func NewDomainRoomMember(roomMember *sqlboiler.RoomMember) model.RoomMember {
	if roomMember == nil {
		return model.RoomMember{}
	}

	return model.RoomMember{
		UserID:    roomMember.UserID,
		RoomID:    roomMember.RoomID,
		CreatedAt: roomMember.CreatedAt,
	}
}

func NewDomainRoomMemberWithErrHandle(roomMember *sqlboiler.RoomMember, err error) (model.RoomMember, error) {
	return NewDomainRoomMember(roomMember), errutil.HandleErr(err)
}

func NewDomainRoomMembers(roomMembers sqlboiler.RoomMemberSlice) model.RoomMembers {
	var domainRoomMembers model.RoomMembers
	for _, roomMember := range roomMembers {
		domainRoomMembers = append(domainRoomMembers, NewDomainRoomMember(roomMember))
	}
	return domainRoomMembers
}

func NewDomainRoomMembersWithErrHandle(roomMembers sqlboiler.RoomMemberSlice, err error) (model.RoomMembers, error) {
	return NewDomainRoomMembers(roomMembers), errutil.HandleErr(err)
}
