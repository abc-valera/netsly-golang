package entity

import (
	"context"
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/global"
	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"
)

type IRoom interface {
	Create(ctx context.Context, req RoomCreateRequest) (model.Room, error)
	Update(ctx context.Context, roomID string, req RoomUpdateRequest) (model.Room, error)
	Delete(ctx context.Context, roomID string) error

	query.IRoom
}

type room struct {
	IDependency

	query.IRoom
}

func newRoom(dep IDependency) IRoom {
	return room{
		IDependency: dep,

		IRoom: dep.Q().Room,
	}
}

type RoomCreateRequest struct {
	Name          string `validate:"required,min=4,max=64"`
	Description   string `validate:"max=256"`
	CreatorUserID string `validate:"required,uuid"`
}

func (e room) Create(ctx context.Context, req RoomCreateRequest) (model.Room, error) {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	if err := global.Validate().Struct(req); err != nil {
		return model.Room{}, err
	}

	room := model.Room{
		ID:            uuid.NewString(),
		Name:          req.Name,
		Description:   req.Description,
		CreatedAt:     time.Now().Truncate(time.Millisecond).Local(),
		CreatorUserID: req.CreatorUserID,
	}

	txFunc := func(
		ctx context.Context,
		txC command.Commands,
		txQ query.Queries,
		txE Entities,
	) error {
		if err := txC.Room.Create(ctx, room); err != nil {
			return err
		}

		if _, err := txE.RoomMember.Create(ctx, RoomMemberCreateRequest{
			UserID: req.CreatorUserID,
			RoomID: room.ID,
		}); err != nil {
			return err
		}

		return nil
	}

	if err := e.RunInTX(ctx, txFunc); err != nil {
		return model.Room{}, err
	}

	return room, nil
}

type RoomUpdateRequest struct {
	Name        *string `validate:"min=4,max=64"`
	Description *string `validate:"max=256"`
}

func (e room) Update(ctx context.Context, roomID string, req RoomUpdateRequest) (model.Room, error) {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	if err := global.Validate().Struct(req); err != nil {
		return model.Room{}, err
	}

	room, err := e.Q().Room.Get(ctx, model.Room{ID: roomID})
	if err != nil {
		return model.Room{}, err
	}

	room.UpdatedAt = time.Now().Truncate(time.Millisecond).Local()

	if req.Name != nil {
		room.Name = *req.Name
	}

	if req.Description != nil {
		room.Description = *req.Description
	}

	if err := e.C().Room.Update(ctx, room); err != nil {
		return model.Room{}, err
	}

	return room, nil
}

func (e room) Delete(ctx context.Context, roomID string) error {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	return e.C().Room.Delete(ctx, model.Room{ID: roomID})
}
