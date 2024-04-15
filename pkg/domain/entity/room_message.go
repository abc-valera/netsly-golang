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

type IRoomMessage interface {
	Create(ctx context.Context, req RoomMessageCreateRequest) (model.RoomMessage, error)
	Update(ctx context.Context, id string, req RoomMessageUpdateRequest) (model.RoomMessage, error)
	Delete(ctx context.Context, id string) error

	query.IRoomMessage
}

type roomMessage struct {
	command command.IRoomMessage
	query.IRoomMessage

	validator service.IValidator
}

func NewRoomMessage(
	command command.IRoomMessage,
	query query.IRoomMessage,
	validator service.IValidator,
) IRoomMessage {
	return roomMessage{
		command:      command,
		IRoomMessage: query,
		validator:    validator,
	}
}

type RoomMessageCreateRequest struct {
	Text   string `validate:"required,min=1,max=2048"`
	UserID string `validate:"required,uuid"`
	RoomID string `validate:"required,uuid"`
}

func (rm roomMessage) Create(ctx context.Context, req RoomMessageCreateRequest) (model.RoomMessage, error) {
	if err := rm.validator.Struct(req); err != nil {
		return model.RoomMessage{}, err
	}

	return rm.command.Create(ctx, model.RoomMessage{
		ID:        uuid.New().String(),
		Text:      req.Text,
		CreatedAt: time.Now(),
		UserID:    req.UserID,
		RoomID:    req.RoomID,
	})
}

type RoomMessageUpdateRequest struct {
	Text *string `validate:"min=1,max=2048"`
}

func (rm roomMessage) Update(ctx context.Context, id string, req RoomMessageUpdateRequest) (model.RoomMessage, error) {
	if err := rm.validator.Struct(req); err != nil {
		return model.RoomMessage{}, err
	}

	return rm.command.Update(ctx, id, command.RoomMessageUpdate{
		Text: req.Text,
	})
}

func (rm roomMessage) Delete(ctx context.Context, id string) error {
	if err := rm.validator.Var(id, "uuid"); err != nil {
		return err
	}

	return rm.command.Delete(ctx, id)
}