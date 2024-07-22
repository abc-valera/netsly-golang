package boilerSqliteQuery

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query/selector"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/boilerSqlite/boilerSqliteDto"
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

func (r roomMember) GetAllByRoomID(ctx context.Context, roomID string, selector selector.Selector) (model.RoomMembers, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	roomMembers, err := sqlboiler.RoomMembers(
		sqlboiler.RoomMemberWhere.RoomID.EQ(roomID),
	).All(ctx, r.executor)
	return boilerSqliteDto.NewDomainRoomMembers(roomMembers), err
}

func (r roomMember) GetByIDs(ctx context.Context, userID string, roomID string) (model.RoomMember, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	roomMembers, err := sqlboiler.FindRoomMember(ctx, r.executor, userID, roomID)
	return boilerSqliteDto.NewDomainRoomMember(roomMembers), err
}
