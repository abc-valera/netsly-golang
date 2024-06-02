package boilerCommand

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/boiler/boilerDto"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/boiler/errutil"
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

func (j joke) Create(ctx context.Context, userID string, req model.Joke) (model.Joke, error) {
	joke := sqlboiler.Joke{
		ID:          req.ID,
		Title:       req.Title,
		Text:        req.Text,
		Explanation: req.Explanation,
		CreatedAt:   req.CreatedAt,

		UserID: userID,
	}
	err := joke.Insert(ctx, j.executor, boil.Infer())
	return boilerDto.NewDomainJokeWithErrHandle(&joke, err)
}

func (j joke) Update(ctx context.Context, id string, req command.JokeUpdate) (model.Joke, error) {
	joke, err := sqlboiler.FindJoke(ctx, j.executor, id)
	if err != nil {
		return model.Joke{}, errutil.HandleErr(err)
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
	return boilerDto.NewDomainJokeWithErrHandle(joke, err)
}

func (j joke) Delete(ctx context.Context, id string) error {
	joke, err := sqlboiler.FindJoke(ctx, j.executor, id)
	if err != nil {
		return errutil.HandleErr(err)
	}
	_, err = joke.Delete(ctx, j.executor)
	return errutil.HandleErr(err)
}
