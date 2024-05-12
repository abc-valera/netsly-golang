package model

import (
	"time"

	"github.com/abc-valera/netsly-api-golang/pkg/core/coderr"
)

var (
	ErrUserNotFound = coderr.NewCodeMessage(coderr.CodeNotFound, "User not found")
)

type User struct {
	ID             string
	Username       string
	Email          string
	HashedPassword string
	Fullname       string
	Status         string
	CreatedAt      time.Time
}

type Users []User
