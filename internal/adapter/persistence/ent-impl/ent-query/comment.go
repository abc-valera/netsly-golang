package entquery

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/gen/ent"
	"github.com/abc-valera/flugo-api-golang/gen/ent/comment"
	"github.com/abc-valera/flugo-api-golang/internal/adapter/persistence/ent-impl/dto"
	"github.com/abc-valera/flugo-api-golang/internal/core/model"
	"github.com/abc-valera/flugo-api-golang/internal/core/persistence/query"
	"github.com/abc-valera/flugo-api-golang/internal/core/persistence/query/spec"
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
		Where(comment.JokeID(jokeID))

	if params.Order == "asc" {
		query = query.Order(ent.Asc("created_at"))
	} else {
		query = query.Order(ent.Desc("created_at"))
	}

	query.Limit(params.Limit)
	query.Offset(params.Offset)

	return dto.FromEntCommentsToCommentsWithErrHandle(query.All(ctx))
}
