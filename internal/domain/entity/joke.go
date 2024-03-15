package entity

import (
	"context"
	"time"

	"github.com/abc-valera/netsly-api-golang/internal/core/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
	"github.com/google/uuid"
)

type IJoke interface {
	Create(ctx context.Context, req JokeCreateRequest) (model.Joke, error)
	Update(ctx context.Context, jokeID string, req JokeUpdateRequest) (model.Joke, error)
	Delete(ctx context.Context, jokeID string) error
}

type joke struct {
	command command.IJoke
}

func NewJoke(
	command command.IJoke,
) IJoke {
	return joke{
		command: command,
	}
}

type JokeCreateRequest struct {
	Title       string `validate:"required,min=4,max=64"`
	Text        string `validate:"required,min=4,max=4096"`
	Explanation string `validate:"max=4096"`
	UserID      string `validate:"required,uuid"`
}

func (j joke) Create(ctx context.Context, req JokeCreateRequest) (model.Joke, error) {
	if err := global.Validator().Struct(req); err != nil {
		return model.Joke{}, err
	}

	return j.command.Create(ctx, model.Joke{
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

func (j joke) Update(ctx context.Context, jokeID string, req JokeUpdateRequest) (model.Joke, error) {
	if err := global.Validator().Struct(req); err != nil {
		return model.Joke{}, err
	}

	return j.command.Update(ctx, jokeID, command.JokeUpdate{
		Title:       req.Title,
		Text:        req.Text,
		Explanation: req.Explanation,
	})
}

func (j joke) Delete(ctx context.Context, jokeID string) error {
	if err := global.Validator().Var(jokeID, "uuid"); err != nil {
		return err
	}

	// Delete in data source
	return j.command.Delete(ctx, jokeID)
}
