package model

import (
	"time"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/codeerr"
)

var (
	ErrChatMemberNotFound = codeerr.NewMessage(codeerr.CodeNotFound, "ChatMember not found")

	ErrChatMemberAlreadyExists = codeerr.NewMessage(codeerr.CodeAlreadyExists, "ChatMember already exists")
)

// ChatMember is a struct that represents a chat member entity.
// Technically, it's a many-to-many relationship between ChatRoom and User entities.
type ChatMember struct {
	ChatRoomID string
	UserID     string
	CreatedAt  time.Time
}

type ChatMembers []ChatMember
