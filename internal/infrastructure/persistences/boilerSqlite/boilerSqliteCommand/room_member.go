package boilerSqliteCommand

import (
	"context"

	"github.com/abc-valera/netsly-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-golang/internal/core/global"
	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/boilerSqlite/boilerSqliteDto"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/boilerSqlite/boilerSqliteErrutil"
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

func (c roomMember) Create(ctx context.Context, req command.RoomMemberCreateRequest) (model.RoomMember, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	roomMember := sqlboiler.RoomMember{
		CreatedAt: req.RoomMember.CreatedAt,
		DeletedAt: req.RoomMember.DeletedAt,

		RoomID: req.RoomID,
		UserID: req.UserID,
	}
	err := roomMember.Insert(ctx, c.executor, boil.Infer())
	return boilerSqliteDto.NewDomainRoomMember(&roomMember), err
}

func (r roomMember) Delete(ctx context.Context, RoomID string, UserID string) error {
	_, span := global.NewSpan(ctx)
	defer span.End()

	roomMember, err := sqlboiler.FindRoomMember(ctx, r.executor, RoomID, UserID)
	if err != nil {
		return err
	}
	_, err = roomMember.Delete(ctx, r.executor)
	return boilerSqliteErrutil.HandleErr(err)
}
