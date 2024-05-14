package boilerCommand

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/model"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/pkg/infrastructure/persistence/boiler/boilerDto"
	"github.com/abc-valera/netsly-api-golang/pkg/infrastructure/persistence/boiler/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type joke struct {
	executor boil.ContextExecutor
}

func NewJoke(executor boil.ContextExecutor) command.IJoke {
	return &joke{
		executor: executor,
	}
}

func (j joke) Create(ctx context.Context, req model.Joke) (model.Joke, error) {
	joke := sqlboiler.Joke{
		ID:          req.ID,
		Title:       req.Title,
		Text:        req.Text,
		Explanation: req.Explanation,
		CreatedAt:   req.CreatedAt,
		UserID:      req.UserID,
	}
	err := joke.Insert(ctx, j.executor, boil.Infer())
	return boilerDto.ToDomainJokeWithErrHandle(&joke, err)
}

func (j joke) Update(ctx context.Context, id string, req command.JokeUpdate) (model.Joke, error) {
	joke, err := sqlboiler.FindJoke(ctx, j.executor, id)
	if err != nil {
		return model.Joke{}, errors.HandleErr(err)
	}
	if req.Title.IsPresent() {
		joke.Title = req.Title.Value()
	}
	if req.Text.IsPresent() {
		joke.Text = req.Text.Value()
	}
	if req.Explanation.IsPresent() {
		joke.Explanation = req.Explanation.Value()
	}
	_, err = joke.Update(ctx, j.executor, boil.Infer())
	return boilerDto.ToDomainJokeWithErrHandle(joke, err)
}

func (j joke) Delete(ctx context.Context, id string) error {
	joke, err := sqlboiler.FindJoke(ctx, j.executor, id)
	if err != nil {
		return errors.HandleErr(err)
	}
	_, err = joke.Delete(ctx, j.executor)
	return errors.HandleErr(err)
}
