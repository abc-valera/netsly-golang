package query

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/model"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/query/spec"
)

type IUser interface {
	GetByID(ctx context.Context, id string) (model.User, error)
	GetByUsername(ctx context.Context, username string) (model.User, error)
	GetByEmail(ctx context.Context, email string) (model.User, error)
	SearchAllByUsername(ctx context.Context, keyword string, params spec.SelectParams) (model.Users, error)
}
