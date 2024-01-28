package command

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/core/coderr"
	"github.com/abc-valera/flugo-api-golang/internal/core/persistence/model"
)

var (
	ErrChatMemberAlreadyExists = coderr.NewMessage(coderr.CodeAlreadyExists, "ChatMember already exists")
)

type IChatMember interface {
	Create(ctx context.Context, req model.ChatMember) error
	Delete(ctx context.Context, ChatRoomID, UserID string) error
}
