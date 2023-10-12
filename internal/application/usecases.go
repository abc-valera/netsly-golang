package application

import (
	"github.com/abc-valera/flugo-api-golang/internal/domain/repository"
	"github.com/abc-valera/flugo-api-golang/internal/domain/service"
)

type UseCases struct {
	SignUseCase    SignUseCase
	UserUseCase    UserUseCase
	JokeUseCase    JokeUseCase
	CommentUseCase CommentUseCase
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
	}, nil
}
