package entity

import (
	"context"
	"time"

	"github.com/abc-valera/netsly-api-golang/pkg/core/global"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/model"
)

type IRoomMember interface {
	Create(ctx context.Context, req RoomMemberCreateRequest) (model.RoomMember, error)
	Delete(ctx context.Context, roomID, userID string) error
}

type roomMember struct {
	command command.IRoomMember
}

func NewRoomMember(
	command command.IRoomMember,
) IRoomMember {
	return roomMember{
		command: command,
	}
}

type RoomMemberCreateRequest struct {
	UserID string `validate:"required,uuid"`
	RoomID string `validate:"required,uuid"`
}

func (rm roomMember) Create(ctx context.Context, req RoomMemberCreateRequest) (model.RoomMember, error) {
	if err := global.Validator().Struct(req); err != nil {
		return model.RoomMember{}, err
	}

	return rm.command.Create(ctx, model.RoomMember{
		CreatedAt: time.Now(),
		UserID:    req.UserID,
		RoomID:    req.RoomID,
	})
}

func (rm roomMember) Delete(ctx context.Context, roomID, userID string) error {
	if err := global.Validator().Var(roomID, "uuid"); err != nil {
		return err
	}

	return rm.command.Delete(ctx, roomID, userID)
}
