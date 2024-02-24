package query

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query/spec"
)

type IRoomMessage interface {
	GetByID(ctx context.Context, id string) (model.RoomMessage, error)
	GetAllByRoomID(ctx context.Context, roomID string) (model.RoomMessages, error)
	SearchAllByText(ctx context.Context, keyword string, spec spec.SelectParams) (model.RoomMessages, error)
}
