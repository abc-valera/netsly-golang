package sqlboilerquery

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query/spec"
	"github.com/abc-valera/netsly-api-golang/internal/persistence/sqlboiler-impl/dto"
	"github.com/abc-valera/netsly-api-golang/internal/persistence/sqlboiler-impl/sqlboiler-query/common"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type joke struct {
	executor boil.ContextExecutor
}

func newJoke(executor boil.ContextExecutor) query.IJoke {
	return &joke{
		executor: executor,
	}
}

func (j joke) GetByID(ctx context.Context, id string) (model.Joke, error) {
	return dto.ToDomainJokeWithErrHandle(sqlboiler.FindJoke(ctx, j.executor, id))
}

func (j joke) GetAllByUserID(ctx context.Context, userID string, params spec.SelectParams) (model.Jokes, error) {
	mods := common.ToBoilerSelectParamsPipe(
		params,
		sqlboiler.JokeWhere.UserID.EQ(userID),
	)
	return dto.ToDomainJokesWithErrHandle(sqlboiler.Jokes(mods...).All(ctx, j.executor))
}

func (j joke) SearchByTitle(ctx context.Context, keyword string, params spec.SelectParams) (model.Jokes, error) {
	mods := common.ToBoilerSelectParamsPipe(
		params,
		sqlboiler.JokeWhere.Title.LIKE("%"+keyword+"%"),
	)
	return dto.ToDomainJokesWithErrHandle(sqlboiler.Jokes(mods...).All(ctx, j.executor))
}
