package application

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/domain/entity"
	"github.com/abc-valera/flugo-api-golang/internal/domain/repository"
)

type LikeUseCase struct {
	likeRepo repository.ILikeRepository
}

func NewLikeUseCase(likeRepo repository.ILikeRepository) LikeUseCase {
	return LikeUseCase{
		likeRepo: likeRepo,
	}
}

type CreateLikeRequest struct {
	UserID string
	JokeID string
}

func (uc LikeUseCase) CreateLike(ctx context.Context, req CreateLikeRequest) error {
	like, err := entity.NewLike(req.UserID, req.JokeID)
	if err != nil {
		return err
	}
	return uc.likeRepo.Create(ctx, like)
}
