package gormSqliteCommand

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/dependencies/gormSqlite/gormSqliteErrutil"
	"gorm.io/gorm"
)

type createUpdateDelete[BaseModel, UpdateModel any, GormModel domainable[BaseModel]] struct {
	create[BaseModel, GormModel]
	update[BaseModel, UpdateModel, GormModel]
	delete[BaseModel, GormModel]
}

func NewCreateUpdateDelete[BaseModel, UpdateModel any, GormModel domainable[BaseModel]](
	db *gorm.DB,
	dto func(BaseModel) GormModel,
	dtoUpdate func(GormModel, UpdateModel) GormModel,
) createUpdateDelete[BaseModel, UpdateModel, GormModel] {
	return createUpdateDelete[BaseModel, UpdateModel, GormModel]{
		create[BaseModel, GormModel]{db, dto},
		update[BaseModel, UpdateModel, GormModel]{db, dto, dtoUpdate},
		delete[BaseModel, GormModel]{db, dto},
	}
}

type createDelete[BaseModel any, GormModel domainable[BaseModel]] struct {
	create[BaseModel, GormModel]
	delete[BaseModel, GormModel]
}

func NewCreateDelete[BaseModel any, GormModel domainable[BaseModel]](
	db *gorm.DB,
	dto func(BaseModel) GormModel,
) createDelete[BaseModel, GormModel] {
	return createDelete[BaseModel, GormModel]{
		create[BaseModel, GormModel]{db, dto},
		delete[BaseModel, GormModel]{db, dto},
	}
}

type create[BaseModel any, GormModel domainable[BaseModel]] struct {
	db *gorm.DB

	dto func(BaseModel) GormModel
}

func (c create[BaseModel, GormModel]) Create(ctx context.Context, model BaseModel) (BaseModel, error) {
	gormModel := c.dto(model)
	res := c.db.WithContext(ctx).Create(&gormModel)
	return gormModel.ToDomain(), gormSqliteErrutil.HandleCommandResult(res)
}

type update[BaseModel any, UpdateModel any, GormModel domainable[BaseModel]] struct {
	db *gorm.DB

	dto       func(BaseModel) GormModel
	updateDto func(GormModel, UpdateModel) GormModel
}

func (u update[BaseModel, UpdateModel, GormModel]) Update(ctx context.Context, ids BaseModel, model UpdateModel) (BaseModel, error) {
	gormModel := u.dto(ids)
	if err := gormSqliteErrutil.HandleQueryResult(u.db.First(&gormModel)); err != nil {
		var emptyValue BaseModel
		return emptyValue, err
	}

	gormModel = u.updateDto(gormModel, model)
	res := u.db.WithContext(ctx).Save(&gormModel)
	return gormModel.ToDomain(), gormSqliteErrutil.HandleCommandResult(res)
}

type delete[BaseModel any, GormModel domainable[BaseModel]] struct {
	db *gorm.DB

	dto func(BaseModel) GormModel
}

func (d delete[BaseModel, GormModel]) Delete(ctx context.Context, model BaseModel) error {
	gormModel := d.dto(model)
	res := d.db.WithContext(ctx).Delete(&gormModel)
	return gormSqliteErrutil.HandleCommandResult(res)
}

type domainable[BaseModel any] interface {
	ToDomain() BaseModel
}
