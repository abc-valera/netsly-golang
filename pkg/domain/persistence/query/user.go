package query

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/model"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/query/selector"
)

type IUser interface {
	GetByID(ctx context.Context, id string) (model.User, error)
	GetByUsername(ctx context.Context, username string) (model.User, error)
	GetByEmail(ctx context.Context, email string) (model.User, error)
	SearchAllByUsername(ctx context.Context, keyword string, params selector.Selector) (model.Users, error)
}
