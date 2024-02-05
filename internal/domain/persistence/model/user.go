package model

import (
	"github.com/abc-valera/netsly-api-golang/internal/domain/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model/common"
)

var (
	ErrUserNotFound = coderr.NewMessage(coderr.CodeNotFound, "User not found")
)

type User struct {
	common.BaseEntity
	Username       string
	Email          string
	HashedPassword string
	Fullname       string
	Status         string
}

type Users []User
