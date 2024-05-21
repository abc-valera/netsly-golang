package commandTransactor

import (
	"context"
	"database/sql"

	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
	domainPersistence "github.com/abc-valera/netsly-api-golang/internal/domain/persistence"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/commandTransactor"
	infraPersistence "github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence"
)

type transactor struct {
	db *sql.DB
}

func NewCommandTransactor(db *sql.DB) commandTransactor.ITransactor {
	return &transactor{
		db: db,
	}
}

func (t transactor) PerformTX(
	ctx context.Context,
	txFunc func(ctx context.Context, txCommands domainPersistence.Commands) error,
) error {
	tx, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		return coderr.NewInternalErr(err)
	}

	txCommands := infraPersistence.NewCommands(tx)
	if err := txFunc(ctx, txCommands); err != nil {
		if err := tx.Rollback(); err != nil {
			return coderr.NewInternalErr(err)
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return coderr.NewInternalErr(err)
	}

	return nil
}
