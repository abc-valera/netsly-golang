package repository

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/domain/entity"
)

type IChatMsgRepository interface {
	GetByID(ctx context.Context, id string) (*entity.ChatMsg, error)
	GetByChatRoomID(ctx context.Context, chatRoomID string) (entity.ChatMsgs, error)
	Create(ctx context.Context, chatMsg *entity.ChatMsg) error
	Update(ctx context.Context, chatMsg *entity.ChatMsg) error
	Delete(ctx context.Context, chatMsgID string) error
}
