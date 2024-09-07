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

type roomMessage struct {
	db *gorm.DB
}

func NewRoomMessage(db *gorm.DB) query.IRoomMessage {
	return &roomMessage{
		db: db,
	}
}

func (q roomMessage) GetByID(ctx context.Context, id string) (model.RoomMessage, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	var roomMessage gormSqliteDto.RoomMessage
	res := q.db.Where("id = ?", id).First(&roomMessage)
	return roomMessage.ToDomain(), gormSqliteErrutil.HandleQueryResult(res)
}

func (q roomMessage) SearchAllByText(ctx context.Context, keyword string, selector selector.Selector) (model.RoomMessages, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	var roomMessages gormSqliteDto.RoomMessages
	res := gormSelector.WithSelector(q.db, selector).WithContext(ctx).
		Where("text LIKE ?", "%"+keyword+"%").
		Find(&roomMessages)
	return roomMessages.ToDomain(), gormSqliteErrutil.HandleQueryResult(res)
}

func (q roomMessage) GetAllByRoomID(ctx context.Context, roomID string, selector selector.Selector) (model.RoomMessages, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	var roomMessages gormSqliteDto.RoomMessages
	res := gormSelector.WithSelector(q.db, selector).WithContext(ctx).
		Where("room_id = ?", roomID).
		Find(&roomMessages)
	return roomMessages.ToDomain(), gormSqliteErrutil.HandleQueryResult(res)
}
