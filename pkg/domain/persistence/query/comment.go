package query

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/pkg/domain/model"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/query/selector"
)

type IComment interface {
	GetByID(ctx context.Context, id string) (model.Comment, error)
	GetAllByJokeID(ctx context.Context, jokeID string, params selector.Selector) (model.Comments, error)
}
