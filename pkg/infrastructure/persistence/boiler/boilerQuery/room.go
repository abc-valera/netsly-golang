package boilerQuery

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/model"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/query"
	selectParams1 "github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/query/selectParams"
	"github.com/abc-valera/netsly-api-golang/pkg/infrastructure/persistence/boiler/boilerQuery/selectParams"
	"github.com/abc-valera/netsly-api-golang/pkg/infrastructure/persistence/boiler/dto"
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
	return dto.ToDomainRoomWithErrHandle(sqlboiler.FindRoom(ctx, r.executor, id))
}

func (r room) GetByName(ctx context.Context, name string) (model.Room, error) {
	return dto.ToDomainRoomWithErrHandle(sqlboiler.Rooms(sqlboiler.RoomWhere.Name.EQ(name)).One(ctx, r.executor))
}

func (r room) GetAllByUserID(ctx context.Context, userID string, params selectParams1.SelectParams) (model.Rooms, error) {
	mods := selectParams.ToBoilerSelectParamsPipe(
		params,
		qm.InnerJoin(sqlboiler.TableNames.RoomMember+" rm ON rm.room_id = room.id"),
		sqlboiler.RoomMemberWhere.UserID.EQ(userID),
	)
	return dto.ToDomainRoomsWithErrHandle(sqlboiler.Rooms(mods...).All(ctx, r.executor))
}

func (r room) SearchAllByName(ctx context.Context, keyword string, params selectParams1.SelectParams) (model.Rooms, error) {
	mods := selectParams.ToBoilerSelectParamsPipe(
		params,
		sqlboiler.RoomWhere.Name.LIKE("%"+keyword+"%"),
	)
	return dto.ToDomainRoomsWithErrHandle(sqlboiler.Rooms(mods...).All(ctx, r.executor))
}