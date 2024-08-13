package boilerSqliteCommand

import (
	"context"

	"github.com/abc-valera/netsly-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-golang/internal/core/global"
	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/boilerSqlite/boilerSqliteDto"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/boilerSqlite/boilerSqliteErrutil"
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

func (c roomMessage) Create(ctx context.Context, req command.RoomMessageCreateRequest) (model.RoomMessage, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	roomMessage := sqlboiler.RoomMessage{
		ID:        req.RoomMessage.ID,
		Text:      req.RoomMessage.Text,
		CreatedAt: req.RoomMessage.CreatedAt,
		UpdatedAt: req.RoomMessage.UpdatedAt,
		DeletedAt: req.RoomMessage.DeletedAt,

		UserID: req.UserID,
		RoomID: req.RoomID,
	}
	err := roomMessage.Insert(ctx, c.executor, boil.Infer())
	return boilerSqliteDto.NewDomainRoomMessage(&roomMessage), boilerSqliteErrutil.HandleErr(err)
}

func (c roomMessage) Update(ctx context.Context, id string, req command.RoomMessageUpdateRequest) (model.RoomMessage, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	roomMessage, err := sqlboiler.FindRoomMessage(ctx, c.executor, id)
	if err != nil {
		return model.RoomMessage{}, boilerSqliteErrutil.HandleErr(err)
	}
	if req.Text != nil {
		roomMessage.Text = *req.Text
	}
	_, err = roomMessage.Update(ctx, c.executor, boil.Infer())
	return boilerSqliteDto.NewDomainRoomMessage(roomMessage), boilerSqliteErrutil.HandleErr(err)
}

func (c roomMessage) Delete(ctx context.Context, id string) error {
	_, span := global.NewSpan(ctx)
	defer span.End()

	roomMessage, err := sqlboiler.FindRoomMessage(ctx, c.executor, id)
	if err != nil {
		return boilerSqliteErrutil.HandleErr(err)
	}
	_, err = roomMessage.Delete(ctx, c.executor)
	return boilerSqliteErrutil.HandleErr(err)
}
