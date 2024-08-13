package transactors

import (
	"context"
	"database/sql"

	"github.com/abc-valera/netsly-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-golang/internal/core/entityTransactor"
	"github.com/abc-valera/netsly-golang/internal/domain"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/commandsAndQueries"
	"gorm.io/gorm"
)

type entityTransactorImpl struct {
	// deps are used to start transactions among all persistence datastores.
	// Note, the same dependencies as in the commandTransactor are used.
	deps Dependencies

	// services are used to create new entities for the transaction
	services domain.Services
}

func NewEntityTransactor(deps Dependencies, services domain.Services) entityTransactor.ITransactor {
	return &entityTransactorImpl{
		deps:     deps,
		services: services,
	}
}

func (t entityTransactorImpl) PerformTX(
	ctx context.Context,
	txFunc func(ctx context.Context, entities domain.Entities) error,
) error {
	// Check every dependency and if it's not empty, start a transaction where possible
	var commandsQueriesDependencies commandsAndQueries.Dependencies
	var transactorsDependencies Dependencies

	var gormTX *gorm.DB
	if t.deps.GormSqlite != nil {
		gormTX = t.deps.GormSqlite.WithContext(ctx).Begin()
		if err := gormTX.Error; err != nil {
			return coderr.NewInternalErr(err)
		}
		commandsQueriesDependencies.GormSqlite = gormTX
		transactorsDependencies.GormSqlite = gormTX
	}

	var boilerTX *sql.Tx
	if t.deps.BoilerSqlite != nil {
		var err error
		boilerTX, err = t.deps.BoilerSqlite.BeginTx(ctx, nil)
		if err != nil {
			return coderr.NewInternalErr(err)
		}
		commandsQueriesDependencies.BoilerSqlite = boilerTX
	}

	// Note, that SQLBoiler doesn't support nested transactions
	transactorsDependencies.BoilerSqlite = t.deps.BoilerSqlite

	// Note, that LocalFileSaver doesn't support transactions
	commandsQueriesDependencies.LocalFileSaver = t.deps.LocalFileSaver
	transactorsDependencies.LocalFileSaver = t.deps.LocalFileSaver

	txCommands, txQueries, err := commandsAndQueries.New(commandsQueriesDependencies)
	if err != nil {
		return coderr.NewInternalErr(err)
	}

	commandTransactor := NewCommandTransactor(transactorsDependencies)

	if err := txFunc(ctx, domain.NewEntities(
		txCommands,
		commandTransactor,
		txQueries,
		t.services,
	)); err != nil {
		if gormTX != nil {
			if err := gormTX.Rollback().Error; err != nil {
				return coderr.NewInternalErr(err)
			}
		}
		if boilerTX != nil {
			if err := boilerTX.Rollback(); err != nil {
				return coderr.NewInternalErr(err)
			}
		}
		return err
	}

	if gormTX != nil {
		if err := gormTX.Commit().Error; err != nil {
			return coderr.NewInternalErr(err)
		}
	}

	if boilerTX != nil {
		if err := boilerTX.Commit(); err != nil {
			return coderr.NewInternalErr(err)
		}
	}

	return nil
}
