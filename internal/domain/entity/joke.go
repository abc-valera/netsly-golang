package entity

import (
	"context"
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/global"
	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query"
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

	return e.C().Joke.Create(ctx, model.Joke{
		ID:          uuid.New().String(),
		Title:       req.Title,
		Text:        req.Text,
		Explanation: req.Explanation,
		CreatedAt:   time.Now(),
		UserID:      req.UserID,
	})
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

	return e.C().Joke.Update(
		ctx,
		model.Joke{ID: jokeID},
		command.JokeUpdateRequest{
			UpdatedAt: time.Now(),

			Title:       req.Title,
			Text:        req.Text,
			Explanation: req.Explanation,
		})
}

func (e joke) Delete(ctx context.Context, jokeID string) error {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	if err := global.Validate().Var(jokeID, "uuid"); err != nil {
		return err
	}

	return e.C().Joke.Delete(ctx, model.Joke{ID: jokeID})
}
