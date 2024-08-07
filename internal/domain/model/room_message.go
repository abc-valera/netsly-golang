package model

import (
	"time"

	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
)

var ErrRoomMessageNotFound = coderr.NewCodeMessage(coderr.CodeNotFound, "RoomMessage not found")

type RoomMessage struct {
	ID        string
	Text      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type RoomMessages []RoomMessage
