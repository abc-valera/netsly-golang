package bunSqliteDto

import (
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
)

type FileInfo struct {
	ID        string         `bun:",pk,type:uuid"`
	Name      string         `bun:",unique,notnull"`
	Type      model.FileType `bun:",notnull"`
	Size      int            `bun:",notnull"`
	CreatedAt time.Time      `bun:",notnull"`
	UpdatedAt time.Time      `bun:",notnull"`
	DeletedAt time.Time      `bun:",notnull"`
}

func (f FileInfo) ToDomain() model.FileInfo {
	return model.FileInfo{
		ID:        f.ID,
		Name:      f.Name,
		Type:      f.Type,
		Size:      f.Size,
		CreatedAt: f.CreatedAt,
		UpdatedAt: f.UpdatedAt,
		DeletedAt: f.DeletedAt,
	}
}

type FileInfos []FileInfo

func (f FileInfos) ToDomain() model.FileInfos {
	fileInfos := make(model.FileInfos, 0, len(f))
	for _, fileInfo := range f {
		fileInfos = append(fileInfos, fileInfo.ToDomain())
	}
	return fileInfos
}
