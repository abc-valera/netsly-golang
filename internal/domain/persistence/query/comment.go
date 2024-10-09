package query

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/selector"
)

type IComment interface {
	GetByID(ctx context.Context, id string) (model.Comment, error)
	GetAllByJokeID(ctx context.Context, jokeID string, s selector.Selector) (model.Comments, error)
}
