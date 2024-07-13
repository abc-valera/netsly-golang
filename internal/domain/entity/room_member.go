package entity

import (
	"context"
	"time"

	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query"
	"go.opentelemetry.io/otel/trace"
)

type IRoomMember interface {
	Create(ctx context.Context, req RoomMemberCreateRequest) (model.RoomMember, error)
	Delete(ctx context.Context, roomID, userID string) error

	query.IRoomMember
}

type roomMember struct {
	command command.IRoomMember
	query.IRoomMember
}

func NewRoomMember(
	command command.IRoomMember,
	query query.IRoomMember,
) IRoomMember {
	return roomMember{
		command:     command,
		IRoomMember: query,
	}
}

type RoomMemberCreateRequest struct {
	UserID string `validate:"required,uuid"`
	RoomID string `validate:"required,uuid"`
}

func (e roomMember) Create(ctx context.Context, req RoomMemberCreateRequest) (model.RoomMember, error) {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	if err := global.Validate().Struct(req); err != nil {
		return model.RoomMember{}, err
	}

	return e.command.Create(ctx, model.RoomMember{
		CreatedAt: time.Now(),
		UserID:    req.UserID,
		RoomID:    req.RoomID,
	})
}

func (e roomMember) Delete(ctx context.Context, roomID, userID string) error {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	if err := global.Validate().Var(roomID, "uuid"); err != nil {
		return err
	}

	return e.command.Delete(ctx, roomID, userID)
}
