package bunSqliteQuery

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/selector"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/bunSqlite/bunSqliteDto"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/bunSqlite/bunSqliteErrutil"
	"github.com/uptrace/bun"
)

type room struct {
	db bun.IDB
}

func NewRoom(db bun.IDB) query.IRoom {
	return &room{
		db: db,
	}
}

func (q room) GetByID(ctx context.Context, id string) (model.Room, error) {
	room := bunSqliteDto.Room{}
	err := q.db.NewSelect().Model(&room).Where("id = ?", id).Scan(ctx)
	return room.ToDomain(), bunSqliteErrutil.HandleQueryResult(err)
}

func (q room) GetByName(ctx context.Context, name string) (model.Room, error) {
	room := bunSqliteDto.Room{}
	err := q.db.NewSelect().Model(&room).Where("name = ?", name).Scan(ctx)
	return room.ToDomain(), bunSqliteErrutil.HandleQueryResult(err)
}

func (q room) GetAllByUserID(ctx context.Context, userID string, selector selector.Selector) (model.Rooms, error) {
	rooms := bunSqliteDto.Rooms{}
	err := q.db.NewSelect().Model(&rooms).Where("user_id = ?", userID).Scan(ctx)
	return rooms.ToDomain(), bunSqliteErrutil.HandleQueryResult(err)
}

func (q room) SearchAllByName(ctx context.Context, keyword string, selector selector.Selector) (model.Rooms, error) {
	rooms := bunSqliteDto.Rooms{}
	err := q.db.NewSelect().Model(&rooms).Where("name LIKE ?", "%"+keyword+"%").Scan(ctx)
	return rooms.ToDomain(), bunSqliteErrutil.HandleQueryResult(err)
}
