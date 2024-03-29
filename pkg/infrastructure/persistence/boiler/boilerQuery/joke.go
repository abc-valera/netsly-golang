package boilerQuery

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/model"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/query"
	selectParams1 "github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/query/selectParams"
	"github.com/abc-valera/netsly-api-golang/pkg/infrastructure/persistence/boiler/boilerQuery/selectParams"
	"github.com/abc-valera/netsly-api-golang/pkg/infrastructure/persistence/boiler/dto"
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
	return dto.ToDomainJokeWithErrHandle(sqlboiler.FindJoke(ctx, j.executor, id))
}

func (j joke) GetAllByUserID(ctx context.Context, userID string, params selectParams1.SelectParams) (model.Jokes, error) {
	mods := selectParams.ToBoilerSelectParamsPipe(
		params,
		sqlboiler.JokeWhere.UserID.EQ(userID),
	)
	return dto.ToDomainJokesWithErrHandle(sqlboiler.Jokes(mods...).All(ctx, j.executor))
}

func (j joke) SearchByTitle(ctx context.Context, keyword string, params selectParams1.SelectParams) (model.Jokes, error) {
	mods := selectParams.ToBoilerSelectParamsPipe(
		params,
		sqlboiler.JokeWhere.Title.LIKE("%"+keyword+"%"),
	)
	return dto.ToDomainJokesWithErrHandle(sqlboiler.Jokes(mods...).All(ctx, j.executor))
}
