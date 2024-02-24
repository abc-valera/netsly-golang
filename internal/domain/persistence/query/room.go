package query

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query/spec"
)

type IRoom interface {
	GetByID(ctx context.Context, id string) (model.Room, error)
	GetByName(ctx context.Context, name string) (model.Room, error)
	GetAllByUserID(ctx context.Context, userID string, params spec.SelectParams) (model.Rooms, error)
	SearchAllByName(ctx context.Context, keyword string, params spec.SelectParams) (model.Rooms, error)
}
