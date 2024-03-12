package model

import (
	"time"

	"github.com/abc-valera/netsly-api-golang/internal/domain/coderr"
)

var (
	ErrRoomNotFound = coderr.NewCodeMessage(coderr.CodeNotFound, "Room not found")
)

type Room struct {
	ID          string
	Name        string
	Description string
	CreatedAt   time.Time

	CreatorID string
}

type Rooms []Room
