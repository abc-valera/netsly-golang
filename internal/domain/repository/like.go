package repository

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/domain/entity"
)

type LikeRepository interface {
	CountByJokeID(ctx context.Context, jokeID string) (int, error)
	Create(ctx context.Context, like *entity.Like) error
	Update(ctx context.Context, like *entity.Like) error
	Delete(ctx context.Context, userID, jokeID string) error

	Transactioneer
}
