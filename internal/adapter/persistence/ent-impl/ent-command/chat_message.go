package entcommand

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/ent"
	"github.com/abc-valera/netsly-api-golang/gen/ent/chatmessage"
	errhandler "github.com/abc-valera/netsly-api-golang/internal/adapter/persistence/ent-impl/errors"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
)

type chatMessageCommand struct {
	*ent.Client
}

func NewChatMessageCommand(client *ent.Client) command.IChatMessage {
	return &chatMessageCommand{
		Client: client,
	}
}

func (cm chatMessageCommand) Create(ctx context.Context, req model.ChatMessage) error {
	_, err := cm.ChatMessage.Create().
		SetID(req.ID).
		SetChatRoomID(req.ChatRoomID).
		SetUserID(req.UserID).
		SetText(req.Text).
		SetCreatedAt(req.CreatedAt).
		Save(ctx)
	return errhandler.HandleErr(err)
}

func (cm chatMessageCommand) Update(ctx context.Context, id string, req command.ChatMessageUpdate) error {
	query := cm.ChatMessage.Update()
	if req.Text != nil {
		query.SetText(*req.Text)
	}

	_, err := query.
		Where(chatmessage.ID(id)).
		Save(ctx)

	return errhandler.HandleErr(err)
}

func (cm chatMessageCommand) Delete(ctx context.Context, id string) error {
	err := cm.ChatMessage.
		DeleteOneID(id).
		Exec(ctx)
	return errhandler.HandleErr(err)
}
