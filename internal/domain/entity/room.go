package entity

import (
	"context"
	"time"

	"github.com/abc-valera/netsly-api-golang/internal/core/optional"
	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query"
	"github.com/google/uuid"
)

type IRoom interface {
	Create(ctx context.Context, req RoomCreateRequest) (model.Room, error)
	Update(ctx context.Context, roomID string, req RoomUpdateRequest) (model.Room, error)
	Delete(ctx context.Context, roomID string) error

	query.IRoom
}

type room struct {
	command command.IRoom
	query.IRoom
}

func NewRoom(
	command command.IRoom,
	query query.IRoom,
) IRoom {
	return room{
		command: command,
		IRoom:   query,
	}
}

type RoomCreateRequest struct {
	Name          string `validate:"required,min=4,max=64"`
	Description   string `validate:"max=256"`
	CreatorUserID string `validate:"required,uuid"`
}

func (r room) Create(ctx context.Context, req RoomCreateRequest) (model.Room, error) {
	if err := global.Validate().Struct(req); err != nil {
		return model.Room{}, err
	}

	return r.command.Create(ctx, model.Room{
		ID:            uuid.New().String(),
		Name:          req.Name,
		Description:   req.Description,
		CreatedAt:     time.Now(),
		CreatorUserID: req.CreatorUserID,
	})
}

type RoomUpdateRequest struct {
	Name        optional.Optional[string] `validate:"min=4,max=64"`
	Description optional.Optional[string] `validate:"max=256"`
}

func (r room) Update(ctx context.Context, roomID string, req RoomUpdateRequest) (model.Room, error) {
	if err := global.Validate().Struct(req); err != nil {
		return model.Room{}, err
	}

	return r.command.Update(ctx, roomID, command.RoomUpdate{
		Description: req.Description,
	})
}

func (r room) Delete(ctx context.Context, roomID string) error {
	if err := global.Validate().Var(roomID, "uuid"); err != nil {
		return err
	}

	return r.command.Delete(ctx, roomID)
}
