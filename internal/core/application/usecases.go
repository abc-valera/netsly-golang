package application

import (
	"github.com/abc-valera/flugo-api-golang/internal/core/domain"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/persistence/transactioneer"
)

type UseCases struct {
	SignUseCase SignUseCase
}

func NewUseCases(
	queries domain.Queries,
	tx transactioneer.ITransactioneer,
	domains domain.Domains,
	services domain.Services,
) UseCases {
	return UseCases{
		SignUseCase: NewSignUseCase(
			queries.User,
			tx,
			domains.User,
			services.PasswordMaker,
			services.TokenMaker,
			services.MessageBroker,
		),
	}
}
