package application

import (
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/domain"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/query"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/transactioneer"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/service"
)

type UseCases struct {
	SignUseCase SignUseCase
}

func NewUseCases(
	queries query.Queries,
	tx transactioneer.ITransactioneer,
	domains domain.Domains,
	services service.Services,
) (UseCases, error) {
	return UseCases{
		SignUseCase: NewSignUseCase(
			queries.User,
			tx,
			domains.User,
			services.PasswordMaker,
			services.TokenMaker,
			services.MessageBroker,
		),
	}, nil
}
