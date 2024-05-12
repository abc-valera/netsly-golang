package boilerQuery

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/model"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/query"
	selector1 "github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/query/selector"
	"github.com/abc-valera/netsly-api-golang/pkg/infrastructure/persistence/boiler/boilerQuery/selector"
	"github.com/abc-valera/netsly-api-golang/pkg/infrastructure/persistence/boiler/dto"
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
	return dto.ToDomainCommentWithErrHandle(sqlboiler.FindComment(ctx, c.executor, id))
}

func (c comment) GetAllByJokeID(ctx context.Context, jokeID string, params selector1.Selector) (model.Comments, error) {
	mods := selector.ToBoilerSelectorPipe(
		params,
		sqlboiler.CommentWhere.JokeID.EQ(jokeID),
	)
	return dto.ToDomainCommentsWithErrHandle(sqlboiler.Comments(mods...).All(ctx, c.executor))
}
