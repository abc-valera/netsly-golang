package entityTransactor

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain"
	"github.com/abc-valera/netsly-api-golang/internal/domain/entityTransactor"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/boiler/errutil"
)

type transactorImpl struct {
	deps persistence.PeristenceDependencies

	// These are used to create new entities for the transaction
	services domain.Services
}

func NewTransactor(
	deps persistence.PeristenceDependencies,
	services domain.Services,
) entityTransactor.ITransactor {
	return &transactorImpl{
		deps:     deps,
		services: services,
	}
}

func (t transactorImpl) PerformTX(
	ctx context.Context,
	txFunc func(ctx context.Context, entities domain.Entities) error,
) error {
	boilerTX, err := t.deps.Boiler.BeginTx(ctx, nil)
	if err != nil {
		return coderr.NewInternalErr(err)
	}

	txEntities := domain.NewEntities(
		persistence.NewCommands(persistence.CommandsDependencies{
			Boiler: boilerTX,
		}),
		persistence.NewQueries(persistence.QueriesDependencies{
			Boiler: boilerTX,
		}),
		t.services,
	)
	if err := txFunc(ctx, txEntities); err != nil {
		if err := boilerTX.Rollback(); err != nil {
			return coderr.NewInternalErr(err)
		}
		return errutil.HandleErr(err)
	}

	if err := boilerTX.Commit(); err != nil {
		return coderr.NewInternalErr(err)
	}

	return nil
}
