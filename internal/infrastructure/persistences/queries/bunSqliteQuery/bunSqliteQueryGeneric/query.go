package bunSqliteQueryGeneric

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/domain/global"
	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/queryUtil/queryGeneric"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/queryUtil/selector"
	"github.com/abc-valera/netsly-golang/internal/domain/util/dto"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/dependencies/bunSqlite/bunSqliteDto"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/dependencies/bunSqlite/bunSqliteErrors"
	"github.com/uptrace/bun"
	"go.opentelemetry.io/otel/trace"
)

type Query[DomainModel any, DtoModel dto.IDomainable[DomainModel]] struct {
	db     bun.IDB
	newDto func(DomainModel) DtoModel
}

// Check if Query implements bunQueryGeneric.Query interface
var _ queryGeneric.IGetOneGetMany[model.User] = Query[model.User, bunSqliteDto.User]{}

func New[DomainModel any, DtoModel dto.IDomainable[DomainModel]](
	db bun.IDB,
	newDto func(DomainModel) DtoModel,
) Query[DomainModel, DtoModel] {
	return Query[DomainModel, DtoModel]{
		db:     db,
		newDto: newDto,
	}
}

func (q Query[DomainModel, DtoModel]) Get(
	ctx context.Context,
	domain DomainModel,
) (DomainModel, error) {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	var dto DtoModel
	query := q.db.NewSelect().Model(&dto)

	if err := ApplyFilter(query, q.newDto, selector.Filter[DomainModel]{By: domain}); err != nil {
		return dto.ToDomain(), bunSqliteErrors.HandleQueryResult(err)
	}

	if err := query.Scan(ctx); err != nil {
		return dto.ToDomain(), bunSqliteErrors.HandleQueryResult(err)
	}

	return dto.ToDomain(), nil
}

func (q Query[DomainModel, DtoModel]) GetMany(
	ctx context.Context,
	options ...selector.Option[DomainModel],
) ([]DomainModel, error) {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	var dtos []DtoModel
	query := q.db.NewSelect().Model(&dtos)

	if err := q.ScanWithSelector(ctx, query, options...); err != nil {
		return nil, bunSqliteErrors.HandleQueryResult(err)
	}

	return dto.NewDomainModels(dtos), nil
}

func (q Query[DomainModel, DtoModel]) ScanWithSelector(
	ctx context.Context,
	query *bun.SelectQuery,
	options ...selector.Option[DomainModel],
) error {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	selector := selector.New(options...)

	ApplyPaging(query, selector.Paging)

	for _, filter := range selector.Filters {
		if err := ApplyFilter(query, q.newDto, filter); err != nil {
			return err
		}
	}

	for _, order := range selector.Orders {
		if err := ApplyOrder(query, q.newDto, order); err != nil {
			return err
		}
	}

	return query.Scan(ctx)
}
