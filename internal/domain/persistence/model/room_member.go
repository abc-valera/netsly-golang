package model

import (
	"time"

	"github.com/abc-valera/netsly-api-golang/internal/domain/coderr"
)

var (
	ErrRoomMemberNotFound = coderr.NewMessage(coderr.CodeNotFound, "RoomMember not found")
)

// RoomMember is a struct that represents a chat member entity.
// Technically, it's a many-to-many relationship between Room and User entities.
type RoomMember struct {
	RoomID    string
	UserID    string
	CreatedAt time.Time
}

type RoomMembers []RoomMember
