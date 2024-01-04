package query

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/model"
)

type IChatMessageQuery interface {
	GetByID(ctx context.Context, id string) (*model.ChatMessage, error)
	GetAllByChatRoomID(ctx context.Context, chatRoomID string) (model.ChatMessages, error)
	SearchAllByText(ctx context.Context, keyword string) (model.ChatMessages, error)
}
