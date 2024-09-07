package persistences

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/domain/global"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-golang/internal/domain/util/coderr"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/commands"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/queries"
	"github.com/uptrace/bun"
	"gorm.io/gorm"
)

// TX is a struct that contains the transactions from all datasources.
type TX struct {
	gormSqlite     *gorm.DB
	bunSqlite      *bun.Tx
	localFileSaver string // Note, that LocalFileSaver doesn't support transactions
}

func newTX(ctx context.Context, db DB) (returnTX TX, returnErr error) {
	// If one of the transactions fails to start, rollback every other one.
	defer func() {
		if returnErr != nil {
			returnTX.Rollback()
		}
	}()

	if db.gormSqlite != nil {
		returnTX.gormSqlite = db.gormSqlite.WithContext(ctx).Begin()
		if err := returnTX.gormSqlite.Error; err != nil {
			return returnTX, coderr.NewInternalErr(err)
		}
	}

	if db.bunSqlite != nil {
		bunTX, err := db.bunSqlite.BeginTx(ctx, nil)
		if err != nil {
			return returnTX, coderr.NewInternalErr(err)
		}
		returnTX.bunSqlite = &bunTX
	}

	if db.localFileSaver != "" {
		// Note, that LocalFileSaver doesn't support transactions
		returnTX.localFileSaver = db.localFileSaver
	}

	return returnTX, nil
}

func newNestedTX(ctx context.Context, tx TX) (returnTX TX, returnErr error) {
	// If one of the transactions fails to start, rollback every other one.
	defer func() {
		if returnErr != nil {
			returnTX.Rollback()
		}
	}()

	if tx.gormSqlite != nil {
		returnTX.gormSqlite = tx.gormSqlite.WithContext(ctx).Begin()
		if err := returnTX.gormSqlite.Error; err != nil {
			return returnTX, coderr.NewInternalErr(err)
		}
	}

	if tx.bunSqlite != nil {
		bunNestedTx, err := tx.bunSqlite.BeginTx(ctx, nil)
		if err != nil {
			return returnTX, coderr.NewInternalErr(err)
		}
		returnTX.bunSqlite = &bunNestedTx
	}

	if tx.localFileSaver != "" {
		// Note, that LocalFileSaver doesn't support transactions at all
		returnTX.localFileSaver = tx.localFileSaver
	}

	return returnTX, nil
}

// Commit commits all the transactions.
// It returns an error only if transactions were already committed/rolledback.
//
// Other errors are handled gracefully, so that the Eventual Consistency is guaranteed.
func (tx TX) Commit() error {
	// TODO: gracefully handle commit errors

	// Note, that here we want to handle commit errors.
	// If one of the transactions fails to commit (after some of them have already committed),
	// we want to queue the failed transation and repeat it till it succeeds.
	// That way the Eventual Consistency is guaranteed.
	//
	// We want to return error only if transactions were already committed/rolledback.
	if tx.gormSqlite != nil {
		if err := tx.gormSqlite.Commit().Error; err != nil {
			global.Log().Error("gorm commit error", "err", err)
		}
	}
	if tx.bunSqlite != nil {
		if err := tx.bunSqlite.Commit(); err != nil {
			global.Log().Error("bun commit error", "err", err)
		}
	}

	return nil
}

// Rollback rolls back all the transactions.
// It returns an error only if transactions were already committed/rolledback.
func (tx TX) Rollback() error {
	// Note, that we don't want to handle rollback errors.
	// Eventually, transaction connections will be closed anyway.
	if tx.gormSqlite != nil {
		if err := tx.gormSqlite.Rollback().Error; err != nil {
			global.Log().Error("gorm rollback error", "err", err)
		}
	}
	if tx.bunSqlite != nil {
		if err := tx.bunSqlite.Rollback(); err != nil {
			global.Log().Error("bun rollback error", "err", err)
		}
	}

	return nil
}

func (tx TX) RunInTX(
	ctx context.Context,
	fn func(context.Context, command.Commands, query.Queries) error,
) error {
	nestedTX, err := newNestedTX(ctx, tx)
	if err != nil {
		return err
	}

	if err := fn(ctx, nestedTX.Commands(), nestedTX.Queries()); err != nil {
		nestedTX.Rollback()
		return err
	}

	return nestedTX.Commit()
}

func (tx TX) BeginTX(ctx context.Context) (persistence.ITX, error) {
	return newNestedTX(ctx, tx)
}

func (tx TX) Commands() command.Commands {
	return commands.New(commands.Dependency{
		GormSqlite:     tx.gormSqlite,
		BunSqlite:      tx.bunSqlite,
		LocalFileSaver: tx.localFileSaver,
	})
}

func (tx TX) Queries() query.Queries {
	return queries.New(queries.Dependency{
		GormSqlite:     tx.gormSqlite,
		BunSqlite:      tx.bunSqlite,
		LocalFileSaver: tx.localFileSaver,
	})
}
