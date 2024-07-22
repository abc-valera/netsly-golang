package gormSqliteQuery

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query/selector"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/gormSqlite/gormSqliteDto"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/gormSqlite/gormSqliteErrutil"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/gormSqlite/gormSqliteQuery/gormSelector"
	"gorm.io/gorm"
)

type room struct {
	db *gorm.DB
}

func NewRoom(db *gorm.DB) query.IRoom {
	return &room{
		db: db,
	}
}

func (q room) GetByID(ctx context.Context, id string) (model.Room, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	var room gormSqliteDto.Room
	res := q.db.Where("id = ?", id).First(&room)
	return gormSqliteDto.NewDomainRoom(room), gormSqliteErrutil.HandleQueryResult(res)
}

func (q room) GetByName(ctx context.Context, name string) (model.Room, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	var room gormSqliteDto.Room
	res := q.db.WithContext(ctx).Where("name = ?", name).First(&room)
	return gormSqliteDto.NewDomainRoom(room), gormSqliteErrutil.HandleQueryResult(res)
}

func (q room) SearchAllByName(ctx context.Context, keyword string, selector selector.Selector) (model.Rooms, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	var rooms gormSqliteDto.Rooms
	res := gormSelector.WithSelector(q.db, selector).WithContext(ctx).
		Where("name LIKE ?", "%"+keyword+"%").
		Find(&rooms)
	return gormSqliteDto.NewDomainRooms(rooms), gormSqliteErrutil.HandleQueryResult(res)
}

func (q room) GetAllByUserID(ctx context.Context, userID string, selector selector.Selector) (model.Rooms, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	var rooms gormSqliteDto.Rooms
	res := gormSelector.WithSelector(q.db, selector).WithContext(ctx).
		Where("user_id = ?", userID).
		Find(&rooms)
	return gormSqliteDto.NewDomainRooms(rooms), gormSqliteErrutil.HandleQueryResult(res)
}
