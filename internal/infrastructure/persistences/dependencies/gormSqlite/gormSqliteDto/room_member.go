package gormSqliteDto

import (
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
)

type RoomMember struct {
	CreatedAt time.Time    `gorm:"not null"`
	DeletedAt time.Time    `gorm:"not null"`
	UserID    string       `gorm:"primaryKey;not null"`
	RoomID    string       `gorm:"primaryKey;not null"`
	Messages  RoomMessages `gorm:"foreignKey:UserID,RoomID;constraint:OnDelete:CASCADE"`
}

func NewRoomMember(roomMember model.RoomMember) RoomMember {
	return RoomMember{
		CreatedAt: roomMember.CreatedAt,
		DeletedAt: roomMember.DeletedAt,
		UserID:    roomMember.UserID,
		RoomID:    roomMember.RoomID,
	}
}

func (dto RoomMember) ToDomain() model.RoomMember {
	return model.RoomMember{
		CreatedAt: dto.CreatedAt,
		DeletedAt: dto.DeletedAt,
		UserID:    dto.UserID,
		RoomID:    dto.RoomID,
	}
}

type RoomMembers []RoomMember

func NewRoomMembers(domainRoomMembers model.RoomMembers) RoomMembers {
	var roomMembers RoomMembers
	for _, domainRoomMember := range domainRoomMembers {
		roomMembers = append(roomMembers, NewRoomMember(domainRoomMember))
	}
	return roomMembers
}

func (dtos RoomMembers) ToDomain() model.RoomMembers {
	var domainRoomMembers model.RoomMembers
	for _, dto := range dtos {
		domainRoomMembers = append(domainRoomMembers, dto.ToDomain())
	}
	return domainRoomMembers
}
