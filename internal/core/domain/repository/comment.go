package repository

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/codeerr"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/entity"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/common"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/spec"
)

var (
	ErrCommentNotFound = codeerr.NewMsgErr(codeerr.CodeNotFound, "Comment not found")
)

type ICommentRepository interface {
	GetByID(ctx context.Context, id string) (*entity.Comment, error)
	GetByJokeID(ctx context.Context, jokeID string, spec spec.SelectParams) (entity.Comments, error)
	Create(ctx context.Context, comment *entity.Comment) error
	Update(ctx context.Context, comment *entity.Comment) error
	Delete(ctx context.Context, commentID string) error

	common.Transactioneer
}
