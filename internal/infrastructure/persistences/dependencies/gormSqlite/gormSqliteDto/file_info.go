package gormSqliteDto

import (
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/command"
)

type FileInfo struct {
	ID        string         `gorm:"primaryKey;not null"`
	Name      string         `gorm:"not null"`
	Type      model.FileType `gorm:"not null"`
	Size      int            `gorm:"not null"`
	CreatedAt time.Time      `gorm:"not null"`
	UpdatedAt time.Time      `gorm:"not null"`
	DeletedAt time.Time      `gorm:"not null"`
}

func NewFileInfo(fileInfo model.FileInfo) FileInfo {
	return FileInfo{
		ID:        fileInfo.ID,
		Name:      fileInfo.Name,
		Type:      fileInfo.Type,
		Size:      fileInfo.Size,
		CreatedAt: fileInfo.CreatedAt,
		UpdatedAt: fileInfo.UpdatedAt,
		DeletedAt: fileInfo.DeletedAt,
	}
}

func NewFileInfoUpdate(fileInfo FileInfo, req command.FileInfoUpdateRequest) FileInfo {
	fileInfo.UpdatedAt = req.UpdatedAt

	if req.Name != nil {
		fileInfo.Name = *req.Name
	}

	return fileInfo
}

func (dto FileInfo) ToDomain() model.FileInfo {
	return model.FileInfo{
		ID:        dto.ID,
		Name:      dto.Name,
		Type:      dto.Type,
		Size:      dto.Size,
		CreatedAt: dto.CreatedAt,
		UpdatedAt: dto.UpdatedAt,
		DeletedAt: dto.DeletedAt,
	}
}

type FileInfos []FileInfo

func NewFileInfos(domainFileInfos model.FileInfos) FileInfos {
	var fileInfos FileInfos
	for _, domainFileInfo := range domainFileInfos {
		fileInfos = append(fileInfos, NewFileInfo(domainFileInfo))
	}
	return fileInfos
}

func (dtos FileInfos) ToDomain() model.FileInfos {
	var domainFileInfos model.FileInfos
	for _, dto := range dtos {
		domainFileInfos = append(domainFileInfos, dto.ToDomain())
	}
	return domainFileInfos
}
