package localFileManager

import (
	"os"

	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/service"
)

type localFileManager struct {
	filesPath string
}

// New creates a new instance of localFileManager, which stores files locally
// in the specified filesPath.
func New(filesPath string) service.IFileManager {
	return localFileManager{
		filesPath: filesPath,
	}
}

func (s localFileManager) GetContent(fileName string) (model.FileContent, error) {
	content, err := os.ReadFile(s.filesPath + "/" + fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, service.ErrFileNotFound
		}
		return nil, coderr.NewInternalErr(err)
	}

	return content, nil
}

func (s localFileManager) Save(fileName string, fileContent model.FileContent) error {
	newFile, err := os.Create(s.filesPath + "/" + fileName)
	if err != nil {
		return coderr.NewInternalErr(err)
	}
	defer newFile.Close()

	if _, err := newFile.Write(fileContent); err != nil {
		return coderr.NewInternalErr(err)
	}

	return nil
}

func (s localFileManager) Update(fileName string, fileContent model.FileContent) error {
	if err := s.Remove(fileName); err != nil {
		return err
	}

	return s.Save(fileName, fileContent)
}

func (s localFileManager) Remove(fileName string) error {
	if err := os.Remove(s.filesPath + "/" + fileName); err != nil {
		if os.IsNotExist(err) {
			return service.ErrFileNotFound
		}
		return coderr.NewInternalErr(err)
	}

	return nil
}
