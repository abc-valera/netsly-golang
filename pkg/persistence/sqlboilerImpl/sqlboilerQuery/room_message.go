package sqlboilerQuery

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/model"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/query"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/query/spec"
	"github.com/abc-valera/netsly-api-golang/pkg/persistence/sqlboilerImpl/dto"
	"github.com/abc-valera/netsly-api-golang/pkg/persistence/sqlboilerImpl/sqlboilerQuery/common"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type roomMessage struct {
	executor boil.ContextExecutor
}

func newRoomMessage(executor boil.ContextExecutor) query.IRoomMessage {
	return &roomMessage{
		executor: executor,
	}
}

func (r roomMessage) GetByID(ctx context.Context, id string) (model.RoomMessage, error) {
	return dto.ToDomainRoomMessageWithErrHandle(sqlboiler.FindRoomMessage(ctx, r.executor, id))
}

func (r roomMessage) GetAllByRoomID(ctx context.Context, roomID string, spec spec.SelectParams) (model.RoomMessages, error) {
	mods := common.ToBoilerSelectParamsPipe(
		spec,
		sqlboiler.RoomMessageWhere.RoomID.EQ(roomID),
	)
	return dto.ToDomainRoomMessagesWithErrHandle(sqlboiler.RoomMessages(mods...).All(ctx, r.executor))
}

func (r roomMessage) SearchAllByText(ctx context.Context, keyword string, spec spec.SelectParams) (model.RoomMessages, error) {
	mods := common.ToBoilerSelectParamsPipe(
		spec,
		sqlboiler.RoomMessageWhere.Text.LIKE("%"+keyword+"%"),
	)
	return dto.ToDomainRoomMessagesWithErrHandle(sqlboiler.RoomMessages(mods...).All(ctx, r.executor))
}
