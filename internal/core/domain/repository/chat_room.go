package repository

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/entity"
)

type IChatRoomRepository interface {
	GetByID(ctx context.Context, id string) (*entity.ChatRoom, error)
	GetByUserID(ctx context.Context, userID string) (entity.ChatRooms, error)
	GetAll(ctx context.Context) (entity.ChatRooms, error)
	Create(ctx context.Context, chatRoom *entity.ChatRoom) error
	Update(ctx context.Context, chatRoom *entity.ChatRoom) error
	Delete(ctx context.Context, chatRoomID string) error
}
