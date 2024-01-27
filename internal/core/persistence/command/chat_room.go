package command

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/core/model"
)

type IChatRoom interface {
	Create(ctx context.Context, req model.ChatRoom) error
	Update(ctx context.Context, id string, req ChatRoomUpdate) error
	Delete(ctx context.Context, id string) error
}

type ChatRoomUpdate struct {
	Description *string
}
