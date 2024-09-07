package bunSqliteDto

import (
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/command"
	"github.com/uptrace/bun"
)

type FileInfo struct {
	bun.BaseModel `bun:"table:file_infos"`

	ID        string         `bun:"id,pk,type:uuid"`
	Name      string         `bun:",unique,notnull"`
	Type      model.FileType `bun:",notnull"`
	Size      int            `bun:",notnull"`
	CreatedAt time.Time      `bun:",notnull"`
	UpdatedAt time.Time      `bun:",notnull"`
	DeletedAt time.Time      `bun:",notnull"`
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

func NewFileInfoUpdate(ids model.FileInfo, req command.FileInfoUpdateRequest) (FileInfo, []string) {
	fileInfo := FileInfo{
		ID: ids.ID,
	}
	var columns []string

	fileInfo.UpdatedAt = req.UpdatedAt
	columns = append(columns, "updated_at")
	if req.Name != nil {
		fileInfo.Name = *req.Name
		columns = append(columns, "name")
	}

	return fileInfo, columns
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

func NewFileInfos(fileInfos model.FileInfos) FileInfos {
	dtos := make(FileInfos, 0, len(fileInfos))
	for _, fileInfo := range fileInfos {
		dtos = append(dtos, NewFileInfo(fileInfo))
	}
	return dtos
}

func (dtos FileInfos) ToDomain() model.FileInfos {
	fileInfos := make(model.FileInfos, 0, len(dtos))
	for _, fileInfo := range dtos {
		fileInfos = append(fileInfos, fileInfo.ToDomain())
	}
	return fileInfos
}
