package boilerSqliteQuery

import (
	"context"

	"github.com/abc-valera/netsly-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-golang/internal/core/global"
	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/selector"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/boilerSqlite/boilerSqliteDto"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/boilerSqlite/boilerSqliteErrutil"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type fileInfo struct {
	executor boil.ContextExecutor
}

func NewFileInfo(executor boil.ContextExecutor) query.IFileInfo {
	return &fileInfo{
		executor: executor,
	}
}

func (f fileInfo) GetByID(ctx context.Context, id string) (model.FileInfo, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	fileInfo, err := sqlboiler.FindFileInfo(ctx, f.executor, id)
	return boilerSqliteDto.NewDomainFileInfo(fileInfo), boilerSqliteErrutil.HandleErr(err)
}

func (f fileInfo) GetAll(ctx context.Context, selector selector.Selector) (model.FileInfos, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	fileInfos, err := sqlboiler.FileInfos().All(ctx, f.executor)
	return boilerSqliteDto.NewDomainFileInfos(fileInfos), boilerSqliteErrutil.HandleErr(err)
}
