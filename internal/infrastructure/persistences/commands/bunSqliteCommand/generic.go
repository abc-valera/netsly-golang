// bunSqliteCommand provides generic struct types for implementing commands interfaces.
// To instantiate the struct the DTOs need to be defined.
//
// Package provides the following structs:
//   - CreateUpdateDelete[DomainModel any, UpdateModel any, BunModel domainable[DomainModel]]
//   - CreateDelete[DomainModel any, BunModel domainable[DomainModel]]
//
// Where:
//   - Base Model is the domain model
//   - UpdateModel is the model used for updating the domain model
//   - BunModel is the model used for bun orm
//
// Also these structs needs to be provided:
//   - dto func(DomainModel) BunModel
//   - dtoUpdate func(UpdateModel) (BunModel, []string)
package bunSqliteCommand

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/dependencies/bunSqlite/bunSqliteErrors"
	"github.com/uptrace/bun"
)

func New[DomainModel, BunModel any](
	db bun.IDB,
	dto func(DomainModel) BunModel,
) command.ICreateUpdateDelete[DomainModel] {
	return struct {
		create[DomainModel, BunModel]
		update[DomainModel, BunModel]
		delete[DomainModel, BunModel]
	}{
		create[DomainModel, BunModel]{db, dto},
		update[DomainModel, BunModel]{db, dto},
		delete[DomainModel, BunModel]{db, dto},
	}
}

type create[DomainModel, BunModel any] struct {
	db  bun.IDB
	dto func(DomainModel) BunModel
}

func (c create[DomainModel, BunModel]) Create(ctx context.Context, req DomainModel) error {
	bunModel := c.dto(req)
	res, err := c.db.NewInsert().Model(&bunModel).Exec(ctx)
	return bunSqliteErrors.HandleCommandResult(res, err)
}

type update[DomainModel, BunModel any] struct {
	db  bun.IDB
	dto func(DomainModel) BunModel
}

func (u update[DomainModel, BunModel]) Update(ctx context.Context, req DomainModel) error {
	bunModel := u.dto(req)
	res, err := u.db.NewUpdate().Model(&bunModel).WherePK().Exec(ctx)
	return bunSqliteErrors.HandleCommandResult(res, err)
}

type delete[DomainModel, BunModel any] struct {
	db  bun.IDB
	dto func(DomainModel) BunModel
}

func (d delete[DomainModel, BunModel]) Delete(ctx context.Context, req DomainModel) error {
	bunModel := d.dto(req)
	res, err := d.db.NewDelete().Model(&bunModel).WherePK().Exec(ctx)
	return bunSqliteErrors.HandleCommandResult(res, err)
}
