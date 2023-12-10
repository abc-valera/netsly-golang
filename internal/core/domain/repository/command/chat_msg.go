package command

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/model"
)

type IChatMessageCommand interface {
	Create(ctx context.Context, req model.ChatMessage) error
	Update(ctx context.Context, id string, req model.ChatMessageUpdate) error
	Delete(ctx context.Context, id string) error
}
