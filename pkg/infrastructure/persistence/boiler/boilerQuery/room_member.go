package boilerQuery

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/model"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/query"
	"github.com/abc-valera/netsly-api-golang/pkg/infrastructure/persistence/boiler/boilerDto"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type roomMember struct {
	executor boil.ContextExecutor
}

func NewRoomMember(executor boil.ContextExecutor) query.IRoomMember {
	return &roomMember{
		executor: executor,
	}
}

func (r roomMember) GetByRoomID(ctx context.Context, roomID string) (model.RoomMembers, error) {
	mods := sqlboiler.RoomMemberWhere.RoomID.EQ(roomID)
	return boilerDto.ToDomainRoomMembersWithErrHandle(sqlboiler.RoomMembers(mods).All(ctx, r.executor))
}

func (r roomMember) GetByIDs(ctx context.Context, userID string, roomID string) (model.RoomMember, error) {
	return boilerDto.ToDomainRoomMemberWithErrHandle(sqlboiler.FindRoomMember(ctx, r.executor, userID, roomID))
}
