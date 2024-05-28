package application

import (
	"github.com/abc-valera/netsly-api-golang/internal/domain"
	"github.com/abc-valera/netsly-api-golang/internal/domain/entityTransactor"
)

type UseCases struct {
	SignUseCase ISignUseCase
}

func NewUseCases(
	tx entityTransactor.ITransactor,
	entities domain.Entities,
	services domain.Services,
) UseCases {
	return UseCases{
		SignUseCase: NewSignUseCase(
			entities.User,
			tx,
			services.PasswordMaker,
			services.TaskQueuer,
		),
	}
}
