package impl

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/gen/ent"
	"github.com/abc-valera/flugo-api-golang/gen/ent/joke"
	"github.com/abc-valera/flugo-api-golang/internal/adapter/persistence/ent/dto"
	"github.com/abc-valera/flugo-api-golang/internal/adapter/persistence/ent/impl/common"
	"github.com/abc-valera/flugo-api-golang/internal/domain/entity"
	"github.com/abc-valera/flugo-api-golang/internal/domain/repository"
	"github.com/abc-valera/flugo-api-golang/internal/domain/repository/spec"
)

type jokeRepository struct {
	common.BaseRepository
}

func NewJokeRepository(client *ent.Client) repository.IJokeRepository {
	return &jokeRepository{
		BaseRepository: common.NewBaseRepository(client),
	}
}

func (r jokeRepository) Create(ctx context.Context, joke *entity.Joke) error {
	_, err := r.Client.Joke.
		Create().
		SetOwnerID(joke.UserID).
		SetID(joke.ID).
		SetUserID(joke.UserID).
		SetTitle(joke.Title).
		SetText(joke.Text).
		SetExplanation(joke.Explanation).
		SetCreatedAt(joke.CreatedAt).
		Save(ctx)
	return common.HandleErr(err)
}

func (r jokeRepository) GetByID(ctx context.Context, id string) (*entity.Joke, error) {
	entJoke, err := r.Client.Joke.
		Query().
		Where(joke.ID(id)).
		Only(ctx)
	return dto.FromEntJokeToJoke(entJoke), common.HandleErr(err)
}

func (r *jokeRepository) GetByUserID(ctx context.Context, userID string, spec spec.SelectParams) (entity.Jokes, error) {
	entJokes, err := r.specToQuery(spec).
		Where(joke.UserID(userID)).
		All(ctx)
	return dto.FromEntJokesToJokes(entJokes), common.HandleErr(err)
}

func (r *jokeRepository) Update(ctx context.Context, domainJoke *entity.Joke) error {
	_, err := r.Client.Joke.
		Update().
		Where(joke.ID(domainJoke.ID)).
		SetTitle(domainJoke.Title).
		SetText(domainJoke.Text).
		SetExplanation(domainJoke.Explanation).
		Save(ctx)
	return common.HandleErr(err)
}

func (r jokeRepository) Delete(ctx context.Context, jokeID string) error {
	_, err := r.Client.Joke.
		Delete().
		Where(joke.ID(jokeID)).
		Exec(ctx)
	return common.HandleErr(err)
}

func (r jokeRepository) specToQuery(spec spec.SelectParams) *ent.JokeQuery {
	// Map orderBy
	var orderBy string
	if spec.OrderBy == "" || spec.OrderBy == "created_at" {
		orderBy = joke.FieldCreatedAt
	}
	if spec.OrderBy == "title" {
		orderBy = joke.FieldTitle
	}

	// Map order
	entSpec := r.Client.Joke.Query()
	if spec.Order == "asc" {
		entSpec.Order(ent.Asc(orderBy))
	}
	if spec.Order == "desc" {
		entSpec.Order(ent.Desc(orderBy))
	}

	// Map limit offset
	return entSpec.
		Limit(int(spec.Limit)).
		Offset(int(spec.Offset))
}
