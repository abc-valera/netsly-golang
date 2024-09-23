package entity

import (
	"context"
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/global"
	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query"
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
	IDependency

	query.IRoomMessage
}

func newRoomMessage(dep IDependency) IRoomMessage {
	return roomMessage{
		IDependency: dep,

		IRoomMessage: dep.Q().RoomMessage,
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

	roomMessage := model.RoomMessage{
		ID:        uuid.New().String(),
		Text:      req.Text,
		CreatedAt: time.Now().Truncate(time.Millisecond).Local(),
		UserID:    req.UserID,
		RoomID:    req.RoomID,
	}

	if err := e.C().RoomMessage.Create(ctx, roomMessage); err != nil {
		return model.RoomMessage{}, err
	}

	return roomMessage, nil
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

	roomMessage, err := e.Q().RoomMessage.GetByID(ctx, id)
	if err != nil {
		return model.RoomMessage{}, err
	}

	roomMessage.UpdatedAt = time.Now().Truncate(time.Millisecond).Local()

	if req.Text != nil {
		roomMessage.Text = *req.Text
	}

	if err := e.C().RoomMessage.Update(ctx, roomMessage); err != nil {
		return model.RoomMessage{}, err
	}

	return roomMessage, nil
}

func (e roomMessage) Delete(ctx context.Context, id string) error {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	return e.C().RoomMessage.Delete(ctx, model.RoomMessage{ID: id})
}
