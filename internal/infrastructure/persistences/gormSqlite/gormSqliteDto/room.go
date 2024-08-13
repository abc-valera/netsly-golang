package gormSqliteDto

import (
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
)

type Room struct {
	ID          string    `gorm:"primaryKey;not null"`
	Name        string    `gorm:"unique;not null"`
	Description string    `gorm:"not null"`
	CreatedAt   time.Time `gorm:"not null"`
	UpdatedAt   time.Time `gorm:"not null"`
	DeletedAt   time.Time `gorm:"not null"`

	CreatorUserID string       `gorm:"not null"`
	Members       RoomMembers  `gorm:"foreignKey:RoomID;constraint:OnDelete:CASCADE"`
	Messages      RoomMessages `gorm:"foreignKey:RoomID;constraint:OnDelete:CASCADE"`
}

func NewDomainRoom(room Room) model.Room {
	return model.Room{
		ID:          room.ID,
		Name:        room.Name,
		Description: room.Description,
		CreatedAt:   room.CreatedAt,
		UpdatedAt:   room.UpdatedAt,
		DeletedAt:   room.DeletedAt,
	}
}

type Rooms []Room

func NewDomainRooms(rooms Rooms) model.Rooms {
	var domainRooms model.Rooms
	for _, room := range rooms {
		domainRooms = append(domainRooms, NewDomainRoom(room))
	}
	return domainRooms
}
