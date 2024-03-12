package entity

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/service"
)

type RoomMessage struct {
	command command.IRoomMessage

	uuidMaker service.IUuidMaker
	timeMaker service.ITimeMaker
}

func NewRoomMessage(
	command command.IRoomMessage,
	uuidMaker service.IUuidMaker,
	timeMaker service.ITimeMaker,
) RoomMessage {
	return RoomMessage{
		command:   command,
		uuidMaker: uuidMaker,
		timeMaker: timeMaker,
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

	return rm.command.Create(ctx, model.RoomMessage{
		ID:        rm.uuidMaker.NewUUID(),
		Text:      req.Text,
		CreatedAt: rm.timeMaker.Now(),
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
