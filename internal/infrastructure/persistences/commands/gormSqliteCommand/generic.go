package gormSqliteCommand

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/dependencies/gormSqlite/gormSqliteErrors"
	"gorm.io/gorm"
)

func New[DomainModel, GormModel any](
	db *gorm.DB,
	dto func(DomainModel) GormModel,
) command.ICreateUpdateDelete[DomainModel] {
	return struct {
		create[DomainModel, GormModel]
		update[DomainModel, GormModel]
		delete[DomainModel, GormModel]
	}{
		create[DomainModel, GormModel]{db, dto},
		update[DomainModel, GormModel]{db, dto},
		delete[DomainModel, GormModel]{db, dto},
	}
}

type create[DomainModel, GormModel any] struct {
	db  *gorm.DB
	dto func(DomainModel) GormModel
}

func (c create[DomainModel, GormModel]) Create(ctx context.Context, req DomainModel) error {
	gormModel := c.dto(req)
	res := c.db.WithContext(ctx).Create(&gormModel)
	return gormSqliteErrors.HandleCommandResult(res)
}

type update[DomainModel, GormModel any] struct {
	db  *gorm.DB
	dto func(DomainModel) GormModel
}

func (u update[DomainModel, GormModel]) Update(ctx context.Context, req DomainModel) error {
	gormModel := u.dto(req)
	res := u.db.WithContext(ctx).Save(&gormModel)
	return gormSqliteErrors.HandleCommandResult(res)
}

type delete[DomainModel, GormModel any] struct {
	db  *gorm.DB
	dto func(DomainModel) GormModel
}

func (d delete[DomainModel, GormModel]) Delete(ctx context.Context, req DomainModel) error {
	gormModel := d.dto(req)
	res := d.db.WithContext(ctx).Delete(&gormModel)
	return gormSqliteErrors.HandleCommandResult(res)
}
