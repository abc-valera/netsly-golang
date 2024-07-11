package boilerQuery

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query"
	selector1 "github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query/selector"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/boiler/boilerDto"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/boiler/boilerQuery/selector"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type joke struct {
	executor boil.ContextExecutor
}

func NewJoke(executor boil.ContextExecutor) query.IJoke {
	return &joke{
		executor: executor,
	}
}

func (j joke) GetByID(ctx context.Context, id string) (model.Joke, error) {
	return boilerDto.NewDomainJokeWithErrHandle(sqlboiler.FindJoke(ctx, j.executor, id))
}

func (j joke) GetAllByUserID(ctx context.Context, userID string, params selector1.Selector) (model.Jokes, error) {
	mods := selector.ToBoilerSelectorPipe(
		params,
		sqlboiler.JokeWhere.UserID.EQ(userID),
	)
	return boilerDto.NewDomainJokesWithErrHandle(sqlboiler.Jokes(mods...).All(ctx, j.executor))
}

func (j joke) SearchAllByTitle(ctx context.Context, keyword string, params selector1.Selector) (model.Jokes, error) {
	mods := selector.ToBoilerSelectorPipe(
		params,
		sqlboiler.JokeWhere.Title.LIKE("%"+keyword+"%"),
	)
	return boilerDto.NewDomainJokesWithErrHandle(sqlboiler.Jokes(mods...).All(ctx, j.executor))
}
