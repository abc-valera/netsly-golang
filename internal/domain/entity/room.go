package entity

import (
	"context"
	"time"

	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/commandTransactor"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query"
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
	roomCommand command.IRoom
	query.IRoom
	commandTransactor commandTransactor.ITransactor
}

func NewRoom(
	roomCommand command.IRoom,
	query query.IRoom,
	commandTransactor commandTransactor.ITransactor,
) IRoom {
	return room{
		roomCommand:       roomCommand,
		IRoom:             query,
		commandTransactor: commandTransactor,
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

	var returnRoom model.Room
	txFunc := func(ctx context.Context, txCommands persistence.Commands) error {
		room, err := txCommands.Room.Create(ctx, model.Room{
			ID:            uuid.New().String(),
			Name:          req.Name,
			Description:   req.Description,
			CreatorUserID: req.CreatorUserID,
			CreatedAt:     time.Now(),
		})
		if err != nil {
			return err
		}
		returnRoom = room

		if _, err := txCommands.RoomMember.Create(ctx, model.RoomMember{
			RoomID:    room.ID,
			UserID:    req.CreatorUserID,
			CreatedAt: time.Now(),
		}); err != nil {
			return err
		}

		return nil
	}

	if err := e.commandTransactor.PerformTX(ctx, txFunc); err != nil {
		return model.Room{}, err
	}

	return returnRoom, nil
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

	return e.roomCommand.Update(ctx, roomID, command.RoomUpdate{
		Description: req.Description,
	})
}

func (e room) Delete(ctx context.Context, roomID string) error {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	if err := global.Validate().Var(roomID, "uuid"); err != nil {
		return err
	}

	return e.roomCommand.Delete(ctx, roomID)
}
