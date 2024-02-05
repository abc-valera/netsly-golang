package entquery

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/abc-valera/netsly-api-golang/gen/ent"
	"github.com/abc-valera/netsly-api-golang/gen/ent/joke"
	"github.com/abc-valera/netsly-api-golang/internal/adapter/persistence/ent-impl/dto"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query/spec"
)

type jokeQuery struct {
	*ent.Client
}

func NewJokeQuery(client *ent.Client) query.IJoke {
	return &jokeQuery{
		Client: client,
	}
}

func (jq *jokeQuery) GetByID(ctx context.Context, id string) (model.Joke, error) {
	return dto.FromEntJokeToJokeWithErrHandle(jq.Joke.Get(ctx, id))
}

func (jq *jokeQuery) GetAllByUserID(ctx context.Context, userID string, params spec.SelectParams) (model.Jokes, error) {
	query := jq.Joke.
		Query().
		Where(joke.UserID(userID))

	if params.Order == "asc" {
		query = query.Order(ent.Asc("created_at"))
	} else {
		query = query.Order(ent.Desc("created_at"))
	}

	query.Limit(params.Limit)
	query.Offset(params.Offset)

	return dto.FromEntJokesToJokesWithErrHandle(query.All(ctx))
}

func (jq *jokeQuery) SearchByTitle(ctx context.Context, keyword string, params spec.SelectParams) (model.Jokes, error) {
	query := jq.Joke.
		Query().
		Where(func(s *sql.Selector) { s.Where(sql.Like("title", "%"+keyword+"%")) })

	if params.Order == "asc" {
		query = query.Order(ent.Asc("created_at"))
	} else {
		query = query.Order(ent.Desc("created_at"))
	}

	query.Limit(params.Limit)
	query.Offset(params.Offset)

	return dto.FromEntJokesToJokesWithErrHandle(query.All(ctx))
}
