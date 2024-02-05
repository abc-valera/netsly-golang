package query

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
)

type IChatMember interface {
	GetByIDs(ctx context.Context, userID, chatRoomID string) (model.ChatMember, error)
}
