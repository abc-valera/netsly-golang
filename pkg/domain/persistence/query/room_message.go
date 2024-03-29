package query

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/model"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/query/selectParams"
)

type IRoomMessage interface {
	GetByID(ctx context.Context, id string) (model.RoomMessage, error)
	GetAllByRoomID(ctx context.Context, roomID string, spec selectParams.SelectParams) (model.RoomMessages, error)
	SearchAllByText(ctx context.Context, keyword string, spec selectParams.SelectParams) (model.RoomMessages, error)
}
