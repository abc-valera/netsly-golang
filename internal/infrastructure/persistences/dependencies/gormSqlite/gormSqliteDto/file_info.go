package gormSqliteDto

import (
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
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

type FileInfoJoke struct {
	FileInfoID string `gorm:"primaryKey;not null"`
	JokeID     string `gorm:"primaryKey;not null"`
}

type FileInfoRoom struct {
	FileInfoID string `gorm:"primaryKey;not null"`
	RoomID     string `gorm:"primaryKey;not null"`
}
