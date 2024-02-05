package enttransactioneer

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/ent"

	errhandler "github.com/abc-valera/netsly-api-golang/internal/adapter/persistence/ent-impl/errors"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/transactioneer"
)

type transactioneerImpl struct {
	client *ent.Client
}

func NewTransactioneer(client *ent.Client) transactioneer.ITransactioneer {
	return &transactioneerImpl{
		client: client,
	}
}

func (t transactioneerImpl) PerformTX(ctx context.Context, txFunc func(ctx context.Context) error) error {
	tx, err := t.client.Tx(ctx)
	if err != nil {
		return errhandler.HandleErr(err)
	}

	oldClient := *t.client
	defer func() {
		*t.client = oldClient
	}()

	*t.client = *tx.Client()
	if err := txFunc(ctx); err != nil {
		if err := tx.Rollback(); err != nil {
			return errhandler.HandleErr(err)
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return errhandler.HandleErr(err)
	}
	return nil
}
