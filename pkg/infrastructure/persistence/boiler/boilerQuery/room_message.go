package boilerQuery

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/model"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/query"
	selector1 "github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/query/selector"
	"github.com/abc-valera/netsly-api-golang/pkg/infrastructure/persistence/boiler/boilerQuery/selector"
	"github.com/abc-valera/netsly-api-golang/pkg/infrastructure/persistence/boiler/dto"
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
	return dto.ToDomainRoomMessageWithErrHandle(sqlboiler.FindRoomMessage(ctx, r.executor, id))
}

func (r roomMessage) GetAllByRoomID(ctx context.Context, roomID string, spec selector1.Selector) (model.RoomMessages, error) {
	mods := selector.ToBoilerSelectorPipe(
		spec,
		sqlboiler.RoomMessageWhere.RoomID.EQ(roomID),
	)
	return dto.ToDomainRoomMessagesWithErrHandle(sqlboiler.RoomMessages(mods...).All(ctx, r.executor))
}

func (r roomMessage) SearchAllByText(ctx context.Context, keyword string, spec selector1.Selector) (model.RoomMessages, error) {
	mods := selector.ToBoilerSelectorPipe(
		spec,
		sqlboiler.RoomMessageWhere.Text.LIKE("%"+keyword+"%"),
	)
	return dto.ToDomainRoomMessagesWithErrHandle(sqlboiler.RoomMessages(mods...).All(ctx, r.executor))
}
