package entity

import (
	"context"
	"time"

	"github.com/abc-valera/netsly-api-golang/internal/domain/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
)

var (
	ErrChatMemberChatIDInvalid = coderr.NewMessage(coderr.CodeInvalidArgument, "Provided invalid chat ID")
	ErrChatMemberUserIDInvalid = coderr.NewMessage(coderr.CodeInvalidArgument, "Provided invalid user ID")
)

type ChatMember struct {
	command command.IChatMember
}

func NewChatMember(
	command command.IChatMember,
) ChatMember {
	return ChatMember{
		command: command,
	}
}

type ChatMemberCreateRequest struct {
	ChatRoomID string
	UserID     string
}

func (c ChatMember) Create(ctx context.Context, req ChatMemberCreateRequest) error {
	// Validation
	if req.ChatRoomID == "" {
		return ErrChatMemberChatIDInvalid
	}
	if req.UserID == "" {
		return ErrChatMemberUserIDInvalid
	}

	createdAt := time.Now()

	return c.command.Create(ctx, model.ChatMember{
		ChatRoomID: req.ChatRoomID,
		UserID:     req.UserID,
		CreatedAt:  createdAt,
	})
}

func (c ChatMember) Delete(ctx context.Context, chatRoomID, userID string) error {
	// Validation
	if chatRoomID == "" {
		return ErrChatMemberChatIDInvalid
	}
	if userID == "" {
		return ErrChatMemberUserIDInvalid
	}

	return c.command.Delete(ctx, chatRoomID, userID)
}
