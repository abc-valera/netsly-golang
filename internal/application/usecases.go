package application

import (
	"github.com/abc-valera/netsly-golang/internal/core/entityTransactor"
	"github.com/abc-valera/netsly-golang/internal/domain"
)

type Usecases struct {
	SignUsecase ISignUsecase
}

func NewUsecases(
	tx entityTransactor.ITransactor,
	entities domain.Entities,
	services domain.Services,
) Usecases {
	return Usecases{
		SignUsecase: NewSignUsecase(
			entities.User,
			tx,
			services.Passworder,
			services.TaskQueuer,
		),
	}
}
