package entity

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/internal/domain/entity/common"
	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
)

type Joke struct {
	command command.IJoke
}

func NewJoke(
	command command.IJoke,
) Joke {
	return Joke{
		command: command,
	}
}

type JokeCreateRequest struct {
	Title       string `validate:"required,min=4,max=64"`
	Text        string `validate:"required,min=4,max=4096"`
	Explanation string `validate:"max=4096"`
	UserID      string `validate:"required,uuid"`
}

func (j Joke) Create(ctx context.Context, req JokeCreateRequest) (model.Joke, error) {
	if err := global.Validator().Struct(req); err != nil {
		return model.Joke{}, err
	}

	baseModel := common.NewBaseEntity()

	return j.command.Create(ctx, model.Joke{
		BaseEntity:  baseModel,
		Title:       req.Title,
		Text:        req.Text,
		Explanation: req.Explanation,
		UserID:      req.UserID,
	})
}

type JokeUpdateRequest struct {
	Title       *string `validate:"min=4,max=64"`
	Text        *string `validate:"min=4,max=4096"`
	Explanation *string `validate:"max=4096"`
}

func (j Joke) Update(ctx context.Context, jokeID string, req JokeUpdateRequest) (model.Joke, error) {
	if err := global.Validator().Struct(req); err != nil {
		return model.Joke{}, err
	}

	return j.command.Update(ctx, jokeID, command.JokeUpdate{
		Title:       req.Title,
		Text:        req.Text,
		Explanation: req.Explanation,
	})
}

func (j Joke) Delete(ctx context.Context, jokeID string) error {
	if err := global.Validator().Var(jokeID, "uuid"); err != nil {
		return err
	}

	// Delete in data source
	return j.command.Delete(ctx, jokeID)
}
