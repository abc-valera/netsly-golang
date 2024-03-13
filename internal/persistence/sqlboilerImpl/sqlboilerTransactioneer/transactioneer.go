package sqlboilerTransactioneer

import (
	"context"
	"database/sql"

	"github.com/abc-valera/netsly-api-golang/internal/domain"
	"github.com/abc-valera/netsly-api-golang/internal/domain/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/transactioneer"
	"github.com/abc-valera/netsly-api-golang/internal/persistence/sqlboilerImpl/errors"
	sqlboilercommand "github.com/abc-valera/netsly-api-golang/internal/persistence/sqlboilerImpl/sqlboilerCommand"
	"github.com/abc-valera/netsly-api-golang/internal/persistence/sqlboilerImpl/sqlboilerQuery"
)

type transactioneerImpl struct {
	db *sql.DB

	// These are used to create new entities for the transaction
	services domain.Services
}

func NewTransactioneer(db *sql.DB, services domain.Services) transactioneer.ITransactioneer {
	return &transactioneerImpl{
		db: db,
	}
}

func (t *transactioneerImpl) PerformTX(ctx context.Context, txFunc func(ctx context.Context, entities domain.Entities) error) error {
	tx, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		return coderr.NewInternalErr(err)
	}

	txEntities := domain.NewEntities(
		sqlboilercommand.NewCommands(tx),
		sqlboilerQuery.NewQueries(tx),
		t.services,
	)
	if err := txFunc(ctx, txEntities); err != nil {
		if err := tx.Rollback(); err != nil {
			return coderr.NewInternalErr(err)
		}
		return errors.HandleErr(err)
	}

	if err := tx.Commit(); err != nil {
		return coderr.NewInternalErr(err)
	}

	return nil
}
