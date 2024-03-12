package sqlboilerTransactioneer

import (
	"context"
	"database/sql"

	"github.com/abc-valera/netsly-api-golang/internal/domain"
	"github.com/abc-valera/netsly-api-golang/internal/domain/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/transactioneer"
	"github.com/abc-valera/netsly-api-golang/internal/persistence/sqlboilerImpl/errors"
	sqlboilercommand "github.com/abc-valera/netsly-api-golang/internal/persistence/sqlboilerImpl/sqlboilerCommand"
)

type transactioneerImpl struct {
	db *sql.DB
}

func NewTransactioneer(db *sql.DB) transactioneer.ITransactioneer {
	return &transactioneerImpl{
		db: db,
	}
}

func (t *transactioneerImpl) PerformTX(ctx context.Context, txFunc func(ctx context.Context, commands domain.Commands) error) error {
	tx, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		return coderr.NewInternalErr(err)
	}

	if err := txFunc(ctx, sqlboilercommand.NewCommands(tx)); err != nil {
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
