package boilerQuery

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/model"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/query"
	selector1 "github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/query/selector"
	"github.com/abc-valera/netsly-api-golang/pkg/infrastructure/persistence/boiler/boilerDto"
	"github.com/abc-valera/netsly-api-golang/pkg/infrastructure/persistence/boiler/boilerQuery/selector"
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
	return boilerDto.ToDomainJokeWithErrHandle(sqlboiler.FindJoke(ctx, j.executor, id))
}

func (j joke) GetAllByUserID(ctx context.Context, userID string, params selector1.Selector) (model.Jokes, error) {
	mods := selector.ToBoilerSelectorPipe(
		params,
		sqlboiler.JokeWhere.UserID.EQ(userID),
	)
	return boilerDto.ToDomainJokesWithErrHandle(sqlboiler.Jokes(mods...).All(ctx, j.executor))
}

func (j joke) SearchByTitle(ctx context.Context, keyword string, params selector1.Selector) (model.Jokes, error) {
	mods := selector.ToBoilerSelectorPipe(
		params,
		sqlboiler.JokeWhere.Title.LIKE("%"+keyword+"%"),
	)
	return boilerDto.ToDomainJokesWithErrHandle(sqlboiler.Jokes(mods...).All(ctx, j.executor))
}
