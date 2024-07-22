package boilerSqliteQuery

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query"
	domainSelector "github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query/selector"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/boilerSqlite/boilerSqliteDto"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/boilerSqlite/boilerSqliteErrutil"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/boilerSqlite/boilerSqliteQuery/selector"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type comment struct {
	executor boil.ContextExecutor
}

func NewComment(executor boil.ContextExecutor) query.IComment {
	return &comment{
		executor: executor,
	}
}

func (c comment) GetByID(ctx context.Context, id string) (model.Comment, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	comment, err := sqlboiler.FindComment(ctx, c.executor, id)
	return boilerSqliteDto.NewDomainComment(comment), boilerSqliteErrutil.HandleErr(err)
}

func (c comment) GetAllByJokeID(ctx context.Context, jokeID string, params domainSelector.Selector) (model.Comments, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	mods := selector.ToBoilerSelectorPipe(
		params,
		sqlboiler.CommentWhere.JokeID.EQ(jokeID),
	)
	comments, err := sqlboiler.Comments(mods...).All(ctx, c.executor)
	return boilerSqliteDto.NewDomainComments(comments), boilerSqliteErrutil.HandleErr(err)
}
