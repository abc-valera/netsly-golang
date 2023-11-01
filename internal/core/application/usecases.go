package application

import (
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/service"
)

type UseCases struct {
	SignUseCase    SignUseCase
	UserUseCase    UserUseCase
	JokeUseCase    JokeUseCase
	CommentUseCase CommentUseCase
	LikeUseCase    LikeUseCase
}

func NewUseCases(
	repos repository.Repositories,
	services service.Services,
) (UseCases, error) {
	return UseCases{
		SignUseCase: NewSignUseCase(
			repos.UserRepo,
			services.PasswordMaker,
			services.TokenMaker,
			services.MessageBroker,
		),
		UserUseCase: NewUserUseCase(
			repos.UserRepo,
		),
		JokeUseCase: NewJokeUseCase(
			repos.JokeRepo,
		),
		CommentUseCase: NewCommentUseCase(
			repos.CommentRepo,
		),
		LikeUseCase: NewLikeUseCase(
			repos.LikeRepo,
		),
	}, nil
}
