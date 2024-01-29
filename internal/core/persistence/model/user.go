package model

import (
	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/core/persistence/model/common"
)

var (
	ErrUserNotFound = coderr.NewMessage(coderr.CodeNotFound, "User not found")
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
