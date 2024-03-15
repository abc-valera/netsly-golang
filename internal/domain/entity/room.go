package entity

import (
	"context"
	"time"

	"github.com/abc-valera/netsly-api-golang/internal/core/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
	"github.com/google/uuid"
)

type IRoom interface {
	Create(ctx context.Context, req RoomCreateRequest) (model.Room, error)
	Update(ctx context.Context, roomID string, req RoomUpdateRequest) (model.Room, error)
	Delete(ctx context.Context, roomID string) error
}

type room struct {
	command command.IRoom
}

func NewRoom(
	command command.IRoom,
) IRoom {
	return room{
		command: command,
	}
}

type RoomCreateRequest struct {
	Name        string `validate:"required,min=4,max=64"`
	Description string `validate:"max=256"`
	CreatorID   string `validate:"required,uuid"`
}

func (r room) Create(ctx context.Context, req RoomCreateRequest) (model.Room, error) {
	if err := global.Validator().Struct(req); err != nil {
		return model.Room{}, err
	}

	return r.command.Create(ctx, model.Room{
		ID:          uuid.New().String(),
		Name:        req.Name,
		Description: req.Description,
		CreatedAt:   time.Now(),
		CreatorID:   req.CreatorID,
	})
}

type RoomUpdateRequest struct {
	Name        *string `validate:"min=4,max=64"`
	Description *string `validate:"max=256"`
}

func (r room) Update(ctx context.Context, roomID string, req RoomUpdateRequest) (model.Room, error) {
	if err := global.Validator().Struct(req); err != nil {
		return model.Room{}, err
	}

	return r.command.Update(ctx, roomID, command.RoomUpdate{
		Description: req.Description,
	})
}

func (r room) Delete(ctx context.Context, roomID string) error {
	if err := global.Validator().Var(roomID, "uuid"); err != nil {
		return err
	}

	return r.command.Delete(ctx, roomID)
}
