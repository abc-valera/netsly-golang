package entity

import (
	"context"
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/global"
	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/queryUtil/filter"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"
)

type IJoke interface {
	Create(ctx context.Context, req JokeCreateRequest) (model.Joke, error)
	Update(ctx context.Context, jokeID string, req JokeUpdateRequest) (model.Joke, error)
	Delete(ctx context.Context, jokeID string) error

	query.IJoke
}

type joke struct {
	IDependency

	query.IJoke
}

func newJoke(dep IDependency) IJoke {
	return joke{
		IDependency: dep,

		IJoke: dep.Q().Joke,
	}
}

type JokeCreateRequest struct {
	Title       string `validate:"required,min=4,max=64"`
	Text        string `validate:"required,min=4,max=4096"`
	Explanation string `validate:"max=4096"`

	UserID string `validate:"required,uuid"`
}

func (e joke) Create(ctx context.Context, req JokeCreateRequest) (model.Joke, error) {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	if err := global.Validate().Struct(req); err != nil {
		return model.Joke{}, err
	}

	joke := model.Joke{
		ID:          uuid.New().String(),
		Title:       req.Title,
		Text:        req.Text,
		Explanation: req.Explanation,
		CreatedAt:   time.Now().Truncate(time.Millisecond).Local(),
		UserID:      req.UserID,
	}

	if err := e.C().Joke.Create(ctx, joke); err != nil {
		return model.Joke{}, err
	}

	return joke, nil
}

type JokeUpdateRequest struct {
	Title       *string `validate:"min=4,max=64"`
	Text        *string `validate:"min=4,max=4096"`
	Explanation *string `validate:"max=4096"`
}

func (e joke) Update(ctx context.Context, jokeID string, req JokeUpdateRequest) (model.Joke, error) {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	if err := global.Validate().Struct(req); err != nil {
		return model.Joke{}, err
	}

	joke, err := e.Q().Joke.GetOne(ctx, filter.By(model.Joke{ID: jokeID}))
	if err != nil {
		return model.Joke{}, err
	}

	joke.UpdatedAt = time.Now().Truncate(time.Millisecond).Local()

	if req.Title != nil {
		joke.Title = *req.Title
	}

	if req.Text != nil {
		joke.Text = *req.Text
	}

	if req.Explanation != nil {
		joke.Explanation = *req.Explanation
	}

	if err := e.C().Joke.Update(ctx, joke); err != nil {
		return model.Joke{}, err
	}

	return joke, nil
}

func (e joke) Delete(ctx context.Context, jokeID string) error {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	return e.C().Joke.Delete(ctx, model.Joke{ID: jokeID})
}
