package model

import (
	"time"

	"github.com/abc-valera/netsly-api-golang/pkg/core/coderr"
)

var (
	ErrRoomMemberNotFound = coderr.NewCodeMessage(coderr.CodeNotFound, "RoomMember not found")
)

// RoomMember is a struct that represents a chat member entity.
// Technically, it's a many-to-many relationship between Room and User entities.
type RoomMember struct {
	CreatedAt time.Time

	UserID string
	RoomID string
}

type RoomMembers []RoomMember
