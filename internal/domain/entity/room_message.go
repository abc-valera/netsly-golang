package entity

import (
	"context"
	"time"

	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"
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
}

func NewRoomMessage(
	command command.IRoomMessage,
	query query.IRoomMessage,
) IRoomMessage {
	return roomMessage{
		command:      command,
		IRoomMessage: query,
	}
}

type RoomMessageCreateRequest struct {
	Text string `validate:"required,min=1,max=2048"`

	UserID string `validate:"required,uuid"`
	RoomID string `validate:"required,uuid"`
}

func (e roomMessage) Create(ctx context.Context, req RoomMessageCreateRequest) (model.RoomMessage, error) {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	if err := global.Validate().Struct(req); err != nil {
		return model.RoomMessage{}, err
	}

	return e.command.Create(ctx, command.RoomMessageCreateRequest{
		RoomMessage: model.RoomMessage{
			ID:        uuid.New().String(),
			Text:      req.Text,
			CreatedAt: time.Now(),
		},
		UserID: req.UserID,
		RoomID: req.RoomID,
	})
}

type RoomMessageUpdateRequest struct {
	Text *string `validate:"min=1,max=2048"`
}

func (e roomMessage) Update(ctx context.Context, id string, req RoomMessageUpdateRequest) (model.RoomMessage, error) {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	if err := global.Validate().Struct(req); err != nil {
		return model.RoomMessage{}, err
	}

	return e.command.Update(ctx, id, command.RoomMessageUpdateRequest{
		Text: req.Text,
	})
}

func (e roomMessage) Delete(ctx context.Context, id string) error {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	if err := global.Validate().Var(id, "uuid"); err != nil {
		return err
	}

	return e.command.Delete(ctx, id)
}
