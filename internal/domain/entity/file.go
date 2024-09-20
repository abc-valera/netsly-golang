package entity

import (
	"context"
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/global"
	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-golang/internal/domain/util/coderr"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"
)

type IFile interface {
	Create(ctx context.Context, req FileCreateRequest) (FileCreateResponse, error)
	Update(ctx context.Context, id string, req FileUpdateRequest) error
	Delete(ctx context.Context, id string) error

	GetByID(ctx context.Context, id string) (model.FileInfo, model.FileContent, error)
}

type file struct {
	IDependency
}

func newFile(dep IDependency) IFile {
	return file{
		IDependency: dep,
	}
}

type FileCreateRequest struct {
	Name string         `validate:"min=1,max=256"`
	Type model.FileType `validate:"enum"`

	FileContent []byte
}

type FileCreateResponse struct {
	FileInfo    model.FileInfo
	FileContent model.FileContent
}

func (e file) Create(ctx context.Context, req FileCreateRequest) (FileCreateResponse, error) {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	if err := global.Validate().Struct(req); err != nil {
		return FileCreateResponse{}, err
	}

	size := len(req.FileContent)
	if size > 32000000 {
		return FileCreateResponse{}, coderr.NewCodeMessage(coderr.CodeInvalidArgument, "File content size is too large")
	}

	var returnFileInfo model.FileInfo
	txFunc := func(
		ctx context.Context,
		txC command.Commands,
		txQ query.Queries,
		txE Entities,
	) error {
		fileInfo, err := txC.FileInfo.Create(ctx, model.FileInfo{
			ID:        uuid.New().String(),
			Name:      req.Name,
			Type:      req.Type,
			Size:      size,
			CreatedAt: time.Now().Truncate(time.Millisecond),
		})
		if err != nil {
			return err
		}
		returnFileInfo = fileInfo

		if _, err := txC.FileContent.Create(ctx, model.FileContent{
			ID:      fileInfo.ID,
			Content: req.FileContent,
		}); err != nil {
			return err
		}

		return nil
	}

	if err := e.RunInTX(ctx, txFunc); err != nil {
		return FileCreateResponse{}, err
	}

	return FileCreateResponse{
		FileInfo:    returnFileInfo,
		FileContent: model.FileContent{ID: returnFileInfo.ID, Content: req.FileContent},
	}, nil
}

type FileUpdateRequest struct {
	Name *string `validate:"omitempty,min=1,max=256"`
}

func (e file) Update(ctx context.Context, id string, req FileUpdateRequest) error {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	if err := global.Validate().Struct(req); err != nil {
		return err
	}

	if req.Name != nil {
		if _, err := e.C().FileInfo.Update(
			ctx,
			model.FileInfo{ID: id},
			command.FileInfoUpdateRequest{
				UpdatedAt: time.Now().Truncate(time.Millisecond),

				Name: req.Name,
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

	txFunc := func(
		ctx context.Context,
		txC command.Commands,
		txQ query.Queries,
		txE Entities,
	) error {
		if err := txC.FileInfo.Delete(ctx, model.FileInfo{ID: id}); err != nil {
			return err
		}

		if err := txC.FileContent.Delete(ctx, model.FileContent{ID: id}); err != nil {
			return err
		}

		return nil
	}

	return e.RunInTX(ctx, txFunc)
}

func (e file) GetByID(ctx context.Context, id string) (model.FileInfo, model.FileContent, error) {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	fileInfo, err := e.Q().FileInfo.GetByID(ctx, id)
	if err != nil {
		return model.FileInfo{}, model.FileContent{}, err
	}

	fileContent, err := e.Q().FileContent.GetByID(ctx, id)
	if err != nil {
		return model.FileInfo{}, model.FileContent{}, err
	}

	return fileInfo, fileContent, nil
}
