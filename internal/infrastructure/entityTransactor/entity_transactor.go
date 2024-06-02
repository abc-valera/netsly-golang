package entityTransactor

import (
	"context"
	"database/sql"

	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain"
	"github.com/abc-valera/netsly-api-golang/internal/domain/entityTransactor"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/boiler/errutil"
)

type transactorImpl struct {
	db *sql.DB

	// These are used to create new entities for the transaction
	services domain.Services
}

func NewTransactor(db *sql.DB, services domain.Services) entityTransactor.ITransactor {
	return &transactorImpl{
		db:       db,
		services: services,
	}
}

func (t *transactorImpl) PerformTX(ctx context.Context, txFunc func(ctx context.Context, entities domain.Entities) error) error {
	tx, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		return coderr.NewInternalErr(err)
	}

	txEntities := domain.NewEntities(
		persistence.NewCommands(tx),
		persistence.NewQueries(tx),
		t.services,
	)
	if err := txFunc(ctx, txEntities); err != nil {
		if err := tx.Rollback(); err != nil {
			return coderr.NewInternalErr(err)
		}
		return errutil.HandleErr(err)
	}

	if err := tx.Commit(); err != nil {
		return coderr.NewInternalErr(err)
	}

	return nil
}
