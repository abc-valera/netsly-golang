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

func (r roomMessage) GetAllByRoomID(ctx context.Context, roomID string, spec selectParams1.SelectParams) (model.RoomMessages, error) {
	mods := selectParams.ToBoilerSelectParamsPipe(
		spec,
		sqlboiler.RoomMessageWhere.RoomID.EQ(roomID),
	)
	return dto.ToDomainRoomMessagesWithErrHandle(sqlboiler.RoomMessages(mods...).All(ctx, r.executor))
}

func (r roomMessage) SearchAllByText(ctx context.Context, keyword string, spec selectParams1.SelectParams) (model.RoomMessages, error) {
	mods := selectParams.ToBoilerSelectParamsPipe(
		spec,
		sqlboiler.RoomMessageWhere.Text.LIKE("%"+keyword+"%"),
	)
	return dto.ToDomainRoomMessagesWithErrHandle(sqlboiler.RoomMessages(mods...).All(ctx, r.executor))
}
