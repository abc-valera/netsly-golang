package command

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/core/model"
)

type IChatMember interface {
	Create(ctx context.Context, req model.ChatMember) error
	Delete(ctx context.Context, ChatRoomID, UserID string) error
}
