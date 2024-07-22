package gormSqliteDto

import (
	"time"

	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
)

type RoomMember struct {
	CreatedAt time.Time `gorm:"not null"`
	DeletedAt time.Time `gorm:"not null"`

	UserID string `gorm:"primaryKey;not null"`
	RoomID string `gorm:"primaryKey;not null"`
}

func NewDomainRoomMember(roomMember RoomMember) model.RoomMember {
	return model.RoomMember{
		CreatedAt: roomMember.CreatedAt,
		DeletedAt: roomMember.DeletedAt,
	}
}

type RoomMembers []RoomMember

func NewDomainRoomMembers(roomMembers RoomMembers) model.RoomMembers {
	var domainRoomMembers model.RoomMembers
	for _, roomMember := range roomMembers {
		domainRoomMembers = append(domainRoomMembers, NewDomainRoomMember(roomMember))
	}
	return domainRoomMembers
}
