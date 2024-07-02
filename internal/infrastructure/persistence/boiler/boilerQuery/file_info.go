package boilerQuery

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/boiler/boilerDto"
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
	return boilerDto.NewDomainFileInfoWithErrHandle(sqlboiler.FindFileInfo(ctx, f.executor, id))
}

func (f fileInfo) GetAll(ctx context.Context) (model.FileInfos, error) {
	return boilerDto.NewDomainFileInfosWithErrHandle(sqlboiler.FileInfos().All(ctx, f.executor))
}
