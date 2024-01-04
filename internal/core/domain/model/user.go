package model

import (
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/codeerr"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/model/common"
)

var (
	ErrUserNotFound = codeerr.NewMessage(codeerr.CodeNotFound, "User not found")

	ErrUserWithUsernameAlreadyExists = codeerr.NewMessage(codeerr.CodeAlreadyExists, "User with such username already exists")
	ErrUserWithEmailAlreadyExists    = codeerr.NewMessage(codeerr.CodeAlreadyExists, "User with such email already exists")
)

type User struct {
	common.BaseModel
	Username       string
	Email          string
	HashedPassword string
	Fullname       string
	Status         string
}

type Users []*User
