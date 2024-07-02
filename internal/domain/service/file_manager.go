package service

import (
	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
)

var ErrFileNotFound = coderr.NewCodeMessage(coderr.CodeNotFound, "File not found")

type IFileManager interface {
	GetContent(fileName string) (model.FileContent, error)
	Save(fileName string, fileContent model.FileContent) error
	Update(fileName string, fileContent model.FileContent) error
	Remove(fileName string) error
}
