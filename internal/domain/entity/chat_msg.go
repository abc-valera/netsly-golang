package entity

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/internal/domain/entity/common"
	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
)

type ChatMessage struct {
	command command.IChatMessage
}

func NewChatMessage(
	command command.IChatMessage,
) ChatMessage {
	return ChatMessage{
		command: command,
	}
}

type ChatMessageCreateRequest struct {
	ChatID string `validate:"required,uuid"`
	UserID string `validate:"required,uuid"`
	Text   string `validate:"required,min=1,max=2048"`
}

func (c ChatMessage) Create(ctx context.Context, req ChatMessageCreateRequest) error {
	if err := global.Validator().Struct(req); err != nil {
		return err
	}

	baseModel := common.NewBaseEntity()

	return c.command.Create(ctx, model.ChatMessage{
		BaseEntity: baseModel,
		ChatRoomID: req.ChatID,
		UserID:     req.UserID,
		Text:       req.Text,
	})
}

type ChatMessageUpdateRequest struct {
	Text *string `validate:"min=1,max=2048"`
}

func (c ChatMessage) Update(ctx context.Context, id string, req ChatMessageUpdateRequest) error {
	if err := global.Validator().Struct(req); err != nil {
		return err
	}

	return c.command.Update(ctx, id, command.ChatMessageUpdate{
		Text: req.Text,
	})
}

func (c ChatMessage) Delete(ctx context.Context, id string) error {
	if err := global.Validator().Var(id, "uuid"); err != nil {
		return err
	}

	return c.command.Delete(ctx, id)
}
