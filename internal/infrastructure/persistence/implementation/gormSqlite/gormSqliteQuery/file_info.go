package gormSqliteQuery

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query/selector"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/gormSqlite/gormSqliteDto"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/gormSqlite/gormSqliteErrutil"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/gormSqlite/gormSqliteQuery/gormSelector"
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
	return gormSqliteDto.NewDomainFileInfo(fileInfo), gormSqliteErrutil.HandleQueryResult(res)
}

func (q fileInfo) GetAll(ctx context.Context, selector selector.Selector) (model.FileInfos, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	var fileInfos gormSqliteDto.FileInfos
	res := gormSelector.WithSelector(q.db, selector).WithContext(ctx).
		Find(&fileInfos)
	return gormSqliteDto.NewDomainFileInfos(fileInfos), gormSqliteErrutil.HandleQueryResult(res)
}
