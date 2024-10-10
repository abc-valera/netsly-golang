package bunSqliteQuery

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/selector"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/dependencies/bunSqlite/bunSqliteDto"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/dependencies/bunSqlite/bunSqliteErrors"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/queries/bunSqliteQuery/bunSqliteSelector"
	"github.com/uptrace/bun"
)

type fileInfo struct {
	db bun.IDB
}

func NewFileInfo(db bun.IDB) query.IFileInfo {
	return &fileInfo{
		db: db,
	}
}

func (q fileInfo) GetByID(ctx context.Context, id string) (model.FileInfo, error) {
	fileInfo := bunSqliteDto.FileInfo{}
	err := q.db.NewSelect().Model(&fileInfo).Where("id = ?", id).Scan(ctx)
	return fileInfo.ToDomain(), bunSqliteErrors.HandleQueryResult(err)
}

func (q fileInfo) GetAll(ctx context.Context, s selector.Selector) (model.FileInfos, error) {
	fileInfos := bunSqliteDto.FileInfos{}
	err := bunSqliteSelector.NewSelect(q.db, s).Model(&fileInfos).Scan(ctx)
	return fileInfos.ToDomain(), bunSqliteErrors.HandleQueryResult(err)
}
