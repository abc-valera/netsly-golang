package boilerSqliteQuery

import (
	"context"

	"github.com/abc-valera/netsly-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-golang/internal/core/global"
	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query"
	selector1 "github.com/abc-valera/netsly-golang/internal/domain/persistence/query/selector"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/boilerSqlite/boilerSqliteDto"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/boilerSqlite/boilerSqliteErrutil"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/boilerSqlite/boilerSqliteQuery/selector"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type room struct {
	executor boil.ContextExecutor
}

func NewRoom(executor boil.ContextExecutor) query.IRoom {
	return &room{
		executor: executor,
	}
}

func (r room) GetByID(ctx context.Context, id string) (model.Room, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	room, err := sqlboiler.FindRoom(ctx, r.executor, id)
	return boilerSqliteDto.NewDomainRoom(room), boilerSqliteErrutil.HandleErr(err)
}

func (r room) GetByName(ctx context.Context, name string) (model.Room, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	romm, err := sqlboiler.Rooms(sqlboiler.RoomWhere.Name.EQ(name)).One(ctx, r.executor)
	return boilerSqliteDto.NewDomainRoom(romm), boilerSqliteErrutil.HandleErr(err)
}

func (r room) GetAllByUserID(ctx context.Context, userID string, params selector1.Selector) (model.Rooms, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	mods := selector.ToBoilerSelectorPipe(
		params,
		qm.InnerJoin(sqlboiler.TableNames.RoomMembers+" rm ON rm.room_id = room.id"),
		sqlboiler.RoomMemberWhere.UserID.EQ(userID),
	)
	rooms, err := sqlboiler.Rooms(mods...).All(ctx, r.executor)
	return boilerSqliteDto.NewDomainRooms(rooms), boilerSqliteErrutil.HandleErr(err)
}

func (r room) SearchAllByName(ctx context.Context, keyword string, params selector1.Selector) (model.Rooms, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	mods := selector.ToBoilerSelectorPipe(
		params,
		sqlboiler.RoomWhere.Name.LIKE("%"+keyword+"%"),
	)
	rooms, err := sqlboiler.Rooms(mods...).All(ctx, r.executor)
	return boilerSqliteDto.NewDomainRooms(rooms), boilerSqliteErrutil.HandleErr(err)
}
