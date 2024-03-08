package entity

import (
	"context"
	"time"

	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
)

type RoomMember struct {
	command command.IRoomMember
}

func NewRoomMember(
	command command.IRoomMember,
) RoomMember {
	return RoomMember{
		command: command,
	}
}

type RoomMemberCreateRequest struct {
	UserID string `validate:"required,uuid"`
	RoomID string `validate:"required,uuid"`
}

func (c RoomMember) Create(ctx context.Context, req RoomMemberCreateRequest) (model.RoomMember, error) {
	if err := global.Validator().Struct(req); err != nil {
		return model.RoomMember{}, err
	}

	createdAt := time.Now()

	return c.command.Create(ctx, model.RoomMember{
		CreatedAt: createdAt,
		UserID:    req.UserID,
		RoomID:    req.RoomID,
	})
}

func (c RoomMember) Delete(ctx context.Context, roomID, userID string) error {
	if err := global.Validator().Var(roomID, "uuid"); err != nil {
		return err
	}

	return c.command.Delete(ctx, roomID, userID)
}
