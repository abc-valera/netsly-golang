// bunSqliteCommand provides generic struct types for implementing commands interfaces.
// To instantiate the struct the DTOs need to be defined.
//
// Package provides the following structs:
//   - CreateUpdateDelete[BaseModel any, UpdateModel any, BunModel domainable[BaseModel]]
//   - CreateDelete[BaseModel any, BunModel domainable[BaseModel]]
//
// Where:
//   - Base Model is the domain model
//   - UpdateModel is the model used for updating the domain model
//   - BunModel is the model used for bun orm
//
// Also these structs needs to be provided:
//   - dto func(BaseModel) BunModel
//   - dtoUpdate func(UpdateModel) (BunModel, []string)
package bunSqliteCommand

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/dependencies/bunSqlite/bunSqliteErrutil"
	"github.com/uptrace/bun"
)

type createUpdateDelete[BaseModel, UpdateModel any, BunModel domainable[BaseModel]] struct {
	create[BaseModel, BunModel]
	update[BaseModel, UpdateModel, BunModel]
	delete[BaseModel, BunModel]
}

func NewCreateUpdateDelete[BaseModel, UpdateModel any, BunModel domainable[BaseModel]](
	db bun.IDB,
	dto func(BaseModel) BunModel,
	dtoUpdate func(BaseModel, UpdateModel) (BunModel, []string),
) createUpdateDelete[BaseModel, UpdateModel, BunModel] {
	return createUpdateDelete[BaseModel, UpdateModel, BunModel]{
		create[BaseModel, BunModel]{db, dto},
		update[BaseModel, UpdateModel, BunModel]{db, dtoUpdate},
		delete[BaseModel, BunModel]{db, dto},
	}
}

type createDelete[BaseModel any, BunModel domainable[BaseModel]] struct {
	create[BaseModel, BunModel]
	delete[BaseModel, BunModel]
}

func NewCreateDelete[BaseModel any, BunModel domainable[BaseModel]](
	db bun.IDB,
	dto func(BaseModel) BunModel,
) createDelete[BaseModel, BunModel] {
	return createDelete[BaseModel, BunModel]{
		create[BaseModel, BunModel]{db, dto},
		delete[BaseModel, BunModel]{db, dto},
	}
}

type create[BaseModel any, BunModel domainable[BaseModel]] struct {
	db bun.IDB

	dto func(BaseModel) BunModel
}

func (c create[BaseModel, BunModel]) Create(ctx context.Context, model BaseModel) (BaseModel, error) {
	bunModel := c.dto(model)
	res, err := c.db.NewInsert().Model(&bunModel).Exec(ctx)
	return bunModel.ToDomain(), bunSqliteErrutil.HandleCommandResult(res, err)
}

type update[BaseModel, UpdateModel any, BunModel domainable[BaseModel]] struct {
	db bun.IDB

	dto func(BaseModel, UpdateModel) (BunModel, []string)
}

func (u update[BaseModel, UpdateModel, BunModel]) Update(ctx context.Context, ids BaseModel, updateModel UpdateModel) (BaseModel, error) {
	bunModel, columns := u.dto(ids, updateModel)
	res, err := u.db.NewUpdate().Model(&bunModel).Column(columns...).WherePK().Exec(ctx)
	return bunModel.ToDomain(), bunSqliteErrutil.HandleCommandResult(res, err)
}

type delete[BaseModel any, BunModel domainable[BaseModel]] struct {
	db bun.IDB

	dto func(BaseModel) BunModel
}

func (d delete[BaseModel, BunModel]) Delete(ctx context.Context, model BaseModel) error {
	bunModel := d.dto(model)
	res, err := d.db.NewDelete().Model(&bunModel).WherePK().Exec(ctx)
	return bunSqliteErrutil.HandleCommandResult(res, err)
}

type domainable[BaseModel any] interface {
	ToDomain() BaseModel
}
