package entity

import (
	"context"
	"time"

	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/model"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/query"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/service"
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

	validator service.IValidator
}

func NewRoom(
	command command.IRoom,
	query query.IRoom,
	validator service.IValidator,
) IRoom {
	return room{
		command:   command,
		IRoom:     query,
		validator: validator,
	}
}

type RoomCreateRequest struct {
	Name          string `validate:"required,min=4,max=64"`
	Description   string `validate:"max=256"`
	CreatorUserID string `validate:"required,uuid"`
}

func (r room) Create(ctx context.Context, req RoomCreateRequest) (model.Room, error) {
	if err := r.validator.Struct(req); err != nil {
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
	Name        *string `validate:"min=4,max=64"`
	Description *string `validate:"max=256"`
}

func (r room) Update(ctx context.Context, roomID string, req RoomUpdateRequest) (model.Room, error) {
	if err := r.validator.Struct(req); err != nil {
		return model.Room{}, err
	}

	return r.command.Update(ctx, roomID, command.RoomUpdate{
		Description: req.Description,
	})
}

func (r room) Delete(ctx context.Context, roomID string) error {
	if err := r.validator.Var(roomID, "uuid"); err != nil {
		return err
	}

	return r.command.Delete(ctx, roomID)
}
