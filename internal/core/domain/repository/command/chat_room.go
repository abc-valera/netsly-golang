package command

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/model"
)

type IChatRoomCommand interface {
	Create(ctx context.Context, req model.ChatRoom) error
	Update(ctx context.Context, id string, req model.ChatRoomUpdate) error
	Delete(ctx context.Context, id string) error
}
