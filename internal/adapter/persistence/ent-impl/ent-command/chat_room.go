package entcommand

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/ent"
	"github.com/abc-valera/netsly-api-golang/internal/adapter/persistence/ent-impl/dto"
	errhandler "github.com/abc-valera/netsly-api-golang/internal/adapter/persistence/ent-impl/errors"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
)

type chatRoomCommand struct {
	*ent.Client
}

func NewChatRoomCommand(client *ent.Client) command.IChatRoom {
	return &chatRoomCommand{
		Client: client,
	}
}

func (cm chatRoomCommand) Create(ctx context.Context, req model.ChatRoom) (model.ChatRoom, error) {
	room, err := cm.ChatRoom.Create().
		SetID(req.ID).
		SetName(req.Name).
		SetDescription(req.Description).
		SetCreatedAt(req.CreatedAt).
		Save(ctx)
	return dto.FromEntChatRoom(room), errhandler.HandleErr(err)
}

func (cm chatRoomCommand) Update(ctx context.Context, id string, req command.ChatRoomUpdate) (model.ChatRoom, error) {
	query := cm.ChatRoom.UpdateOneID(id)
	if req.Description != nil {
		query.SetDescription(*req.Description)
	}

	room, err := query.
		Save(ctx)

	return dto.FromEntChatRoom(room), errhandler.HandleErr(err)
}

func (cm chatRoomCommand) Delete(ctx context.Context, id string) error {
	err := cm.ChatRoom.
		DeleteOneID(id).
		Exec(ctx)
	return errhandler.HandleErr(err)
}
