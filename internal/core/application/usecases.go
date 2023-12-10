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
	tx transactioneer.ITransactioneer,
	queries query.Queries,
	domains domain.Domains,
	services service.Services,
) (UseCases, error) {
	return UseCases{
		SignUseCase: NewSignUseCase(
			tx,
			queries.User,
			domains.User,
			services.PasswordMaker,
			services.TokenMaker,
			services.MessageBroker,
		),
	}, nil
}
