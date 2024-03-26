package sqlboilerCommand

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/model"
	"github.com/abc-valera/netsly-api-golang/pkg/persistence/sqlboilerImpl/dto"
	"github.com/abc-valera/netsly-api-golang/pkg/persistence/sqlboilerImpl/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type roomMessage struct {
	executor boil.ContextExecutor
}

func newRoomMessage(executor boil.ContextExecutor) command.IRoomMessage {
	return &roomMessage{
		executor: executor,
	}
}

func (r roomMessage) Create(ctx context.Context, req model.RoomMessage) (model.RoomMessage, error) {
	roomMessage := sqlboiler.RoomMessage{
		ID:        req.ID,
		Text:      req.Text,
		CreatedAt: req.CreatedAt,
		UserID:    req.UserID,
		RoomID:    req.RoomID,
	}
	err := roomMessage.Insert(ctx, r.executor, boil.Infer())
	return dto.ToDomainRoomMessageWithErrHandle(&roomMessage, err)
}

func (r roomMessage) Update(ctx context.Context, id string, req command.RoomMessageUpdate) (model.RoomMessage, error) {
	roomMessage, err := sqlboiler.FindRoomMessage(ctx, r.executor, id)
	if err != nil {
		return model.RoomMessage{}, errors.HandleErr(err)
	}
	if req.Text != nil {
		roomMessage.Text = *req.Text
	}
	_, err = roomMessage.Update(ctx, r.executor, boil.Infer())
	return dto.ToDomainRoomMessageWithErrHandle(roomMessage, err)
}

func (r roomMessage) Delete(ctx context.Context, id string) error {
	roomMessage, err := sqlboiler.FindRoomMessage(ctx, r.executor, id)
	if err != nil {
		return errors.HandleErr(err)
	}
	_, err = roomMessage.Delete(ctx, r.executor)
	return errors.HandleErr(err)
}
