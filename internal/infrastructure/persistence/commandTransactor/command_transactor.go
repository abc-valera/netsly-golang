package commandTransactor

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
	domainPersistence "github.com/abc-valera/netsly-api-golang/internal/domain/persistence"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/commandTransactor"
	infraPersistence "github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence"
)

type transactor struct {
	deps infraPersistence.PeristenceDependencies
}

func New(deps infraPersistence.PeristenceDependencies) commandTransactor.ITransactor {
	return &transactor{
		deps: deps,
	}
}

func (t transactor) PerformTX(
	ctx context.Context,
	txFunc func(ctx context.Context, txCommands domainPersistence.Commands) error,
) error {
	boilerTX, err := t.deps.Boiler.BeginTx(ctx, nil)
	if err != nil {
		return coderr.NewInternalErr(err)
	}

	txCommands := infraPersistence.NewCommands(infraPersistence.CommandsDependencies{
		Boiler: boilerTX,
	})
	if err := txFunc(ctx, txCommands); err != nil {
		if err := boilerTX.Rollback(); err != nil {
			return coderr.NewInternalErr(err)
		}
		return err
	}

	if err := boilerTX.Commit(); err != nil {
		return coderr.NewInternalErr(err)
	}

	return nil
}
