package entquery

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/ent"
	"github.com/abc-valera/netsly-api-golang/gen/ent/comment"
	"github.com/abc-valera/netsly-api-golang/gen/ent/joke"
	"github.com/abc-valera/netsly-api-golang/internal/adapter/persistence/ent-impl/dto"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query/spec"
)

type commentQuery struct {
	*ent.Client
}

func NewCommentQuery(client *ent.Client) query.IComment {
	return &commentQuery{
		Client: client,
	}
}

func (cq *commentQuery) GetByID(ctx context.Context, id string) (model.Comment, error) {
	return dto.FromEntCommentToCommentWithErrHandle(cq.Comment.Get(ctx, id))
}

func (cq *commentQuery) GetAllByJokeID(ctx context.Context, jokeID string, params spec.SelectParams) (model.Comments, error) {
	query := cq.Comment.
		Query().
		Where(comment.HasJokeWith(joke.ID(jokeID)))

	if params.Order() == "asc" {
		query = query.Order(ent.Asc("created_at"))
	} else {
		query = query.Order(ent.Desc("created_at"))
	}

	query.Limit(params.Limit())
	query.Offset(params.Offset())

	return dto.FromEntCommentsToCommentsWithErrHandle(query.All(ctx))
}
