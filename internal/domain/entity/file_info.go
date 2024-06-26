package entity

import (
	"context"
	"time"

	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/commandTransactor"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-api-golang/internal/domain/service"
	"github.com/google/uuid"
)

type IFile interface {
	Create(ctx context.Context, req FileCreateRequest) (model.FileInfo, error)
	Update(ctx context.Context, id string, req FileUpdateRequest) error
	Delete(ctx context.Context, id string) error

	GetByID(ctx context.Context, id string) (model.FileInfo, model.FileContent, error)
}

type file struct {
	commandTransactor commandTransactor.ITransactor
	fileInfoCommand   command.IFileInfo
	fileInfoQuery     query.IFileInfo
	fileManger        service.IFileManager
}

func NewFile(
	commandTransactor commandTransactor.ITransactor,
	fileInfoCommand command.IFileInfo,
	fileInfoQuery query.IFileInfo,
	fileManger service.IFileManager,
) IFile {
	return file{
		commandTransactor: commandTransactor,
		fileInfoCommand:   fileInfoCommand,
		fileInfoQuery:     fileInfoQuery,
		fileManger:        fileManger,
	}
}

type FileCreateRequest struct {
	Name string         `validate:"min=1,max=256"`
	Type model.FileType `validate:"enum"`
	Size int            `validate:"min=1,max=32000000"`

	FileContent model.FileContent
}

func (f file) Create(ctx context.Context, req FileCreateRequest) (model.FileInfo, error) {
	if err := global.Validate().Struct(req); err != nil {
		return model.FileInfo{}, err
	}

	var returnFileInfo model.FileInfo
	txFunc := func(ctx context.Context, txCommands persistence.Commands) error {
		guid := uuid.New().String()
		fileInfo, err := txCommands.FileInfo.Create(ctx, model.FileInfo{
			ID:        guid,
			Name:      req.Name,
			Type:      req.Type,
			Size:      req.Size,
			CreatedAt: time.Now(),
		})
		if err != nil {
			return err
		}
		returnFileInfo = fileInfo

		if err := f.fileManger.Save(guid, req.FileContent); err != nil {
			return err
		}

		return nil
	}

	if err := f.commandTransactor.PerformTX(ctx, txFunc); err != nil {
		return model.FileInfo{}, err
	}

	return returnFileInfo, nil
}

type FileUpdateRequest struct {
	Name *string `validate:"omitempty,min=1,max=256"`

	FileContent *model.FileContent `validate:"omitempty,min=1,max=32000000"`
}

func (e file) Update(ctx context.Context, id string, req FileUpdateRequest) error {
	if err := global.Validate().Struct(req); err != nil {
		return err
	}

	if req.Name != nil {
		if _, err := e.fileInfoCommand.Update(ctx, id, command.FileInfoUpdateRequest{
			Name: req.Name,
		}); err != nil {
			return err
		}
	}

	if req.FileContent != nil {
		if err := e.fileManger.Update(id, *req.FileContent); err != nil {
			return err
		}
	}

	return nil
}

func (e file) Delete(ctx context.Context, id string) error {
	txFunc := func(ctx context.Context, txCommands persistence.Commands) error {
		if err := txCommands.FileInfo.Delete(ctx, id); err != nil {
			return err
		}

		if err := e.fileManger.Remove(id); err != nil {
			return err
		}

		return nil
	}

	return e.commandTransactor.PerformTX(ctx, txFunc)
}

func (e file) GetByID(ctx context.Context, id string) (model.FileInfo, model.FileContent, error) {
	fileInfo, err := e.fileInfoQuery.GetByID(ctx, id)
	if err != nil {
		return model.FileInfo{}, model.FileContent{}, err
	}

	fileContent, err := e.fileManger.GetContent(id)
	if err != nil {
		return model.FileInfo{}, model.FileContent{}, err
	}

	return fileInfo, fileContent, nil
}
