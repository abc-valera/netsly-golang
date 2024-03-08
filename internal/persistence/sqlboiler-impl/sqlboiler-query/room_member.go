package sqlboilerquery

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-api-golang/internal/persistence/sqlboiler-impl/dto"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type roomMember struct {
	executor boil.ContextExecutor
}

func newRoomMember(executor boil.ContextExecutor) query.IRoomMember {
	return &roomMember{
		executor: executor,
	}
}

func (r roomMember) GetByRoomID(ctx context.Context, roomID string) (model.RoomMembers, error) {
	mods := sqlboiler.RoomMemberWhere.RoomID.EQ(roomID)
	return dto.ToDomainRoomMembersWithErrHandle(sqlboiler.RoomMembers(mods).All(ctx, r.executor))
}

func (r roomMember) GetByIDs(ctx context.Context, userID string, roomID string) (model.RoomMember, error) {
	return dto.ToDomainRoomMemberWithErrHandle(sqlboiler.FindRoomMember(ctx, r.executor, userID, roomID))
}