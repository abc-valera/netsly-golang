package bunSqliteQuery

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/selector"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/dependencies/bunSqlite/bunSqliteDto"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/dependencies/bunSqlite/bunSqliteErrors"
	"github.com/uptrace/bun"
)

type roomMember struct {
	db bun.IDB
}

func NewRoomMember(db bun.IDB) query.IRoomMember {
	return &roomMember{
		db: db,
	}
}

func (q roomMember) GetByIDs(ctx context.Context, userID string, roomID string) (model.RoomMember, error) {
	roomMember := bunSqliteDto.RoomMember{}
	err := q.db.NewSelect().Model(&roomMember).Where("user_id = ? AND room_id = ?", userID, roomID).Scan(ctx)
	return roomMember.ToDomain(), bunSqliteErrors.HandleQueryResult(err)
}

func (q roomMember) GetAllByRoomID(ctx context.Context, roomID string, s selector.Selector) (model.RoomMembers, error) {
	roomMembers := bunSqliteDto.RoomMembers{}
	err := q.db.NewSelect().Model(&roomMembers).Where("room_id = ?", roomID).Scan(ctx)
	return roomMembers.ToDomain(), bunSqliteErrors.HandleQueryResult(err)
}
