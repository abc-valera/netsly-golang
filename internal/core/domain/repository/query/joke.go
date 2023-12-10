package query

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/model"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/query/spec"
)

type IJokeQuery interface {
	GetAll(ctx context.Context, params JokeSelectParams) (model.Jokes, error)
	GetOne(ctx context.Context, fields JokeGetFields) (*model.Joke, error)
}

type JokeSearchByFields struct {
	Title string
	Text  string
}

type JokeOrderByFields struct {
	Title     bool
	CreatedAt bool
}

type JokeSelectParams struct {
	SearchBy JokeSearchByFields
	OrderBy  JokeOrderByFields
	spec.SelectParams
}

func NewJokeSelectParams(
	searchBy JokeSearchByFields,
	orderBy JokeOrderByFields,
	order string,
	limit int,
	offset int,
) (JokeSelectParams, error) {
	commonSelectParams, err := spec.NewSelectParams(order, limit, offset)
	if err != nil {
		return JokeSelectParams{}, err
	}
	return JokeSelectParams{
		SearchBy:     searchBy,
		OrderBy:      orderBy,
		SelectParams: commonSelectParams,
	}, nil
}

type JokeGetFields struct {
	ID     string
	UserID string
	Title  string
}
