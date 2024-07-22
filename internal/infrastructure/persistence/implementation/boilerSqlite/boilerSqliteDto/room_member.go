package boilerSqliteDto

import (
	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
)

func NewDomainRoomMember(roomMember *sqlboiler.RoomMember) model.RoomMember {
	if roomMember == nil {
		return model.RoomMember{}
	}

	return model.RoomMember{
		CreatedAt: roomMember.CreatedAt,
		DeletedAt: roomMember.DeletedAt,
	}
}

func NewDomainRoomMembers(roomMembers sqlboiler.RoomMemberSlice) model.RoomMembers {
	var domainRoomMembers model.RoomMembers
	for _, roomMember := range roomMembers {
		domainRoomMembers = append(domainRoomMembers, NewDomainRoomMember(roomMember))
	}
	return domainRoomMembers
}
