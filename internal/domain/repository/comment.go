package repository

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/domain/entity"
	"github.com/abc-valera/flugo-api-golang/internal/domain/repository/common"
)

type CommentRepository interface {
	GetByID(ctx context.Context, id string) (*entity.Comment, error)
	GetByJokeID(ctx context.Context, jokeID string) (entity.Comments, error)
	Create(ctx context.Context, comment *entity.Comment) error
	Update(ctx context.Context, comment *entity.Comment) error
	Delete(ctx context.Context, id string) error

	common.Transactioneer
}
