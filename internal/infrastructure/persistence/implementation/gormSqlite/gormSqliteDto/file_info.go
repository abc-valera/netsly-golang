package gormSqliteDto

import (
	"time"

	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
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

func NewDomainFileInfo(fileInfo FileInfo) model.FileInfo {
	return model.FileInfo{
		ID:        fileInfo.ID,
		Name:      fileInfo.Name,
		Type:      fileInfo.Type,
		Size:      fileInfo.Size,
		CreatedAt: fileInfo.CreatedAt,
		UpdatedAt: fileInfo.UpdatedAt,
		DeletedAt: fileInfo.DeletedAt,
	}
}

type FileInfos []FileInfo

func NewDomainFileInfos(fileInfos FileInfos) model.FileInfos {
	var domainFileInfos model.FileInfos
	for _, fileInfo := range fileInfos {
		domainFileInfos = append(domainFileInfos, NewDomainFileInfo(fileInfo))
	}
	return domainFileInfos
}
