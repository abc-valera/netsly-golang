package query

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/model"
)

type IChatMember interface {
	GetByIDs(ctx context.Context, userID, chatRoomID string) (model.ChatMember, error)
}
