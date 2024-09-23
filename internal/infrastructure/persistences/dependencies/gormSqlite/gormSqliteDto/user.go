package gormSqliteDto

import (
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
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

	Jokes        Jokes       `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Comments     Comments    `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Likes        Likes       `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	CreatedRooms Rooms       `gorm:"foreignKey:CreatorUserID;constraint:OnDelete:CASCADE"`
	RoomMembers  RoomMembers `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

func NewUser(user model.User) User {
	return User{
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

func (dto User) ToDomain() model.User {
	return model.User{
		ID:             dto.ID,
		Username:       dto.Username,
		Email:          dto.Email,
		HashedPassword: dto.HashedPassword,
		Fullname:       dto.Fullname,
		Status:         dto.Status,
		CreatedAt:      dto.CreatedAt,
		UpdatedAt:      dto.UpdatedAt,
		DeletedAt:      dto.DeletedAt,
	}
}

type Users []User

func NewUsers(users model.Users) Users {
	var dtos Users
	for _, user := range users {
		dtos = append(dtos, NewUser(user))
	}
	return dtos
}

func (dtos Users) ToDomain() model.Users {
	var domainUsers model.Users
	for _, user := range dtos {
		domainUsers = append(domainUsers, user.ToDomain())
	}
	return domainUsers
}
