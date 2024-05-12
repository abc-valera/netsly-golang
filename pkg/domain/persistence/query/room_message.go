package query

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/pkg/domain/model"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/query/selector"
)

type IRoomMessage interface {
	GetByID(ctx context.Context, id string) (model.RoomMessage, error)
	GetAllByRoomID(ctx context.Context, roomID string, spec selector.Selector) (model.RoomMessages, error)
	SearchAllByText(ctx context.Context, keyword string, spec selector.Selector) (model.RoomMessages, error)
}
