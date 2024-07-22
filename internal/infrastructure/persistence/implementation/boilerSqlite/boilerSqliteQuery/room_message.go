package boilerSqliteQuery

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query"
	selector1 "github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query/selector"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/boilerSqlite/boilerSqliteDto"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/boilerSqlite/boilerSqliteQuery/selector"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type roomMessage struct {
	executor boil.ContextExecutor
}

func NewRoomMessage(executor boil.ContextExecutor) query.IRoomMessage {
	return &roomMessage{
		executor: executor,
	}
}

func (r roomMessage) GetByID(ctx context.Context, id string) (model.RoomMessage, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	roomMessage, err := sqlboiler.FindRoomMessage(ctx, r.executor, id)
	return boilerSqliteDto.NewDomainRoomMessage(roomMessage), err
}

func (r roomMessage) GetAllByRoomID(ctx context.Context, roomID string, spec selector1.Selector) (model.RoomMessages, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	mods := selector.ToBoilerSelectorPipe(
		spec,
		sqlboiler.RoomMessageWhere.RoomID.EQ(roomID),
	)
	roomMessages, err := sqlboiler.RoomMessages(mods...).All(ctx, r.executor)
	return boilerSqliteDto.NewDomainRoomMessages(roomMessages), err
}

func (r roomMessage) SearchAllByText(ctx context.Context, keyword string, spec selector1.Selector) (model.RoomMessages, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	mods := selector.ToBoilerSelectorPipe(
		spec,
		sqlboiler.RoomMessageWhere.Text.LIKE("%"+keyword+"%"),
	)
	roomMessages, err := sqlboiler.RoomMessages(mods...).All(ctx, r.executor)
	return boilerSqliteDto.NewDomainRoomMessages(roomMessages), err
}
