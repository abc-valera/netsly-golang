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

type ChatMemberDomain struct {
	command command.IChatMemberCommand
}

func NewChatMemberDomain(
	command command.IChatMemberCommand,
) ChatMemberDomain {
	return ChatMemberDomain{
		command: command,
	}
}

type ChatMemberCreateRequest struct {
	ChatRoomID string
	UserID     string
}

func (c ChatMemberDomain) Create(ctx context.Context, req ChatMemberCreateRequest) error {
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

func (c ChatMemberDomain) Delete(ctx context.Context, chatRoomID, userID string) error {
	// Validation
	if chatRoomID == "" {
		return ErrChatMemberChatIDInvalid
	}
	if userID == "" {
		return ErrChatMemberUserIDInvalid
	}

	return c.command.Delete(ctx, chatRoomID, userID)
}
