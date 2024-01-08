package domain

import (
	"context"
	"time"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/codeerr"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/model"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/command"
)

var (
	ErrChatMemberChatIDInvalid = codeerr.NewMessage(codeerr.CodeInvalidArgument, "Provided invalid chat ID")
	ErrChatMemberUserIDInvalid = codeerr.NewMessage(codeerr.CodeInvalidArgument, "Provided invalid user ID")
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
