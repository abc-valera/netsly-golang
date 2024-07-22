package commandTransactor

import (
	"context"
	"database/sql"

	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/commandTransactor"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation"
	"gorm.io/gorm"
)

type Dependencies struct {
	GormSqlite     *gorm.DB
	BoilerSqlite   *sql.DB
	LocalFileSaver string
}

type transactor struct {
	Deps Dependencies
}

func New(deps Dependencies) commandTransactor.ITransactor {
	return &transactor{
		Deps: deps,
	}
}

func (t transactor) PerformTX(
	ctx context.Context,
	txFunc func(ctx context.Context, txCommands persistence.Commands, txQueries persistence.Queries) error,
) error {
	// Check every dependency and if it's not empty, start a transaction where possible
	var commandsQueriesDependencies implementation.CommandsAndQueriesDependencies

	var gormTX *gorm.DB
	if t.Deps.GormSqlite != nil {
		gormTX = t.Deps.GormSqlite.WithContext(ctx).Begin()
		if err := gormTX.Error; err != nil {
			return coderr.NewInternalErr(err)
		}
		commandsQueriesDependencies.GormSqlite = gormTX
	}

	var boilerTX *sql.Tx
	if t.Deps.BoilerSqlite != nil {
		var err error
		boilerTX, err = t.Deps.BoilerSqlite.BeginTx(ctx, nil)
		if err != nil {
			return coderr.NewInternalErr(err)
		}
		commandsQueriesDependencies.BoilerSqlite = boilerTX
	}

	// Note, that LocalFileSaver doesn't support transactions
	commandsQueriesDependencies.LocalFileSaver = t.Deps.LocalFileSaver

	txCommands, txQueries, err := implementation.NewCommandsAndQueries(commandsQueriesDependencies)
	if err != nil {
		return coderr.NewInternalErr(err)
	}

	if err := txFunc(ctx, txCommands, txQueries); err != nil {
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
