package boilerCommand

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/model"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/pkg/infrastructure/persistence/boiler/boilerDto"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type roomMember struct {
	executor boil.ContextExecutor
}

func NewRoomMember(executor boil.ContextExecutor) command.IRoomMember {
	return &roomMember{
		executor: executor,
	}
}

func (r roomMember) Create(ctx context.Context, req model.RoomMember) (model.RoomMember, error) {
	roomMember := sqlboiler.RoomMember{
		CreatedAt: req.CreatedAt,
		RoomID:    req.RoomID,
		UserID:    req.UserID,
	}
	err := roomMember.Insert(ctx, r.executor, boil.Infer())
	return boilerDto.ToDomainRoomMemberWithErrHandle(&roomMember, err)
}

func (r roomMember) Delete(ctx context.Context, RoomID string, UserID string) error {
	roomMember, err := sqlboiler.FindRoomMember(ctx, r.executor, RoomID, UserID)
	if err != nil {
		return err
	}
	_, err = roomMember.Delete(ctx, r.executor)
	return err
}
