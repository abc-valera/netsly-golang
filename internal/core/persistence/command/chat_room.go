package command

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/core/persistence/model"
)

var (
	ErrChatRoomNameAlreadyExists = coderr.NewMessage(coderr.CodeAlreadyExists, "ChatRoom already exists")
)

type IChatRoom interface {
	Create(ctx context.Context, req model.ChatRoom) error
	Update(ctx context.Context, id string, req ChatRoomUpdate) error
	Delete(ctx context.Context, id string) error
}

type ChatRoomUpdate struct {
	Description *string
}
