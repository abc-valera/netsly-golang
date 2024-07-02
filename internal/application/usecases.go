package application

import (
	"github.com/abc-valera/netsly-api-golang/internal/domain"
	"github.com/abc-valera/netsly-api-golang/internal/domain/transactor"
)

type Usecases struct {
	SignUsecase ISignUsecase
}

func NewUsecases(
	tx transactor.ITransactor,
	entities domain.Entities,
	services domain.Services,
) Usecases {
	return Usecases{
		SignUsecase: NewSignUsecase(
			entities.User,
			tx,
			services.PasswordMaker,
			services.TaskQueuer,
		),
	}
}
