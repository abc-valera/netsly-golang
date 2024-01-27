package query

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/core/model"
	"github.com/abc-valera/flugo-api-golang/internal/core/persistence/query/spec"
)

type IComment interface {
	GetByID(ctx context.Context, id string) (model.Comment, error)
	GetAllByJokeID(ctx context.Context, jokeID string, params spec.SelectParams) (model.Comments, error)
}
