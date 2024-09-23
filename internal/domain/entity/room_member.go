package entity

import (
	"context"
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/global"
	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query"
	"go.opentelemetry.io/otel/trace"
)

type IRoomMember interface {
	Create(ctx context.Context, req RoomMemberCreateRequest) (model.RoomMember, error)
	Delete(ctx context.Context, roomID, userID string) error

	query.IRoomMember
}

type roomMember struct {
	IDependency

	query.IRoomMember
}

func newRoomMember(dep IDependency) IRoomMember {
	return roomMember{
		IDependency: dep,

		IRoomMember: dep.Q().RoomMember,
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

	roomMember := model.RoomMember{
		CreatedAt: time.Now().Truncate(time.Millisecond).Local(),
		UserID:    req.UserID,
		RoomID:    req.RoomID,
	}

	if err := e.C().RoomMember.Create(ctx, roomMember); err != nil {
		return model.RoomMember{}, err
	}

	return roomMember, nil
}

func (e roomMember) Delete(ctx context.Context, roomID, userID string) error {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	return e.C().RoomMember.Delete(ctx, model.RoomMember{UserID: userID, RoomID: roomID})
}
