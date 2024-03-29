package query

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/model"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/query/selectParams"
)

type IComment interface {
	GetByID(ctx context.Context, id string) (model.Comment, error)
	GetAllByJokeID(ctx context.Context, jokeID string, params selectParams.SelectParams) (model.Comments, error)
}
