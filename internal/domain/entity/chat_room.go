package entity

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/internal/domain/entity/common"
	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
)

type ChatRoom struct {
	command command.IChatRoom
}

func NewChatRoom(
	command command.IChatRoom,
) ChatRoom {
	return ChatRoom{
		command: command,
	}
}

type ChatRoomCreateRequest struct {
	Name        string `validate:"required,min=4,max=64"`
	Description string `validate:"max=256"`
}

func (c ChatRoom) Create(ctx context.Context, req ChatRoomCreateRequest) (model.ChatRoom, error) {
	if err := global.Validator().Struct(req); err != nil {
		return model.ChatRoom{}, err
	}

	baseModel := common.NewBaseEntity()

	return c.command.Create(ctx, model.ChatRoom{
		BaseEntity:  baseModel,
		Name:        req.Name,
		Description: req.Description,
	})
}

type ChatRoomUpdateRequest struct {
	Name        *string `validate:"min=4,max=64"`
	Description *string `validate:"max=256"`
}

func (c ChatRoom) Update(ctx context.Context, chatRoomID string, req ChatRoomUpdateRequest) (model.ChatRoom, error) {
	if err := global.Validator().Struct(req); err != nil {
		return model.ChatRoom{}, err
	}

	return c.command.Update(ctx, chatRoomID, command.ChatRoomUpdate{
		Description: req.Description,
	})
}

func (c ChatRoom) Delete(ctx context.Context, chatRoomID string) error {
	if err := global.Validator().Var(chatRoomID, "uuid"); err != nil {
		return err
	}

	return c.command.Delete(ctx, chatRoomID)
}
