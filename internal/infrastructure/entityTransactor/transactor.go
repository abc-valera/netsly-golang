package entityTransactor

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/internal/domain"
	"github.com/abc-valera/netsly-api-golang/internal/domain/entityTransactor"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence"
	domainCommandTransactor "github.com/abc-valera/netsly-api-golang/internal/domain/persistence/commandTransactor"
	infraCommandTransactor "github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/commandTransactor"
)

type transactorImpl struct {
	commandTransactor domainCommandTransactor.ITransactor

	// These are used to create new entities for the transaction
	queries  persistence.Queries
	services domain.Services
}

func New(
	commandTransactor domainCommandTransactor.ITransactor,
	queries persistence.Queries,
	services domain.Services,
) entityTransactor.ITransactor {
	return &transactorImpl{
		commandTransactor: commandTransactor,
		services:          services,
	}
}

func (t transactorImpl) PerformTX(
	ctx context.Context,
	txFunc func(ctx context.Context, entities domain.Entities) error,
) error {
	commandTransactorTxFunc := func(ctx context.Context, txCommands persistence.Commands) error {
		txEntities := domain.NewEntities(
			txCommands,
			infraCommandTransactor.NewAntiNested(),
			t.queries,
			t.services,
		)

		if err := txFunc(ctx, txEntities); err != nil {
			return err
		}

		return nil
	}

	return t.commandTransactor.PerformTX(ctx, commandTransactorTxFunc)
}
