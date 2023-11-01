package repository

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/entity"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/common"
)

type ILikeRepository interface {
	CountByJokeID(ctx context.Context, jokeID string) (int, error)
	Create(ctx context.Context, like *entity.Like) error
	Delete(ctx context.Context, userID, jokeID string) error

	common.Transactioneer
}
