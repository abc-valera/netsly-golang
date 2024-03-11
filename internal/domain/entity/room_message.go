package entity

import (
	"context"
	"time"

	newbasemodel "github.com/abc-valera/netsly-api-golang/internal/domain/entity/new-base-model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
	"github.com/google/uuid"
)

type RoomMessage struct {
	newUUID func() string
	timeNow func() time.Time

	command command.IRoomMessage
}

func NewRoomMessage(
	command command.IRoomMessage,
) RoomMessage {
	return RoomMessage{
		newUUID: uuid.New().String,
		timeNow: time.Now,

		command: command,
	}
}

type RoomMessageCreateRequest struct {
	Text   string `validate:"required,min=1,max=2048"`
	UserID string `validate:"required,uuid"`
	RoomID string `validate:"required,uuid"`
}

func (rm RoomMessage) Create(ctx context.Context, req RoomMessageCreateRequest) (model.RoomMessage, error) {
	if err := global.Validator().Struct(req); err != nil {
		return model.RoomMessage{}, err
	}

	baseModel := newbasemodel.NewBaseModel(rm.newUUID(), rm.timeNow())

	return rm.command.Create(ctx, model.RoomMessage{
		BaseModel: baseModel,
		Text:      req.Text,
		UserID:    req.UserID,
		RoomID:    req.RoomID,
	})
}

type RoomMessageUpdateRequest struct {
	Text *string `validate:"min=1,max=2048"`
}

func (rm RoomMessage) Update(ctx context.Context, id string, req RoomMessageUpdateRequest) (model.RoomMessage, error) {
	if err := global.Validator().Struct(req); err != nil {
		return model.RoomMessage{}, err
	}

	return rm.command.Update(ctx, id, command.RoomMessageUpdate{
		Text: req.Text,
	})
}

func (rm RoomMessage) Delete(ctx context.Context, id string) error {
	if err := global.Validator().Var(id, "uuid"); err != nil {
		return err
	}

	return rm.command.Delete(ctx, id)
}
