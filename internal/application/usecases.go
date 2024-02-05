package application

import (
	"github.com/abc-valera/netsly-api-golang/internal/domain"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/transactioneer"
)

type UseCases struct {
	SignUseCase SignUseCase
}

func NewUseCases(
	queries domain.Queries,
	tx transactioneer.ITransactioneer,
	entities domain.Entities,
	services domain.Services,
) UseCases {
	return UseCases{
		SignUseCase: NewSignUseCase(
			queries.User,
			tx,
			entities.User,
			services.PasswordMaker,
			services.TokenMaker,
			services.MessageBroker,
		),
	}
}
