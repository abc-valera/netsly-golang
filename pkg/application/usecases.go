package application

import (
	"github.com/abc-valera/netsly-api-golang/pkg/domain"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/transactor"
)

type UseCases struct {
	SignUseCase ISignUseCase
}

func NewUseCases(
	tx transactor.ITransactor,
	entities domain.Entities,
	services domain.Services,
) UseCases {
	return UseCases{
		SignUseCase: NewSignUseCase(
			entities.User,
			tx,
			services.PasswordMaker,
			services.TokenMaker,
			services.TaskQueuer,
		),
	}
}
