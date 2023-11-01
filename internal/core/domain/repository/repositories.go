package repository

import (
	"errors"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/codeerr"
)

type Repositories struct {
	UserRepo    IUserRepository
	JokeRepo    IJokeRepository
	CommentRepo ICommentRepository
	LikeRepo    ILikeRepository
}

func NewRepositories(
	userRepo IUserRepository,
	jokeRepo IJokeRepository,
	commentRepo ICommentRepository,
	likeRepo ILikeRepository,
) (Repositories, error) {
	if userRepo == nil {
		return Repositories{}, codeerr.NewInternal("NewRepositories", errors.New("userRepo is nil"))
	}
	if jokeRepo == nil {
		return Repositories{}, codeerr.NewInternal("NewRepositories", errors.New("jokeRepo is nil"))
	}
	if commentRepo == nil {
		return Repositories{}, codeerr.NewInternal("NewRepositories", errors.New("commentRepo is nil"))
	}
	if likeRepo == nil {
		return Repositories{}, codeerr.NewInternal("NewRepositories", errors.New("likeRepo is nil"))
	}
	return Repositories{
		UserRepo:    userRepo,
		JokeRepo:    jokeRepo,
		CommentRepo: commentRepo,
		LikeRepo:    likeRepo,
	}, nil
}
