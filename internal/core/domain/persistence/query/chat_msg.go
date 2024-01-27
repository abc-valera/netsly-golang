package query

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/model"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/persistence/query/spec"
)

type IChatMessage interface {
	GetByID(ctx context.Context, id string) (model.ChatMessage, error)
	GetAllByChatRoomID(ctx context.Context, chatRoomID string) (model.ChatMessages, error)
	SearchAllByText(ctx context.Context, keyword string, spec spec.SelectParams) (model.ChatMessages, error)
}
