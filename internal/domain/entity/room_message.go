package entity

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/internal/domain/entity/common"
	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
)

type RoomMessage struct {
	command command.IRoomMessage
}

func NewRoomMessage(
	command command.IRoomMessage,
) RoomMessage {
	return RoomMessage{
		command: command,
	}
}

type RoomMessageCreateRequest struct {
	Text   string `validate:"required,min=1,max=2048"`
	UserID string `validate:"required,uuid"`
	RoomID string `validate:"required,uuid"`
}

func (c RoomMessage) Create(ctx context.Context, req RoomMessageCreateRequest) (model.RoomMessage, error) {
	if err := global.Validator().Struct(req); err != nil {
		return model.RoomMessage{}, err
	}

	baseModel := common.NewBaseEntity()

	return c.command.Create(ctx, model.RoomMessage{
		BaseEntity: baseModel,
		Text:       req.Text,
		UserID:     req.UserID,
		RoomID:     req.RoomID,
	})
}

type RoomMessageUpdateRequest struct {
	Text *string `validate:"min=1,max=2048"`
}

func (c RoomMessage) Update(ctx context.Context, id string, req RoomMessageUpdateRequest) (model.RoomMessage, error) {
	if err := global.Validator().Struct(req); err != nil {
		return model.RoomMessage{}, err
	}

	return c.command.Update(ctx, id, command.RoomMessageUpdate{
		Text: req.Text,
	})
}

func (c RoomMessage) Delete(ctx context.Context, id string) error {
	if err := global.Validator().Var(id, "uuid"); err != nil {
		return err
	}

	return c.command.Delete(ctx, id)
}
