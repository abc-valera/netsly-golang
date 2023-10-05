package repository

import (
	"errors"

	"github.com/abc-valera/flugo-api-golang/internal/domain/codeerr"
)

type Repositories struct {
	UserRepo UserRepository
	// JokeRepo    JokeRepository
	// CommentRepo CommentRepository
	// LikeRepo    LikeRepository
}

func NewRepositories(userRepo UserRepository) (Repositories, error) {
	if userRepo == nil {
		return Repositories{}, codeerr.NewInternal("NewRepositories", errors.New("userRepo is nil"))
	}
	return Repositories{
		UserRepo: userRepo,
	}, nil
}
