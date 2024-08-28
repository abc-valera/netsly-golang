package bunSqliteCommand

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/bunSqlite/bunSqliteDto"
	bunSqlitErrutil "github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/bunSqlite/bunSqliteErrutil"
	"github.com/uptrace/bun"
)

type roomMessage struct {
	db bun.IDB
}

func NewRoomMessage(db bun.IDB) command.IRoomMessage {
	return &roomMessage{
		db: db,
	}
}

func (c roomMessage) Create(ctx context.Context, req command.RoomMessageCreateRequest) (model.RoomMessage, error) {
	roomMessage := bunSqliteDto.RoomMessage{
		ID:        req.RoomMessage.ID,
		Text:      req.RoomMessage.Text,
		CreatedAt: req.RoomMessage.CreatedAt,
		UpdatedAt: req.RoomMessage.UpdatedAt,
		DeletedAt: req.RoomMessage.DeletedAt,

		UserID: req.UserID,
		RoomID: req.RoomID,
	}

	res, err := c.db.NewInsert().Model(&roomMessage).Exec(ctx)
	return roomMessage.ToDomain(), bunSqlitErrutil.HandleCommandResult(res, err)
}

func (c roomMessage) Update(ctx context.Context, id string, req command.RoomMessageUpdateRequest) (model.RoomMessage, error) {
	roomMessage := bunSqliteDto.RoomMessage{
		ID: id,
	}
	var columns []string

	if req.Text != nil {
		roomMessage.Text = *req.Text
		columns = append(columns, "text")
	}

	if len(columns) == 0 {
		return model.RoomMessage{}, nil
	}

	res, err := c.db.NewUpdate().Model(&roomMessage).Column(columns...).WherePK().Exec(ctx)
	return roomMessage.ToDomain(), bunSqlitErrutil.HandleCommandResult(res, err)
}

func (c roomMessage) Delete(ctx context.Context, id string) error {
	roomMessage := bunSqliteDto.RoomMessage{
		ID: id,
	}
	res, err := c.db.NewDelete().Model(&roomMessage).WherePK().Exec(ctx)
	return bunSqlitErrutil.HandleCommandResult(res, err)
}
