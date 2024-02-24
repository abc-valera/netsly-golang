package entcommand

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/ent"
	"github.com/abc-valera/netsly-api-golang/internal/adapter/persistence/ent-impl/dto"
	errhandler "github.com/abc-valera/netsly-api-golang/internal/adapter/persistence/ent-impl/errors"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
)

type roomMessageCommand struct {
	*ent.Client
}

func NewRoomMessageCommand(client *ent.Client) command.IRoomMessage {
	return &roomMessageCommand{
		Client: client,
	}
}

func (cm roomMessageCommand) Create(ctx context.Context, req model.RoomMessage) (model.RoomMessage, error) {
	msg, err := cm.RoomMessage.Create().
		SetID(req.ID).
		SetRoomID(req.RoomID).
		SetUserID(req.UserID).
		SetText(req.Text).
		SetCreatedAt(req.CreatedAt).
		Save(ctx)
	return dto.FromEntRoomMessage(msg), errhandler.HandleErr(err)
}

func (cm roomMessageCommand) Update(ctx context.Context, id string, req command.RoomMessageUpdate) (model.RoomMessage, error) {
	query := cm.RoomMessage.UpdateOneID(id)
	if req.Text != nil {
		query.SetText(*req.Text)
	}

	msg, err := query.
		Save(ctx)

	return dto.FromEntRoomMessage(msg), errhandler.HandleErr(err)
}

func (cm roomMessageCommand) Delete(ctx context.Context, id string) error {
	err := cm.RoomMessage.
		DeleteOneID(id).
		Exec(ctx)
	return errhandler.HandleErr(err)
}
