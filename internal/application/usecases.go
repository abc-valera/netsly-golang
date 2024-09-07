package application

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/domain/entity"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence"
	"github.com/abc-valera/netsly-golang/internal/domain/service"
)

type Usecases struct {
	SignUsecase ISignUsecase
}

func NewUsecases(dep IDependency) Usecases {
	return Usecases{
		SignUsecase: newSignUsecase(dep),
	}
}

type IDependency interface {
	BeginTX(ctx context.Context) (IDependencyTX, error)
	RunInTX(
		ctx context.Context,
		fn func(context.Context, entity.Entities, Usecases) error,
	) error

	E() entity.Entities
	U() Usecases
}

type IDependencyTX interface {
	IDependency

	Commit() error
	Rollback() error
}

type dependency struct {
	db       persistence.IDB
	services service.Services
}

func NewDependency(db persistence.IDB, services service.Services) IDependency {
	return dependency{
		db:       db,
		services: services,
	}
}

func (dep dependency) BeginTX(ctx context.Context) (IDependencyTX, error) {
	tx, err := dep.db.BeginTX(ctx)
	if err != nil {
		return nil, err
	}

	return dependencyTX{
		tx:       tx,
		services: dep.services,
	}, nil
}

func (dep dependency) RunInTX(
	ctx context.Context,
	fn func(context.Context, entity.Entities, Usecases) error,
) error {
	tx, err := dep.db.BeginTX(ctx)
	if err != nil {
		return err
	}

	if err := fn(
		ctx,
		entity.NewEntities(entity.NewDependency(tx, dep.services)),
		NewUsecases(NewDependency(tx, dep.services)),
	); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (dep dependency) U() Usecases {
	return NewUsecases(dep)
}

func (dep dependency) E() entity.Entities {
	return entity.NewEntities(entity.NewDependency(dep.db, dep.services))
}

type dependencyTX struct {
	tx       persistence.ITX
	services service.Services
}

func (depTX dependencyTX) RunInTX(
	ctx context.Context,
	fn func(context.Context, entity.Entities, Usecases) error,
) error {
	nestedTX, err := depTX.tx.BeginTX(ctx)
	if err != nil {
		return err
	}

	if err := fn(
		ctx,
		entity.NewEntities(entity.NewDependency(nestedTX, depTX.services)),
		NewUsecases(NewDependency(nestedTX, depTX.services)),
	); err != nil {
		nestedTX.Rollback()
		return err
	}

	return nestedTX.Commit()
}

func (depTX dependencyTX) BeginTX(ctx context.Context) (IDependencyTX, error) {
	tx, err := depTX.tx.BeginTX(ctx)
	if err != nil {
		return nil, err
	}

	return dependencyTX{
		tx:       tx,
		services: depTX.services,
	}, nil
}

func (depTX dependencyTX) Rollback() error {
	return depTX.tx.Rollback()
}

func (depTX dependencyTX) Commit() error {
	return depTX.tx.Commit()
}

func (depTX dependencyTX) U() Usecases {
	return NewUsecases(depTX)
}

func (depTX dependencyTX) E() entity.Entities {
	return entity.NewEntities(entity.NewDependency(depTX.tx, depTX.services))
}

func (depTX dependencyTX) Services() service.Services {
	return depTX.services
}
