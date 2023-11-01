package entity

import (
	"time"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/codeerr"
)

// ChatMember is a struct that represents a chat member entity.
// Technically, Technically, it's a many-to-many relationship between ChatRoom and User entities.
type ChatMember struct {
	ChatRoomID string
	UserID     string
	CreatedAt  time.Time
}

func NewChatMember(chatRoomID, userID string) (*ChatMember, error) {
	if chatRoomID == "" {
		return nil, codeerr.NewMsgErr(codeerr.CodeInvalidArgument, "Provided invalid chat room ID")
	}
	if userID == "" {
		return nil, codeerr.NewMsgErr(codeerr.CodeInvalidArgument, "Provided invalid user ID")
	}

	return &ChatMember{
		ChatRoomID: chatRoomID,
		UserID:     userID,
		CreatedAt:  time.Now(),
	}, nil
}

type ChatMembers []*ChatMember
