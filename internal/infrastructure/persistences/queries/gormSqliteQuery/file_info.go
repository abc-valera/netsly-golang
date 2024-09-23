package gormSqliteQuery

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/domain/global"
	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/selector"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/dependencies/gormSqlite/gormSqliteDto"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/dependencies/gormSqlite/gormSqliteErrors"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/queries/gormSqliteQuery/gormSelector"
	"gorm.io/gorm"
)

type fileInfo struct {
	db *gorm.DB
}

func NewFileInfo(db *gorm.DB) query.IFileInfo {
	return &fileInfo{
		db: db,
	}
}

func (q fileInfo) GetByID(ctx context.Context, id string) (model.FileInfo, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	var fileInfo gormSqliteDto.FileInfo
	res := q.db.Where("id = ?", id).First(&fileInfo)
	return fileInfo.ToDomain(), gormSqliteErrors.HandleQueryResult(res)
}

func (q fileInfo) GetAll(ctx context.Context, selector selector.Selector) (model.FileInfos, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	var fileInfos gormSqliteDto.FileInfos
	res := gormSelector.WithSelector(q.db, selector).WithContext(ctx).
		Find(&fileInfos)
	return fileInfos.ToDomain(), gormSqliteErrors.HandleQueryResult(res)
}
