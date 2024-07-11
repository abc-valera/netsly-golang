package commandTransactor

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/commandTransactor"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation"
)

type transactor struct {
	deps implementation.PersistenceDependencies
}

func New(deps implementation.PersistenceDependencies) commandTransactor.ITransactor {
	return &transactor{
		deps: deps,
	}
}

func (t transactor) PerformTX(
	ctx context.Context,
	txFunc func(ctx context.Context, txCommands persistence.Commands) error,
) error {
	boilerTX, err := t.deps.Boiler.BeginTx(ctx, nil)
	if err != nil {
		return coderr.NewInternalErr(err)
	}

	txCommands := implementation.NewCommands(implementation.CommandsDependencies{
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
