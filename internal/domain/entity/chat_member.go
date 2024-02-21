package entity

import (
	"context"
	"time"

	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
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
	ChatRoomID string `validate:"required,uuid"`
	UserID     string `validate:"required,uuid"`
}

func (c ChatMember) Create(ctx context.Context, req ChatMemberCreateRequest) (model.ChatMember, error) {
	if err := global.Validator().Struct(req); err != nil {
		return model.ChatMember{}, err
	}

	createdAt := time.Now()

	return c.command.Create(ctx, model.ChatMember{
		ChatRoomID: req.ChatRoomID,
		UserID:     req.UserID,
		CreatedAt:  createdAt,
	})
}

func (c ChatMember) Delete(ctx context.Context, chatRoomID, userID string) error {
	if err := global.Validator().Var(chatRoomID, "uuid"); err != nil {
		return err
	}

	return c.command.Delete(ctx, chatRoomID, userID)
}
