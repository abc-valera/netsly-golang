package boilerQuery

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query"
	selector1 "github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query/selector"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/boiler/boilerDto"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/boiler/boilerQuery/selector"
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
	return boilerDto.ToDomainRoomMessageWithErrHandle(sqlboiler.FindRoomMessage(ctx, r.executor, id))
}

func (r roomMessage) GetAllByRoomID(ctx context.Context, roomID string, spec selector1.Selector) (model.RoomMessages, error) {
	mods := selector.ToBoilerSelectorPipe(
		spec,
		sqlboiler.RoomMessageWhere.RoomID.EQ(roomID),
	)
	return boilerDto.ToDomainRoomMessagesWithErrHandle(sqlboiler.RoomMessages(mods...).All(ctx, r.executor))
}

func (r roomMessage) SearchAllByText(ctx context.Context, keyword string, spec selector1.Selector) (model.RoomMessages, error) {
	mods := selector.ToBoilerSelectorPipe(
		spec,
		sqlboiler.RoomMessageWhere.Text.LIKE("%"+keyword+"%"),
	)
	return boilerDto.ToDomainRoomMessagesWithErrHandle(sqlboiler.RoomMessages(mods...).All(ctx, r.executor))
}
