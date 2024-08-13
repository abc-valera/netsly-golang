package query

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/selector"
)

type IRoomMessage interface {
	GetByID(ctx context.Context, id string) (model.RoomMessage, error)
	GetAllByRoomID(ctx context.Context, roomID string, selector selector.Selector) (model.RoomMessages, error)
	SearchAllByText(ctx context.Context, keyword string, selector selector.Selector) (model.RoomMessages, error)
}
