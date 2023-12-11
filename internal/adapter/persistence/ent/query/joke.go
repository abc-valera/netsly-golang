package query

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/abc-valera/flugo-api-golang/gen/ent"
	"github.com/abc-valera/flugo-api-golang/gen/ent/joke"
	"github.com/abc-valera/flugo-api-golang/internal/adapter/persistence/ent/dto"
	errhandler "github.com/abc-valera/flugo-api-golang/internal/adapter/persistence/ent/err-handler"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/model"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/query"
)

type jokeQuery struct {
	*ent.Client
}

func NewJokeQuery(client *ent.Client) query.IJokeQuery {
	return &jokeQuery{
		Client: client,
	}
}

func (jq *jokeQuery) GetAll(ctx context.Context, params query.JokeSelectParams) (model.Jokes, error) {
	query := jq.Joke.Query()

	// Where
	if params.SearchBy.Title != "" {
		query.Where(func(s *sql.Selector) {
			s.Where(sql.Like("title", "%"+params.SearchBy.Title+"%"))
		})
	}
	if params.SearchBy.Text != "" {
		query.Where(func(s *sql.Selector) {
			s.Where(sql.Like("text", "%"+params.SearchBy.Text+"%"))
		})
	}

	// Order
	orderByField := "created_at"
	if params.OrderBy.Title == true {
		orderByField = "username"
	}

	if params.Order == "asc" {
		query.Order(ent.Asc(orderByField))
	} else {
		query.Order(ent.Desc(orderByField))
	}

	// Limit and Offset
	query.Limit(params.Limit)
	query.Offset(params.Offset)

	entJokes, err := query.All(ctx)
	return dto.FromEntJokesToJokes(entJokes), errhandler.HandleErr(err)
}

func (jq *jokeQuery) GetOne(ctx context.Context, fields query.JokeGetFields) (*model.Joke, error) {
	query := jq.Joke.Query()

	// Where
	if fields.ID != "" {
		query.Where(joke.ID(fields.ID))
	}
	if fields.Title != "" {
		query.Where(joke.Title(fields.Title))
	}
	if fields.UserID != "" {
		query.Where(joke.UserID(fields.UserID))
	}

	entJoke, err := query.Only(ctx)
	return dto.FromEntJokeToJoke(entJoke), errhandler.HandleErr(err)
}
