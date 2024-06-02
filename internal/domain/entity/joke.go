package entity

import (
	"context"
	"time"

	"github.com/abc-valera/netsly-api-golang/internal/core/optional"
	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query"
	"github.com/google/uuid"
)

type IJoke interface {
	Create(ctx context.Context, req JokeCreateRequest) (model.Joke, error)
	Update(ctx context.Context, jokeID string, req JokeUpdateRequest) (model.Joke, error)
	Delete(ctx context.Context, jokeID string) error

	query.IJoke
}

type joke struct {
	command command.IJoke
	query.IJoke
}

func NewJoke(
	command command.IJoke,
	query query.IJoke,
) IJoke {
	return joke{
		command: command,
		IJoke:   query,
	}
}

type JokeCreateRequest struct {
	Title       string `validate:"required,min=4,max=64"`
	Text        string `validate:"required,min=4,max=4096"`
	Explanation string `validate:"max=4096"`

	UserID string `validate:"required,uuid"`
}

func (j joke) Create(ctx context.Context, req JokeCreateRequest) (model.Joke, error) {
	if err := global.Validate().Struct(req); err != nil {
		return model.Joke{}, err
	}

	return j.command.Create(ctx, req.UserID, model.Joke{
		ID:          uuid.New().String(),
		Title:       req.Title,
		Text:        req.Text,
		Explanation: req.Explanation,
		CreatedAt:   time.Now(),
	})
}

type JokeUpdateRequest struct {
	Title       optional.Optional[string] `validate:"min=4,max=64"`
	Text        optional.Optional[string] `validate:"min=4,max=4096"`
	Explanation optional.Optional[string] `validate:"max=4096"`
}

func (j joke) Update(ctx context.Context, jokeID string, req JokeUpdateRequest) (model.Joke, error) {
	if err := global.Validate().Struct(req); err != nil {
		return model.Joke{}, err
	}

	return j.command.Update(ctx, jokeID, command.JokeUpdate{
		Title:       req.Title,
		Text:        req.Text,
		Explanation: req.Explanation,
	})
}

func (j joke) Delete(ctx context.Context, jokeID string) error {
	if err := global.Validate().Var(jokeID, "uuid"); err != nil {
		return err
	}

	return j.command.Delete(ctx, jokeID)
}
