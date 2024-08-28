package bunSqliteCommand

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/bunSqlite/bunSqliteDto"
	bunSqlitErrutil "github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/bunSqlite/bunSqliteErrutil"
	"github.com/uptrace/bun"
)

type fileInfo struct {
	db bun.IDB
}

func NewFileInfo(db bun.IDB) command.IFileInfo {
	return &fileInfo{
		db: db,
	}
}

func (c fileInfo) Create(ctx context.Context, req model.FileInfo) (model.FileInfo, error) {
	fileInfo := bunSqliteDto.FileInfo{
		ID:        req.ID,
		Name:      req.Name,
		Type:      req.Type,
		Size:      req.Size,
		CreatedAt: req.CreatedAt,
		UpdatedAt: req.UpdatedAt,
		DeletedAt: req.DeletedAt,
	}

	res, err := c.db.NewInsert().Model(&fileInfo).Exec(ctx)
	return fileInfo.ToDomain(), bunSqlitErrutil.HandleCommandResult(res, err)
}

func (c fileInfo) Update(ctx context.Context, id string, req command.FileInfoUpdateRequest) (model.FileInfo, error) {
	fileInfo := bunSqliteDto.FileInfo{
		ID: id,
	}
	var columns []string

	if req.Name != nil {
		fileInfo.Name = *req.Name
		columns = append(columns, "name")
	}

	if len(columns) == 0 {
		return model.FileInfo{}, nil
	}

	res, err := c.db.NewUpdate().Model(&fileInfo).Column(columns...).WherePK().Exec(ctx)
	return fileInfo.ToDomain(), bunSqlitErrutil.HandleCommandResult(res, err)
}

func (c fileInfo) Delete(ctx context.Context, id string) error {
	fileInfo := bunSqliteDto.FileInfo{
		ID: id,
	}
	res, err := c.db.NewDelete().Model(&fileInfo).WherePK().Exec(ctx)
	return bunSqlitErrutil.HandleCommandResult(res, err)
}
