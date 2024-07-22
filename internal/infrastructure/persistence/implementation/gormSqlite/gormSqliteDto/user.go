package gormSqliteDto

import (
	"time"

	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
)

type User struct {
	ID             string    `gorm:"primaryKey;not null"`
	Username       string    `gorm:"unique;not null"`
	Email          string    `gorm:"unique;not null"`
	HashedPassword string    `gorm:"not null"`
	Fullname       string    `gorm:"not null"`
	Status         string    `gorm:"not null"`
	CreatedAt      time.Time `gorm:"not null"`
	UpdatedAt      time.Time `gorm:"not null"`
	DeletedAt      time.Time `gorm:"not null"`

	Jokes        Jokes        `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Comments     Comments     `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Likes        Likes        `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	CreatedRooms Rooms        `gorm:"foreignKey:CreatorUserID;constraint:OnDelete:CASCADE"`
	RoomMembers  RoomMembers  `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	RoomMessages RoomMessages `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

func NewDomainUser(user User) model.User {
	return model.User{
		ID:             user.ID,
		Username:       user.Username,
		Email:          user.Email,
		HashedPassword: user.HashedPassword,
		Fullname:       user.Fullname,
		Status:         user.Status,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
		DeletedAt:      user.DeletedAt,
	}
}

type Users []User

func NewDomainUsers(users Users) model.Users {
	var domainUsers model.Users
	for _, user := range users {
		domainUsers = append(domainUsers, NewDomainUser(user))
	}
	return domainUsers
}
