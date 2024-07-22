package entityTransactor

import (
	"context"
	"database/sql"

	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain"
	"github.com/abc-valera/netsly-api-golang/internal/domain/entityTransactor"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/commandTransactor"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation"
	"gorm.io/gorm"
)

type Dependencies struct {
	GormSqlite     *gorm.DB
	BoilerSqlite   *sql.DB
	LocalFileSaver string
}

type entityTransactorImpl struct {
	deps Dependencies

	// services are used to create new entities for the transaction
	services domain.Services
}

func New(deps Dependencies, services domain.Services) entityTransactor.ITransactor {
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
	var commandsQueriesDependencies implementation.CommandsAndQueriesDependencies
	var commandTransactorDependencies commandTransactor.Dependencies

	var gormTX *gorm.DB
	if t.deps.GormSqlite != nil {
		gormTX = t.deps.GormSqlite.WithContext(ctx).Begin()
		if err := gormTX.Error; err != nil {
			return coderr.NewInternalErr(err)
		}
		commandsQueriesDependencies.GormSqlite = gormTX
		commandTransactorDependencies.GormSqlite = gormTX
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
	commandTransactorDependencies.BoilerSqlite = t.deps.BoilerSqlite

	// Note, that LocalFileSaver doesn't support transactions
	commandsQueriesDependencies.LocalFileSaver = t.deps.LocalFileSaver
	commandTransactorDependencies.LocalFileSaver = t.deps.LocalFileSaver

	txCommands, txQueries, err := implementation.NewCommandsAndQueries(commandsQueriesDependencies)
	if err != nil {
		return coderr.NewInternalErr(err)
	}

	commandTransactor := commandTransactor.New(commandTransactorDependencies)

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
