package bunSqliteQuery

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/selector"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/dependencies/bunSqlite/bunSqliteDto"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/dependencies/bunSqlite/bunSqliteErrors"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/queries/bunSqliteQuery/bunSqliteSelector"
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
	return roomMessage.ToDomain(), bunSqliteErrors.HandleQueryResult(err)
}

func (q roomMessage) GetAllByRoomID(ctx context.Context, roomID string, s selector.Selector) (model.RoomMessages, error) {
	roomMessages := bunSqliteDto.RoomMessages{}
	err := bunSqliteSelector.NewSelectQuery(q.db, s).Model(&roomMessages).Where("room_id = ?", roomID).Scan(ctx)
	return roomMessages.ToDomain(), bunSqliteErrors.HandleQueryResult(err)
}

func (q roomMessage) SearchAllByText(ctx context.Context, keyword string, s selector.Selector) (model.RoomMessages, error) {
	roomMessages := bunSqliteDto.RoomMessages{}
	err := bunSqliteSelector.NewSelectQuery(q.db, s).Model(&roomMessages).Where("text LIKE ?", "%"+keyword+"%").Scan(ctx)
	return roomMessages.ToDomain(), bunSqliteErrors.HandleQueryResult(err)
}
