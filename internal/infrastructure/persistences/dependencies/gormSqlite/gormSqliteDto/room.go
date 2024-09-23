package gormSqliteDto

import (
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
)

type Room struct {
	ID            string       `gorm:"primaryKey;not null"`
	Name          string       `gorm:"unique;not null"`
	Description   string       `gorm:"not null"`
	CreatedAt     time.Time    `gorm:"not null"`
	UpdatedAt     time.Time    `gorm:"not null"`
	DeletedAt     time.Time    `gorm:"not null"`
	CreatorUserID string       `gorm:"not null"`
	Members       RoomMembers  `gorm:"foreignKey:RoomID;constraint:OnDelete:CASCADE"`
	Messages      RoomMessages `gorm:"foreignKey:RoomID;constraint:OnDelete:CASCADE"`
}

func NewRoom(room model.Room) Room {
	return Room{
		ID:            room.ID,
		Name:          room.Name,
		Description:   room.Description,
		CreatedAt:     room.CreatedAt,
		UpdatedAt:     room.UpdatedAt,
		DeletedAt:     room.DeletedAt,
		CreatorUserID: room.CreatorUserID,
	}
}

func (dto Room) ToDomain() model.Room {
	return model.Room{
		ID:            dto.ID,
		Name:          dto.Name,
		Description:   dto.Description,
		CreatedAt:     dto.CreatedAt,
		UpdatedAt:     dto.UpdatedAt,
		DeletedAt:     dto.DeletedAt,
		CreatorUserID: dto.CreatorUserID,
	}
}

type Rooms []Room

func NewRooms(domainRooms model.Rooms) Rooms {
	var rooms Rooms
	for _, domainRoom := range domainRooms {
		rooms = append(rooms, NewRoom(domainRoom))
	}
	return rooms
}

func (dtos Rooms) ToDomain() model.Rooms {
	var domainRooms model.Rooms
	for _, dto := range dtos {
		domainRooms = append(domainRooms, dto.ToDomain())
	}
	return domainRooms
}
