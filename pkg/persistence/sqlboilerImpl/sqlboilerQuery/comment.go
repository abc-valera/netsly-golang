package sqlboilerQuery

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/model"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/query"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/query/spec"
	"github.com/abc-valera/netsly-api-golang/pkg/persistence/sqlboilerImpl/dto"
	"github.com/abc-valera/netsly-api-golang/pkg/persistence/sqlboilerImpl/sqlboilerQuery/common"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type comment struct {
	executor boil.ContextExecutor
}

func newComment(executor boil.ContextExecutor) query.IComment {
	return &comment{
		executor: executor,
	}
}

func (c comment) GetByID(ctx context.Context, id string) (model.Comment, error) {
	return dto.ToDomainCommentWithErrHandle(sqlboiler.FindComment(ctx, c.executor, id))
}

func (c comment) GetAllByJokeID(ctx context.Context, jokeID string, params spec.SelectParams) (model.Comments, error) {
	mods := common.ToBoilerSelectParamsPipe(
		params,
		sqlboiler.CommentWhere.JokeID.EQ(jokeID),
	)
	return dto.ToDomainCommentsWithErrHandle(sqlboiler.Comments(mods...).All(ctx, c.executor))
}
