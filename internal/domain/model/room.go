package model

import (
	"time"

	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
)

var ErrRoomNotFound = coderr.NewCodeMessage(coderr.CodeNotFound, "Room not found")

type Room struct {
	ID            string
	Name          string
	Description   string
	CreatorUserID string
	CreatedAt     time.Time
}

type Rooms []Room
