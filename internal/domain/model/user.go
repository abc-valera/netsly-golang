package model

import (
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/util/coderr"
	"github.com/abc-valera/netsly-golang/internal/domain/util/enum"
)

var ErrUserNotFound = coderr.NewCodeMessage(coderr.CodeNotFound, "User not found")

type User struct {
	ID             string
	Username       string
	Email          string
	HashedPassword string
	Fullname       string
	Status         string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
}

type UserMood int

const (
	UserStatusHappy UserMood = iota + 1
	UserStatusSad
	UserStatusAngry

	userStatusEnd
)

func (s UserMood) IsValid() bool {
	return s > 0 && s < userStatusEnd
}

var _ enum.IEnum = (*UserMood)(nil)
