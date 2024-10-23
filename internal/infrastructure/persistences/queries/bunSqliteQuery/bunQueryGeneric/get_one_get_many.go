package bunQueryGeneric

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/queryUtil/filter"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/queryUtil/queryGeneric"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/queryUtil/selector"
	"github.com/abc-valera/netsly-golang/internal/domain/util/dto"
	"github.com/uptrace/bun"
)

func NewGetOneGetMany[DomainModel any, DtoModel dto.IDomainable[DomainModel]](
	db bun.IDB,
	newDto func(DomainModel) DtoModel,
) queryGeneric.IGetOneGetMany[DomainModel] {
	return struct {
		getOne[DomainModel, DtoModel]
		getMany[DomainModel, DtoModel]
	}{
		getOne[DomainModel, DtoModel]{db: db, newDto: newDto},
		getMany[DomainModel, DtoModel]{db: db, newDto: newDto},
	}
}

type getOne[DomainModel any, DtoModel dto.IDomainable[DomainModel]] struct {
	db bun.IDB

	newDto func(DomainModel) DtoModel
}

func NewGetOne[DomainModel any, DtoModel dto.IDomainable[DomainModel]](
	db bun.IDB,
	newDto func(DomainModel) DtoModel,
) queryGeneric.IGetOne[DomainModel] {
	return &getOne[DomainModel, DtoModel]{
		db:     db,
		newDto: newDto,
	}
}

func (q getOne[DomainModel, DtoModel]) GetOne(
	ctx context.Context,
	fitlerOptions ...filter.Option[DomainModel],
) (DomainModel, error) {
	filters := filter.New(fitlerOptions...)

	var dto DtoModel
	query := q.db.NewSelect().Model(&dto)

	for _, f := range filters {
		ApplyFilter[DomainModel, DtoModel](query, q.newDto, f)
	}

	err := query.Scan(ctx)

	return dto.ToDomain(), err
}

type getMany[DomainModel any, DtoModel dto.IDomainable[DomainModel]] struct {
	db bun.IDB

	newDomain func(DtoModel) DomainModel
	newDto    func(DomainModel) DtoModel
}

func NewGetMany[DomainModel any, DtoModel dto.IDomainable[DomainModel]](
	db bun.IDB,

	newDomain func(DtoModel) DomainModel,
	newDto func(DomainModel) DtoModel,
) queryGeneric.IGetMany[DomainModel] {
	return &getMany[DomainModel, DtoModel]{
		db: db,

		newDomain: newDomain,
		newDto:    newDto,
	}
}

func (q getMany[DomainModel, DtoModel]) GetMany(
	ctx context.Context,
	selectorOptions ...selector.Option[DomainModel],
) ([]DomainModel, error) {
	selector := selector.New(selectorOptions...)

	var bunModel []DtoModel
	query := q.db.NewSelect().Model(&bunModel)

	ApplySelector[DomainModel, DtoModel](query, q.newDto, selector)

	err := query.Scan(ctx)

	return dto.NewDomainModels(bunModel), err
}
