package boilerCommand

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/boiler/boilerDto"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/boiler/errutil"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type fileInfo struct {
	executor boil.ContextExecutor
}

func NewFileInfo(executor boil.ContextExecutor) command.IFileInfo {
	return &fileInfo{
		executor: executor,
	}
}

func (f fileInfo) Create(ctx context.Context, req model.FileInfo) (model.FileInfo, error) {
	fileInfo := sqlboiler.FileInfo{
		ID:        req.ID,
		Name:      req.Name,
		Type:      int(req.Type),
		Size:      req.Size,
		CreatedAt: req.CreatedAt,
	}
	err := fileInfo.Insert(ctx, f.executor, boil.Infer())
	return boilerDto.NewDomainFileInfoWithErrHandle(&fileInfo, err)
}

func (f fileInfo) Update(ctx context.Context, id string, req command.FileInfoUpdateRequest) (model.FileInfo, error) {
	fileInfo, err := sqlboiler.FindFileInfo(ctx, f.executor, id)
	if err != nil {
		return model.FileInfo{}, errutil.HandleErr(err)
	}
	if req.Name != nil {
		fileInfo.Name = *req.Name
	}
	_, err = fileInfo.Update(ctx, f.executor, boil.Infer())
	return boilerDto.NewDomainFileInfoWithErrHandle(fileInfo, err)
}

func (f fileInfo) Delete(ctx context.Context, id string) error {
	fileInfo, err := sqlboiler.FindFileInfo(ctx, f.executor, id)
	if err != nil {
		return errutil.HandleErr(err)
	}
	_, err = fileInfo.Delete(ctx, f.executor)
	return errutil.HandleErr(err)
}
