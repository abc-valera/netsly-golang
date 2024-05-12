package entity

import (
	"context"
	"time"

	"github.com/abc-valera/netsly-api-golang/pkg/domain/model"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/query"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/service"
)

type IRoomMember interface {
	Create(ctx context.Context, req RoomMemberCreateRequest) (model.RoomMember, error)
	Delete(ctx context.Context, roomID, userID string) error

	query.IRoomMember
}

type roomMember struct {
	command command.IRoomMember
	query.IRoomMember

	validator service.IValidator
}

func NewRoomMember(
	command command.IRoomMember,
	query query.IRoomMember,
	validator service.IValidator,
) IRoomMember {
	return roomMember{
		command:     command,
		IRoomMember: query,
		validator:   validator,
	}
}

type RoomMemberCreateRequest struct {
	UserID string `validate:"required,uuid"`
	RoomID string `validate:"required,uuid"`
}

func (rm roomMember) Create(ctx context.Context, req RoomMemberCreateRequest) (model.RoomMember, error) {
	if err := rm.validator.Struct(req); err != nil {
		return model.RoomMember{}, err
	}

	return rm.command.Create(ctx, model.RoomMember{
		CreatedAt: time.Now(),
		UserID:    req.UserID,
		RoomID:    req.RoomID,
	})
}

func (rm roomMember) Delete(ctx context.Context, roomID, userID string) error {
	if err := rm.validator.Var(roomID, "uuid"); err != nil {
		return err
	}

	return rm.command.Delete(ctx, roomID, userID)
}
