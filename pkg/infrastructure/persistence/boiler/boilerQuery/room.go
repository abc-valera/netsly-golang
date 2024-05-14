package boilerQuery

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/model"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/query"
	selector1 "github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/query/selector"
	"github.com/abc-valera/netsly-api-golang/pkg/infrastructure/persistence/boiler/boilerDto"
	"github.com/abc-valera/netsly-api-golang/pkg/infrastructure/persistence/boiler/boilerQuery/selector"
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
	return boilerDto.ToDomainRoomWithErrHandle(sqlboiler.FindRoom(ctx, r.executor, id))
}

func (r room) GetByName(ctx context.Context, name string) (model.Room, error) {
	return boilerDto.ToDomainRoomWithErrHandle(sqlboiler.Rooms(sqlboiler.RoomWhere.Name.EQ(name)).One(ctx, r.executor))
}

func (r room) GetAllByUserID(ctx context.Context, userID string, params selector1.Selector) (model.Rooms, error) {
	mods := selector.ToBoilerSelectorPipe(
		params,
		qm.InnerJoin(sqlboiler.TableNames.RoomMember+" rm ON rm.room_id = room.id"),
		sqlboiler.RoomMemberWhere.UserID.EQ(userID),
	)
	return boilerDto.ToDomainRoomsWithErrHandle(sqlboiler.Rooms(mods...).All(ctx, r.executor))
}

func (r room) SearchAllByName(ctx context.Context, keyword string, params selector1.Selector) (model.Rooms, error) {
	mods := selector.ToBoilerSelectorPipe(
		params,
		sqlboiler.RoomWhere.Name.LIKE("%"+keyword+"%"),
	)
	return boilerDto.ToDomainRoomsWithErrHandle(sqlboiler.Rooms(mods...).All(ctx, r.executor))
}
