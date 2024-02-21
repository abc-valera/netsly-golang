package command

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
)

type IChatMessage interface {
	Create(ctx context.Context, req model.ChatMessage) (model.ChatMessage, error)
	Update(ctx context.Context, id string, req ChatMessageUpdate) (model.ChatMessage, error)
	Delete(ctx context.Context, id string) error
}

type ChatMessageUpdate struct {
	Text *string
}
