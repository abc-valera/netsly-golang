package boilerQuery

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query"
	selector1 "github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query/selector"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/boiler/boilerDto"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/boiler/boilerQuery/selector"
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
	return boilerDto.ToDomainCommentWithErrHandle(sqlboiler.FindComment(ctx, c.executor, id))
}

func (c comment) GetAllByJokeID(ctx context.Context, jokeID string, params selector1.Selector) (model.Comments, error) {
	mods := selector.ToBoilerSelectorPipe(
		params,
		sqlboiler.CommentWhere.JokeID.EQ(jokeID),
	)
	return boilerDto.ToDomainCommentsWithErrHandle(sqlboiler.Comments(mods...).All(ctx, c.executor))
}
