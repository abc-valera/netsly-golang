package boilerCommand

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/boiler/boilerDto"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/boiler/errutil"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type roomMessage struct {
	executor boil.ContextExecutor
}

func NewRoomMessage(executor boil.ContextExecutor) command.IRoomMessage {
	return &roomMessage{
		executor: executor,
	}
}

func (c roomMessage) Create(ctx context.Context, userID, roomID string, req model.RoomMessage) (model.RoomMessage, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	roomMessage := sqlboiler.RoomMessage{
		ID:        req.ID,
		Text:      req.Text,
		CreatedAt: req.CreatedAt,

		UserID: userID,
		RoomID: roomID,
	}
	err := roomMessage.Insert(ctx, c.executor, boil.Infer())
	return boilerDto.NewDomainRoomMessageWithErrHandle(&roomMessage, err)
}

func (c roomMessage) Update(ctx context.Context, id string, req command.RoomMessageUpdate) (model.RoomMessage, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	roomMessage, err := sqlboiler.FindRoomMessage(ctx, c.executor, id)
	if err != nil {
		return model.RoomMessage{}, errutil.HandleErr(err)
	}
	if req.Text != nil {
		roomMessage.Text = *req.Text
	}
	_, err = roomMessage.Update(ctx, c.executor, boil.Infer())
	return boilerDto.NewDomainRoomMessageWithErrHandle(roomMessage, err)
}

func (c roomMessage) Delete(ctx context.Context, id string) error {
	_, span := global.NewSpan(ctx)
	defer span.End()

	roomMessage, err := sqlboiler.FindRoomMessage(ctx, c.executor, id)
	if err != nil {
		return errutil.HandleErr(err)
	}
	_, err = roomMessage.Delete(ctx, c.executor)
	return errutil.HandleErr(err)
}
