package bunSqliteCommand

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/bunSqlite/bunSqliteDto"
	bunSqlitErrutil "github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/bunSqlite/bunSqliteErrutil"
	"github.com/uptrace/bun"
)

type roomMember struct {
	db bun.IDB
}

func NewRoomMember(db bun.IDB) command.IRoomMember {
	return &roomMember{
		db: db,
	}
}

func (c roomMember) Create(ctx context.Context, req command.RoomMemberCreateRequest) (model.RoomMember, error) {
	roomMember := bunSqliteDto.RoomMember{
		CreatedAt: req.RoomMember.CreatedAt,
		DeletedAt: req.RoomMember.DeletedAt,

		UserID: req.UserID,
		RoomID: req.RoomID,
	}

	res, err := c.db.NewInsert().Model(&roomMember).Exec(ctx)
	return roomMember.ToDomain(), bunSqlitErrutil.HandleCommandResult(res, err)
}

func (c roomMember) Delete(ctx context.Context, userID string, roomID string) error {
	roomMember := bunSqliteDto.RoomMember{
		UserID: userID,
		RoomID: roomID,
	}
	res, err := c.db.NewDelete().Model(&roomMember).WherePK().Exec(ctx)
	return bunSqlitErrutil.HandleCommandResult(res, err)
}
