package query

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/codeerr"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/model"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/query/spec"
)

var (
	ErrCommentNotFound = codeerr.NewMessageErr(codeerr.CodeNotFound, "Comment not found")
)

type ICommentQuery interface {
	GetAll(ctx context.Context, params CommentSelectParams) (model.Comments, error)
	GetOne(ctx context.Context, fiedls CommentGetFields) (*model.Comment, error)
}

type CommentSearchFields struct {
	UserID string
	JokeID string
}

type CommentOrderByFields struct {
	CreatedAt bool
}

type CommentSelectParams struct {
	SearchBy CommentSearchFields
	OrderBy  CommentOrderByFields
	spec.SelectParams
}

func NewCommentSelectParams(
	searchBy CommentSearchFields,
	orderBy CommentOrderByFields,
	order string,
	limit int,
	offset int,
) (CommentSelectParams, error) {
	commonSelectParams, err := spec.NewSelectParams(order, limit, offset)
	if err != nil {
		return CommentSelectParams{}, err
	}
	return CommentSelectParams{
		SearchBy:     searchBy,
		OrderBy:      orderBy,
		SelectParams: commonSelectParams,
	}, nil
}

type CommentGetFields struct {
	ID string
}
