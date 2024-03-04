package model

import (
	"github.com/abc-valera/netsly-api-golang/internal/domain/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model/common"
)

var (
	ErrRoomNotFound = coderr.NewCodeMessage(coderr.CodeNotFound, "Room not found")
)

type Room struct {
	common.BaseEntity
	Name        string
	CreatorID   string
	Description string
}

type Rooms []Room
