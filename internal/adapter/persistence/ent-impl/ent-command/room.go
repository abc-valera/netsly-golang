package entcommand

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/ent"
	"github.com/abc-valera/netsly-api-golang/internal/adapter/persistence/ent-impl/dto"
	errhandler "github.com/abc-valera/netsly-api-golang/internal/adapter/persistence/ent-impl/errors"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
)

type roomCommand struct {
	*ent.Client
}

func NewRoomCommand(client *ent.Client) command.IRoom {
	return &roomCommand{
		Client: client,
	}
}

func (cm roomCommand) Create(ctx context.Context, req model.Room) (model.Room, error) {
	room, err := cm.Room.Create().
		SetID(req.ID).
		SetName(req.Name).
		SetCreatorID(req.CreatorID).
		SetDescription(req.Description).
		SetCreatedAt(req.CreatedAt).
		Save(ctx)
	return dto.FromEntRoom(room), errhandler.HandleErr(err)
}

func (cm roomCommand) Update(ctx context.Context, id string, req command.RoomUpdate) (model.Room, error) {
	query := cm.Room.UpdateOneID(id)
	if req.Description != nil {
		query.SetDescription(*req.Description)
	}

	room, err := query.
		Save(ctx)

	return dto.FromEntRoom(room), errhandler.HandleErr(err)
}

func (cm roomCommand) Delete(ctx context.Context, id string) error {
	err := cm.Room.
		DeleteOneID(id).
		Exec(ctx)
	return errhandler.HandleErr(err)
}
