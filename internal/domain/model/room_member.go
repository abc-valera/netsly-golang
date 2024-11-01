package model

import (
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/util/coderr"
)

var ErrRoomMemberNotFound = coderr.NewCodeMessage(coderr.CodeNotFound, "RoomMember not found")

// RoomMember is a struct that represents a chat member entity.
// Technically, it's a many-to-many relationship between Room and User entities.
type RoomMember struct {
	UserID    string
	RoomID    string
	CreatedAt time.Time
	DeletedAt time.Time
}
