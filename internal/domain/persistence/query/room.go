package query

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/selector"
)

type IRoom interface {
	GetByID(ctx context.Context, id string) (model.Room, error)
	GetByName(ctx context.Context, name string) (model.Room, error)
	GetAllByUserID(ctx context.Context, userID string, s selector.Selector) (model.Rooms, error)
	SearchAllByName(ctx context.Context, keyword string, s selector.Selector) (model.Rooms, error)
}
