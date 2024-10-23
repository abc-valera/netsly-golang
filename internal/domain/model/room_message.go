package model

import (
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/util/coderr"
)

var ErrRoomMessageNotFound = coderr.NewCodeMessage(coderr.CodeNotFound, "RoomMessage not found")

type RoomMessage struct {
	ID        string
	Text      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time

	UserID string
	RoomID string
}
