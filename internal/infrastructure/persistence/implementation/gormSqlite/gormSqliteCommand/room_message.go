package gormSqliteCommand

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/gormSqlite/gormSqliteDto"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/gormSqlite/gormSqliteErrutil"
	"gorm.io/gorm"
)

type roomMessage struct {
	db *gorm.DB
}

func NewRoomMessage(db *gorm.DB) command.IRoomMessage {
	return &roomMessage{
		db: db,
	}
}

func (c roomMessage) Create(ctx context.Context, req command.RoomMessageCreateRequest) (model.RoomMessage, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	roomMessage := gormSqliteDto.RoomMessage{
		ID:        req.RoomMessage.ID,
		Text:      req.RoomMessage.Text,
		CreatedAt: req.RoomMessage.CreatedAt,
		UpdatedAt: req.RoomMessage.UpdatedAt,
		DeletedAt: req.RoomMessage.DeletedAt,
		UserID:    req.UserID,
		RoomID:    req.RoomID,
	}
	res := c.db.Create(&roomMessage)
	return gormSqliteDto.NewDomainRoomMessage(roomMessage), gormSqliteErrutil.HandleCommandResult(res)
}

func (c roomMessage) Update(ctx context.Context, id string, req command.RoomMessageUpdateRequest) (model.RoomMessage, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	var roomMessage gormSqliteDto.RoomMessage
	queryRes := c.db.Where("id = ?", id).First(&roomMessage)
	if err := gormSqliteErrutil.HandleQueryResult(queryRes); err != nil {
		return model.RoomMessage{}, err
	}

	if req.Text != nil {
		roomMessage.Text = *req.Text
	}

	updateRes := c.db.Save(&roomMessage)
	return gormSqliteDto.NewDomainRoomMessage(roomMessage), gormSqliteErrutil.HandleCommandResult(updateRes)
}

func (c roomMessage) Delete(ctx context.Context, id string) error {
	_, span := global.NewSpan(ctx)
	defer span.End()

	roomMessage := gormSqliteDto.RoomMessage{
		ID: id,
	}
	res := c.db.Delete(&roomMessage)
	return gormSqliteErrutil.HandleCommandResult(res)
}
