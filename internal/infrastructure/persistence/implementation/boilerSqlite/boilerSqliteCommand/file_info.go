package boilerSqliteCommand

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/boilerSqlite/boilerSqliteDto"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/boilerSqlite/boilerSqliteErrutil"
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
	_, span := global.NewSpan(ctx)
	defer span.End()

	fileInfo := sqlboiler.FileInfo{
		ID:        req.ID,
		Name:      req.Name,
		Type:      int64(req.Type),
		Size:      int64(req.Size),
		CreatedAt: req.CreatedAt,
		UpdatedAt: req.UpdatedAt,
		DeletedAt: req.DeletedAt,
	}
	err := fileInfo.Insert(ctx, f.executor, boil.Infer())
	return boilerSqliteDto.NewDomainFileInfo(&fileInfo), boilerSqliteErrutil.HandleErr(err)
}

func (f fileInfo) Update(ctx context.Context, id string, req command.FileInfoUpdateRequest) (model.FileInfo, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	fileInfo, err := sqlboiler.FindFileInfo(ctx, f.executor, id)
	if err != nil {
		return model.FileInfo{}, boilerSqliteErrutil.HandleErr(err)
	}
	if req.Name != nil {
		fileInfo.Name = *req.Name
	}
	_, err = fileInfo.Update(ctx, f.executor, boil.Infer())
	return boilerSqliteDto.NewDomainFileInfo(fileInfo), boilerSqliteErrutil.HandleErr(err)
}

func (f fileInfo) Delete(ctx context.Context, id string) error {
	_, span := global.NewSpan(ctx)
	defer span.End()

	fileInfo, err := sqlboiler.FindFileInfo(ctx, f.executor, id)
	if err != nil {
		return boilerSqliteErrutil.HandleErr(err)
	}
	_, err = fileInfo.Delete(ctx, f.executor)
	return boilerSqliteErrutil.HandleErr(err)
}
