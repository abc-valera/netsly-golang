package gormSqliteCommand

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/core/global"
	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/gormSqlite/gormSqliteDto"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/gormSqlite/gormSqliteErrutil"
	"gorm.io/gorm"
)

type fileInfo struct {
	db *gorm.DB
}

func NewFileInfo(db *gorm.DB) command.IFileInfo {
	return &fileInfo{
		db: db,
	}
}

func (c fileInfo) Create(ctx context.Context, req model.FileInfo) (model.FileInfo, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	fileInfo := gormSqliteDto.FileInfo{
		ID:        req.ID,
		Name:      req.Name,
		Type:      req.Type,
		Size:      req.Size,
		CreatedAt: req.CreatedAt,
		UpdatedAt: req.UpdatedAt,
		DeletedAt: req.DeletedAt,
	}
	res := c.db.Create(&fileInfo)
	return gormSqliteDto.NewDomainFileInfo(fileInfo), gormSqliteErrutil.HandleCommandResult(res)
}

func (c fileInfo) Update(ctx context.Context, id string, req command.FileInfoUpdateRequest) (model.FileInfo, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	var fileInfo gormSqliteDto.FileInfo
	queryRes := c.db.Where("id = ?", id).First(&fileInfo)
	if err := gormSqliteErrutil.HandleQueryResult(queryRes); err != nil {
		return model.FileInfo{}, err
	}

	if req.Name != nil {
		fileInfo.Name = *req.Name
	}

	updateRes := c.db.Save(&fileInfo)
	return gormSqliteDto.NewDomainFileInfo(fileInfo), gormSqliteErrutil.HandleCommandResult(updateRes)
}

func (c fileInfo) Delete(ctx context.Context, id string) error {
	_, span := global.NewSpan(ctx)
	defer span.End()

	fileInfo := gormSqliteDto.FileInfo{
		ID: id,
	}
	res := c.db.Delete(&fileInfo)
	return gormSqliteErrutil.HandleCommandResult(res)
}
