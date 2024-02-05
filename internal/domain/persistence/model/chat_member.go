package model

import (
	"time"

	"github.com/abc-valera/netsly-api-golang/internal/domain/coderr"
)

var (
	ErrChatMemberNotFound = coderr.NewMessage(coderr.CodeNotFound, "ChatMember not found")
)

// ChatMember is a struct that represents a chat member entity.
// Technically, it's a many-to-many relationship between ChatRoom and User entities.
type ChatMember struct {
	ChatRoomID string
	UserID     string
	CreatedAt  time.Time
}

type ChatMembers []ChatMember
