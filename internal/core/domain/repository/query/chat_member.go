package query

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/model"
)

type IChatMemberQuery interface {
	GetByUserIDChatRoomID(ctx context.Context, userID, chatRoomID string) (*model.ChatMember, error)
}
