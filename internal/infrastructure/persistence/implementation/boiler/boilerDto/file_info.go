package boilerDto

import (
	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
)

func NewDomainFileInfo(fileInfo *sqlboiler.FileInfo) model.FileInfo {
	if fileInfo == nil {
		return model.FileInfo{}
	}

	return model.FileInfo{
		ID:        fileInfo.ID,
		Name:      fileInfo.Name,
		Type:      model.FileType(fileInfo.Type),
		Size:      fileInfo.Size,
		CreatedAt: fileInfo.CreatedAt,
	}
}

func NewDomainFileInfoWithErrHandle(fileInfo *sqlboiler.FileInfo, err error) (model.FileInfo, error) {
	return NewDomainFileInfo(fileInfo), err
}

func NewDomainFileInfos(fileInfos sqlboiler.FileInfoSlice) model.FileInfos {
	var domainFileInfos model.FileInfos
	for _, fileInfo := range fileInfos {
		domainFileInfos = append(domainFileInfos, NewDomainFileInfo(fileInfo))
	}
	return domainFileInfos
}

func NewDomainFileInfosWithErrHandle(fileInfos sqlboiler.FileInfoSlice, err error) (model.FileInfos, error) {
	return NewDomainFileInfos(fileInfos), err
}
