package model

import (
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/util/coderr"
)

var ErrRoomNotFound = coderr.NewCodeMessage(coderr.CodeNotFound, "Room not found")

type Room struct {
	ID          string
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time

	CreatorUserID string
}

type Rooms []Room
