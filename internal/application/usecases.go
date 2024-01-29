package application

import (
	"github.com/abc-valera/netsly-api-golang/internal/core"
	"github.com/abc-valera/netsly-api-golang/internal/core/persistence/transactioneer"
)

type UseCases struct {
	SignUseCase SignUseCase
}

func NewUseCases(
	queries core.Queries,
	tx transactioneer.ITransactioneer,
	domains core.Domains,
	services core.Services,
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
