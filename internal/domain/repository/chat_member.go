package repository

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/domain/entity"
)

type IChatMemberRepository interface {
	GetByChatRoomIDAndUserID(ctx context.Context, chatRoomID, userID string) (*entity.ChatMember, error)
	GetByChatRoomID(ctx context.Context, chatRoomID string) (entity.ChatMembers, error)
	GetByUserID(ctx context.Context, userID string) (entity.ChatMembers, error)
	Create(ctx context.Context, chatMember *entity.ChatMember) error
	Update(ctx context.Context, chatMember *entity.ChatMember) error
	Delete(ctx context.Context, chatMember *entity.ChatMember) error
}
