package entcommand

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/ent"
	"github.com/abc-valera/netsly-api-golang/gen/ent/chatroom"
	errhandler "github.com/abc-valera/netsly-api-golang/internal/adapter/persistence/ent-impl/errors"
	"github.com/abc-valera/netsly-api-golang/internal/core/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/core/persistence/model"
)

type chatRoomCommand struct {
	*ent.Client
}

func NewChatRoomCommand(client *ent.Client) command.IChatRoom {
	return &chatRoomCommand{
		Client: client,
	}
}

func (cm chatRoomCommand) Create(ctx context.Context, req model.ChatRoom) error {
	_, err := cm.ChatRoom.Create().
		SetID(req.ID).
		SetName(req.Name).
		SetDescription(req.Description).
		SetCreatedAt(req.CreatedAt).
		Save(ctx)
	return err
}

func (cm chatRoomCommand) Update(ctx context.Context, id string, req command.ChatRoomUpdate) error {
	query := cm.ChatRoom.Update()
	if req.Description != nil {
		query.SetDescription(*req.Description)
	}

	_, err := query.
		Where(chatroom.ID(id)).
		Save(ctx)

	return err
}

func (cm chatRoomCommand) Delete(ctx context.Context, id string) error {
	err := cm.ChatRoom.
		DeleteOneID(id).
		Exec(ctx)
	return errhandler.HandleErr(err)
}
