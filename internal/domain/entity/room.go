package entity

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/internal/domain/entity/common"
	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
)

type Room struct {
	command command.IRoom
}

func NewRoom(
	command command.IRoom,
) Room {
	return Room{
		command: command,
	}
}

type RoomCreateRequest struct {
	Name        string `validate:"required,min=4,max=64"`
	Description string `validate:"max=256"`
	CreatorID   string `validate:"required,uuid"`
}

func (c Room) Create(ctx context.Context, req RoomCreateRequest) (model.Room, error) {
	if err := global.Validator().Struct(req); err != nil {
		return model.Room{}, err
	}

	baseModel := common.NewBaseEntity()

	return c.command.Create(ctx, model.Room{
		BaseEntity:  baseModel,
		Name:        req.Name,
		Description: req.Description,
		CreatorID:   req.CreatorID,
	})
}

type RoomUpdateRequest struct {
	Name        *string `validate:"min=4,max=64"`
	Description *string `validate:"max=256"`
}

func (c Room) Update(ctx context.Context, roomID string, req RoomUpdateRequest) (model.Room, error) {
	if err := global.Validator().Struct(req); err != nil {
		return model.Room{}, err
	}

	return c.command.Update(ctx, roomID, command.RoomUpdate{
		Description: req.Description,
	})
}

func (c Room) Delete(ctx context.Context, roomID string) error {
	if err := global.Validator().Var(roomID, "uuid"); err != nil {
		return err
	}

	return c.command.Delete(ctx, roomID)
}
