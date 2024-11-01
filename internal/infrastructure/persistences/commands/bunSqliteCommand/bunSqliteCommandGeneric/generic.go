package bunSqliteCommandGeneric

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/domain/global"
	"github.com/abc-valera/netsly-golang/internal/domain/util/dto"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/dependencies/bunSqlite/bunSqliteErrors"
	"github.com/uptrace/bun"
	"go.opentelemetry.io/otel/trace"
)

type Command[Model any, DtoModel dto.IDomainable[Model]] struct {
	db  bun.IDB
	dto func(Model) DtoModel
}

func New[Model any, DtoModel dto.IDomainable[Model]](
	db bun.IDB,
	dto func(Model) DtoModel,
) Command[Model, DtoModel] {
	return Command[Model, DtoModel]{
		db:  db,
		dto: dto,
	}
}

func (c Command[Model, DtoModel]) Create(ctx context.Context, model Model) error {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	bunDto := c.dto(model)
	res, err := c.db.NewInsert().Model(&bunDto).Exec(ctx)
	return bunSqliteErrors.HandleCommandResult(res, err)
}

func (c Command[Model, DtoModel]) Update(ctx context.Context, model Model) error {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	bunDto := c.dto(model)
	res, err := c.db.NewUpdate().Model(&bunDto).WherePK().Exec(ctx)
	return bunSqliteErrors.HandleCommandResult(res, err)
}

func (c Command[Model, DtoModel]) Delete(ctx context.Context, model Model) error {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	bunDto := c.dto(model)
	res, err := c.db.NewDelete().Model(&bunDto).WherePK().Exec(ctx)
	return bunSqliteErrors.HandleCommandResult(res, err)
}
