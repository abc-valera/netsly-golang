package boilerSqliteCommand

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/boilerSqlite/boilerSqliteDto"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/boilerSqlite/boilerSqliteErrutil"
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

func (j joke) Create(ctx context.Context, req command.JokeCreateRequest) (model.Joke, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	joke := sqlboiler.Joke{
		ID:          req.Joke.ID,
		Title:       req.Joke.Title,
		Text:        req.Joke.Text,
		Explanation: req.Joke.Explanation,
		CreatedAt:   req.Joke.CreatedAt,
		UpdatedAt:   req.Joke.UpdatedAt,
		DeletedAt:   req.Joke.DeletedAt,

		UserID: req.UserID,
	}
	err := joke.Insert(ctx, j.executor, boil.Infer())
	return boilerSqliteDto.NewDomainJoke(&joke), boilerSqliteErrutil.HandleErr(err)
}

func (j joke) Update(ctx context.Context, id string, req command.JokeUpdateRequest) (model.Joke, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	joke, err := sqlboiler.FindJoke(ctx, j.executor, id)
	if err != nil {
		return model.Joke{}, boilerSqliteErrutil.HandleErr(err)
	}
	if req.Title != nil {
		joke.Title = *req.Title
	}
	if req.Text != nil {
		joke.Text = *req.Text
	}
	if req.Explanation != nil {
		joke.Explanation = *req.Explanation
	}
	_, err = joke.Update(ctx, j.executor, boil.Infer())
	return boilerSqliteDto.NewDomainJoke(joke), boilerSqliteErrutil.HandleErr(err)
}

func (j joke) Delete(ctx context.Context, id string) error {
	_, span := global.NewSpan(ctx)
	defer span.End()

	joke, err := sqlboiler.FindJoke(ctx, j.executor, id)
	if err != nil {
		return boilerSqliteErrutil.HandleErr(err)
	}
	_, err = joke.Delete(ctx, j.executor)
	return boilerSqliteErrutil.HandleErr(err)
}
