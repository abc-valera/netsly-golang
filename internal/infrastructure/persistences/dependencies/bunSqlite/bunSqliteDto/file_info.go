package bunSqliteDto

import (
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/uptrace/bun"
)

type FileInfo struct {
	bun.BaseModel

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

type FileInfosJoke struct {
	bun.BaseModel `bun:"table:file_info_joke"`

	FileInfoID string `bun:"file_info_id,pk,type:uuid"`
	JokeID     string `bun:"joke_id,pk,type:uuid"`
}

type FileInfosRoom struct {
	bun.BaseModel `bun:"table:file_info_room"`

	FileInfoID string `bun:"file_info_id,pk,type:uuid"`
	RoomID     string `bun:"room_id,pk,type:uuid"`
}
