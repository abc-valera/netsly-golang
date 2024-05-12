package query

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/model"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/query/selector"
)

type IRoom interface {
	GetByID(ctx context.Context, id string) (model.Room, error)
	GetByName(ctx context.Context, name string) (model.Room, error)
	GetAllByUserID(ctx context.Context, userID string, params selector.Selector) (model.Rooms, error)
	SearchAllByName(ctx context.Context, keyword string, params selector.Selector) (model.Rooms, error)
}
