package query

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/model"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/query/spec"
)

type ICommentQuery interface {
	GetByID(ctx context.Context, id string) (*model.Comment, error)
	GetAllByJokeID(ctx context.Context, jokeID string, params spec.SelectParams) (model.Comments, error)
}
