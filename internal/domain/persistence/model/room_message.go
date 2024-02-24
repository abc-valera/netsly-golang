package model

import (
	"github.com/abc-valera/netsly-api-golang/internal/domain/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model/common"
)

var (
	ErrRoomMessageNotFound = coderr.NewMessage(coderr.CodeNotFound, "RoomMessage not found")
)

type RoomMessage struct {
	common.BaseEntity
	RoomID string
	UserID string
	Text   string
}

type RoomMessages []RoomMessage
