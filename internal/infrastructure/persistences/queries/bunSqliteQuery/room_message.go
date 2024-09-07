package bunSqliteQuery

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/selector"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/dependencies/bunSqlite/bunSqliteDto"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/dependencies/bunSqlite/bunSqliteErrutil"
	"github.com/uptrace/bun"
)

type roomMessage struct {
	db bun.IDB
}

func NewRoomMessage(db bun.IDB) query.IRoomMessage {
	return &roomMessage{
		db: db,
	}
}

func (q roomMessage) GetByID(ctx context.Context, id string) (model.RoomMessage, error) {
	roomMessage := bunSqliteDto.RoomMessage{}
	err := q.db.NewSelect().Model(&roomMessage).Where("id = ?", id).Scan(ctx)
	return roomMessage.ToDomain(), bunSqliteErrutil.HandleQueryResult(err)
}

func (q roomMessage) GetAllByRoomID(ctx context.Context, roomID string, selector selector.Selector) (model.RoomMessages, error) {
	roomMessages := bunSqliteDto.RoomMessages{}
	err := q.db.NewSelect().Model(&roomMessages).Where("room_id = ?", roomID).Scan(ctx)
	return roomMessages.ToDomain(), bunSqliteErrutil.HandleQueryResult(err)
}

func (q roomMessage) SearchAllByText(ctx context.Context, keyword string, selector selector.Selector) (model.RoomMessages, error) {
	roomMessages := bunSqliteDto.RoomMessages{}
	err := q.db.NewSelect().Model(&roomMessages).Where("text LIKE ?", "%"+keyword+"%").Scan(ctx)
	return roomMessages.ToDomain(), bunSqliteErrutil.HandleQueryResult(err)
}
