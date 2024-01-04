package domain

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/codeerr"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/domain/common"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/model"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/command"
)

var (
	ErrChatMessageChatIDInvalid  = codeerr.NewMessageErr(codeerr.CodeInvalidArgument, "Provided invalid chat ID")
	ErrChatMessageUserIDInvalid  = codeerr.NewMessageErr(codeerr.CodeInvalidArgument, "Provided invalid user ID")
	ErrChatMessageMessageInvalid = codeerr.NewMessageErr(codeerr.CodeInvalidArgument, "Provided invalid message")
)

type ChatMessageDomain struct {
	command command.IChatMessageCommand
}

func NewChatMessageDomain(
	command command.IChatMessageCommand,
) ChatMessageDomain {
	return ChatMessageDomain{
		command: command,
	}
}

type ChatMessageCreateRequest struct {
	ChatID string
	UserID string
	Text   string
}

func (c ChatMessageDomain) Create(ctx context.Context, req ChatMessageCreateRequest) error {
	// Validation
	if req.ChatID == "" {
		return ErrChatMessageChatIDInvalid
	}
	if req.UserID == "" {
		return ErrChatMessageUserIDInvalid
	}
	if req.Text == "" || len(req.Text) > 2048 {
		return ErrChatMessageMessageInvalid
	}

	baseModel := common.NewBaseModel()

	// Create chat message
	return c.command.Create(ctx, model.ChatMessage{
		BaseModel:  baseModel,
		ChatRoomID: req.ChatID,
		UserID:     req.UserID,
		Text:       req.Text,
	})
}

type ChatMessageUpdateRequest struct {
	Text *string
}

func (c ChatMessageDomain) Update(ctx context.Context, id string, req ChatMessageUpdateRequest) error {
	// Validation
	if id == "" {
		return ErrChatMessageChatIDInvalid
	}
	if req.Text != nil && (*req.Text == "" || len(*req.Text) > 2048) {
		return ErrChatMessageMessageInvalid
	}

	return c.command.Update(ctx, id, command.ChatMessageUpdate{
		Text: req.Text,
	})
}

func (c ChatMessageDomain) Delete(ctx context.Context, id string) error {
	// Validation
	if id == "" {
		return ErrChatMessageChatIDInvalid
	}

	return c.command.Delete(ctx, id)
}
