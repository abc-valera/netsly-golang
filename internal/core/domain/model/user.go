package model

import (
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/coderr"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/model/common"
)

var (
	ErrUserNotFound = coderr.NewMessage(coderr.CodeNotFound, "User not found")

	ErrUserWithUsernameAlreadyExists = coderr.NewMessage(coderr.CodeAlreadyExists, "User with such username already exists")
	ErrUserWithEmailAlreadyExists    = coderr.NewMessage(coderr.CodeAlreadyExists, "User with such email already exists")
)

type User struct {
	common.BaseModel
	Username       string
	Email          string
	HashedPassword string
	Fullname       string
	Status         string
}

type Users []User
