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
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"
)

type IFile interface {
	Create(ctx context.Context, req FileCreateRequest) (model.FileInfo, error)
	Update(ctx context.Context, id string, req FileUpdateRequest) error
	Delete(ctx context.Context, id string) error

	GetByID(ctx context.Context, id string) (model.FileInfo, []byte, error)
}

type file struct {
	fileInfoCommand    command.IFileInfo
	fileInfoQuery      query.IFileInfo
	fileContentCommand command.IFileContent
	fileContentQuery   query.IFileContent
	commandTransactor  commandTransactor.ITransactor
}

func NewFile(
	fileInfoCommand command.IFileInfo,
	fileInfoQuery query.IFileInfo,
	fileContentCommand command.IFileContent,
	fileContentQuery query.IFileContent,
	commandTransactor commandTransactor.ITransactor,
) IFile {
	return file{
		fileInfoCommand:    fileInfoCommand,
		fileInfoQuery:      fileInfoQuery,
		fileContentCommand: fileContentCommand,
		fileContentQuery:   fileContentQuery,
		commandTransactor:  commandTransactor,
	}
}

type FileCreateRequest struct {
	Name string         `validate:"min=1,max=256"`
	Type model.FileType `validate:"enum"`
	Size int            `validate:"min=1,max=32000000"`

	FileContent []byte
}

func (e file) Create(ctx context.Context, req FileCreateRequest) (model.FileInfo, error) {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	if err := global.Validate().Struct(req); err != nil {
		return model.FileInfo{}, err
	}

	var returnFileInfo model.FileInfo
	txFunc := func(ctx context.Context, txCommands persistence.Commands, txQueries persistence.Queries) error {
		fileInfo, err := txCommands.FileInfo.Create(ctx, model.FileInfo{
			ID:        uuid.New().String(),
			Name:      req.Name,
			Type:      req.Type,
			Size:      req.Size,
			CreatedAt: time.Now(),
		})
		if err != nil {
			return err
		}
		returnFileInfo = fileInfo

		if _, err := txCommands.FileContent.Create(ctx, command.FileContentCreateRequest{
			Content: req.FileContent,
			ID:      fileInfo.ID,
		}); err != nil {
			return err
		}

		return nil
	}

	if err := e.commandTransactor.PerformTX(ctx, txFunc); err != nil {
		return model.FileInfo{}, err
	}

	return returnFileInfo, nil
}

type FileUpdateRequest struct {
	Name *string `validate:"omitempty,min=1,max=256"`

	FileContent *[]byte `validate:"omitempty,min=1,max=32000000"`
}

func (e file) Update(ctx context.Context, id string, req FileUpdateRequest) error {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

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
		if _, err := e.fileContentCommand.Update(ctx, id, command.FileContentUpdateRequest{
			Content: *req.FileContent,
		}); err != nil {
			return err
		}
	}

	return nil
}

func (e file) Delete(ctx context.Context, id string) error {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	txFunc := func(ctx context.Context, txCommands persistence.Commands, txQueries persistence.Queries) error {
		if err := txCommands.FileInfo.Delete(ctx, id); err != nil {
			return err
		}

		if err := e.fileContentCommand.Delete(ctx, id); err != nil {
			return err
		}

		return nil
	}

	return e.commandTransactor.PerformTX(ctx, txFunc)
}

func (e file) GetByID(ctx context.Context, id string) (model.FileInfo, []byte, error) {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	fileInfo, err := e.fileInfoQuery.GetByID(ctx, id)
	if err != nil {
		return model.FileInfo{}, []byte{}, err
	}

	fileContent, err := e.fileContentQuery.GetByID(ctx, id)
	if err != nil {
		return model.FileInfo{}, []byte{}, err
	}

	return fileInfo, fileContent, nil
}
