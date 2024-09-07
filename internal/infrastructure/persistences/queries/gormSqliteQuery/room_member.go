package gormSqliteQuery

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/domain/global"
	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/selector"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/dependencies/gormSqlite/gormSqliteDto"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/dependencies/gormSqlite/gormSqliteErrutil"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/queries/gormSqliteQuery/gormSelector"
	"gorm.io/gorm"
)

type roomMember struct {
	db *gorm.DB
}

func NewRoomMember(db *gorm.DB) query.IRoomMember {
	return &roomMember{
		db: db,
	}
}

func (q roomMember) GetByIDs(ctx context.Context, userID string, roomID string) (model.RoomMember, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	var roomMember gormSqliteDto.RoomMember
	res := q.db.Where("user_id = ? AND room_id = ?", userID, roomID).First(&roomMember)
	return roomMember.ToDomain(), gormSqliteErrutil.HandleQueryResult(res)
}

func (q roomMember) GetAllByRoomID(ctx context.Context, roomID string, selector selector.Selector) (model.RoomMembers, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	var roomMembers gormSqliteDto.RoomMembers
	res := gormSelector.WithSelector(q.db, selector).WithContext(ctx).
		Where("room_id = ?", roomID).
		Find(&roomMembers)
	return roomMembers.ToDomain(), gormSqliteErrutil.HandleQueryResult(res)
}
