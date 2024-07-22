package boilerSqliteQuery

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query"
	selector1 "github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query/selector"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/boilerSqlite/boilerSqliteDto"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/boilerSqlite/boilerSqliteQuery/selector"
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
	_, span := global.NewSpan(ctx)
	defer span.End()

	joke, err := sqlboiler.FindJoke(ctx, j.executor, id)
	return boilerSqliteDto.NewDomainJoke(joke), err
}

func (j joke) GetAllByUserID(ctx context.Context, userID string, params selector1.Selector) (model.Jokes, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	mods := selector.ToBoilerSelectorPipe(
		params,
		sqlboiler.JokeWhere.UserID.EQ(userID),
	)
	jokes, err := sqlboiler.Jokes(mods...).All(ctx, j.executor)
	return boilerSqliteDto.NewDomainJokes(jokes), err
}

func (j joke) SearchAllByTitle(ctx context.Context, keyword string, params selector1.Selector) (model.Jokes, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	mods := selector.ToBoilerSelectorPipe(
		params,
		sqlboiler.JokeWhere.Title.LIKE("%"+keyword+"%"),
	)
	jokes, err := sqlboiler.Jokes(mods...).All(ctx, j.executor)
	return boilerSqliteDto.NewDomainJokes(jokes), err
}
