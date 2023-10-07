package repository

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/domain/entity"
	"github.com/abc-valera/flugo-api-golang/internal/domain/repository/common"
)

type LikeRepository interface {
	CountByJokeID(ctx context.Context, jokeID string) (int, error)
	Create(ctx context.Context, like *entity.Like) error
	Update(ctx context.Context, like *entity.Like) error
	Delete(ctx context.Context, userID, jokeID string) error

	common.Transactioneer
}
