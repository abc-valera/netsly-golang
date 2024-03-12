package entity

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/service"
)

type Room struct {
	command command.IRoom

	uuidMaker service.IUuidMaker
	timeMaker service.ITimeMaker
}

func NewRoom(
	command command.IRoom,
	uuidMaker service.IUuidMaker,
	timeMaker service.ITimeMaker,
) Room {
	return Room{
		command:   command,
		uuidMaker: uuidMaker,
		timeMaker: timeMaker,
	}
}

type RoomCreateRequest struct {
	Name        string `validate:"required,min=4,max=64"`
	Description string `validate:"max=256"`
	CreatorID   string `validate:"required,uuid"`
}

func (r Room) Create(ctx context.Context, req RoomCreateRequest) (model.Room, error) {
	if err := global.Validator().Struct(req); err != nil {
		return model.Room{}, err
	}

	return r.command.Create(ctx, model.Room{
		ID:          r.uuidMaker.NewUUID(),
		Name:        req.Name,
		Description: req.Description,
		CreatedAt:   r.timeMaker.Now(),
		CreatorID:   req.CreatorID,
	})
}

type RoomUpdateRequest struct {
	Name        *string `validate:"min=4,max=64"`
	Description *string `validate:"max=256"`
}

func (r Room) Update(ctx context.Context, roomID string, req RoomUpdateRequest) (model.Room, error) {
	if err := global.Validator().Struct(req); err != nil {
		return model.Room{}, err
	}

	return r.command.Update(ctx, roomID, command.RoomUpdate{
		Description: req.Description,
	})
}

func (r Room) Delete(ctx context.Context, roomID string) error {
	if err := global.Validator().Var(roomID, "uuid"); err != nil {
		return err
	}

	return r.command.Delete(ctx, roomID)
}
