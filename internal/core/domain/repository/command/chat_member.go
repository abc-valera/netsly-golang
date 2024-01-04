package command

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/model"
)

type IChatMemberCommand interface {
	Create(ctx context.Context, req model.ChatMember) error
	Delete(ctx context.Context, ChatRoomID, UserID string) error
}
