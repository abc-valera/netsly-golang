package entcommand

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/ent"
	"github.com/abc-valera/netsly-api-golang/gen/ent/chatmember"
	"github.com/abc-valera/netsly-api-golang/internal/adapter/persistence/ent-impl/dto"
	errhandler "github.com/abc-valera/netsly-api-golang/internal/adapter/persistence/ent-impl/errors"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
)

type chatMemberCommand struct {
	*ent.Client
}

func NewChatMemberCommand(client *ent.Client) command.IChatMember {
	return &chatMemberCommand{
		Client: client,
	}
}

func (cm chatMemberCommand) Create(ctx context.Context, req model.ChatMember) (model.ChatMember, error) {
	chatMember, err := cm.ChatMember.Create().
		SetChatRoomID(req.ChatRoomID).
		SetUserID(req.UserID).
		SetCreatedAt(req.CreatedAt).
		Save(ctx)
	return dto.FromEntChatMember(chatMember), errhandler.HandleErr(err)
}

func (cm chatMemberCommand) Delete(ctx context.Context, ChatRoomID string, UserID string) error {
	_, err := cm.ChatMember.Delete().
		Where(
			chatmember.ChatRoomIDEQ(ChatRoomID),
			chatmember.UserID(UserID),
		).Exec(ctx)
	return errhandler.HandleErr(err)
}
