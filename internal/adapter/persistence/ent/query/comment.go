package query

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/gen/ent"
	"github.com/abc-valera/flugo-api-golang/gen/ent/comment"
	"github.com/abc-valera/flugo-api-golang/internal/adapter/persistence/ent/dto"
	errhandler "github.com/abc-valera/flugo-api-golang/internal/adapter/persistence/ent/err-handler"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/model"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/query"
)

type commentQuery struct {
	*ent.Client
}

func NewCommentQuery(client *ent.Client) query.ICommentQuery {
	return &commentQuery{
		Client: client,
	}
}

func (cq *commentQuery) GetAll(ctx context.Context, params query.CommentSelectParams) (model.Comments, error) {
	query := cq.Comment.Query()

	// Order
	orderByField := "created_at"

	if params.Order == "asc" {
		query.Order(ent.Asc(orderByField))
	} else {
		query.Order(ent.Desc(orderByField))
	}

	// Limit and Offset
	query.Limit(params.Limit)
	query.Offset(params.Offset)

	entUsers, err := query.All(ctx)
	return dto.FromEntCommentsToComments(entUsers), errhandler.HandleErr(err)
}

func (cq *commentQuery) GetOne(ctx context.Context, fiedls query.CommentGetFields) (*model.Comment, error) {
	query := cq.Comment.Query()

	// Where
	if fiedls.ID != "" {
		query.Where(comment.ID(fiedls.ID))
	}

	entUser, err := query.Only(ctx)
	return dto.FromEntCommentToComment(entUser), errhandler.HandleErr(err)
}
