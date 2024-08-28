package transactors

import (
	"context"
	"database/sql"

	"github.com/abc-valera/netsly-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-golang/internal/core/commandTransactor"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/commandsAndQueries"
	"github.com/uptrace/bun"
	"gorm.io/gorm"
)

type transactor struct {
	Deps Dependencies
}

func NewCommandTransactor(deps Dependencies) commandTransactor.ITransactor {
	return &transactor{
		Deps: deps,
	}
}

func (t transactor) PerformTX(
	ctx context.Context,
	txFunc func(ctx context.Context, txCommands persistence.Commands, txQueries persistence.Queries) error,
) error {
	// Check every dependency and if it's not empty, start a transaction where possible
	var commandsQueriesDependencies commandsAndQueries.Dependencies

	var gormTX *gorm.DB
	if t.Deps.GormSqlite != nil {
		gormTX = t.Deps.GormSqlite.WithContext(ctx).Begin()
		if err := gormTX.Error; err != nil {
			return coderr.NewInternalErr(err)
		}
		commandsQueriesDependencies.GormSqlite = gormTX
	}

	var bunTX bun.Tx
	if t.Deps.BunSqlite != nil {
		var err error
		bunTX, err = t.Deps.BunSqlite.BeginTx(ctx, nil)
		if err != nil {
			return coderr.NewInternalErr(err)
		}
		commandsQueriesDependencies.BunSqlite = bunTX
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

	txCommands, txQueries, err := commandsAndQueries.New(commandsQueriesDependencies)
	if err != nil {
		return coderr.NewInternalErr(err)
	}

	if err := txFunc(ctx, txCommands, txQueries); err != nil {
		if gormTX != nil {
			if err := gormTX.Rollback().Error; err != nil {
				return coderr.NewInternalErr(err)
			}
		}
		if bunTX != (bun.Tx{}) {
			if err := bunTX.Rollback(); err != nil {
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

	if bunTX != (bun.Tx{}) {
		if err := bunTX.Commit(); err != nil {
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
