package entity

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/service"
)

type RoomMember struct {
	command command.IRoomMember

	timeMaker service.ITimeMaker
}

func NewRoomMember(
	command command.IRoomMember,
	timeMaker service.ITimeMaker,
) RoomMember {
	return RoomMember{
		command:   command,
		timeMaker: timeMaker,
	}
}

type RoomMemberCreateRequest struct {
	UserID string `validate:"required,uuid"`
	RoomID string `validate:"required,uuid"`
}

func (rm RoomMember) Create(ctx context.Context, req RoomMemberCreateRequest) (model.RoomMember, error) {
	if err := global.Validator().Struct(req); err != nil {
		return model.RoomMember{}, err
	}

	return rm.command.Create(ctx, model.RoomMember{
		CreatedAt: rm.timeMaker.Now(),
		UserID:    req.UserID,
		RoomID:    req.RoomID,
	})
}

func (rm RoomMember) Delete(ctx context.Context, roomID, userID string) error {
	if err := global.Validator().Var(roomID, "uuid"); err != nil {
		return err
	}

	return rm.command.Delete(ctx, roomID, userID)
}
