package query

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query/spec"
)

type IChatRoom interface {
	GetByID(ctx context.Context, id string) (model.ChatRoom, error)
	GetByName(ctx context.Context, name string) (model.ChatRoom, error)
	GetAllByUserID(ctx context.Context, userID string, params spec.SelectParams) (model.ChatRooms, error)
	SearchAllByName(ctx context.Context, keyword string, params spec.SelectParams) (model.ChatRooms, error)
}
