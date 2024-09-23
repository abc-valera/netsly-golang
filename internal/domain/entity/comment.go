package entity

import (
	"context"
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/global"
	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"
)

type IComment interface {
	Create(ctx context.Context, req CommentCreateRequest) (model.Comment, error)
	Update(ctx context.Context, commentID string, req CommentUpdateRequest) (model.Comment, error)
	Delete(ctx context.Context, commentID string) error

	query.IComment
}

type comment struct {
	IDependency

	query.IComment
}

func newComment(dep IDependency) IComment {
	return comment{
		IDependency: dep,

		IComment: dep.Q().Comment,
	}
}

type CommentCreateRequest struct {
	Text string `validate:"required,min=4,max=256"`

	UserID string `validate:"required,uuid"`
	JokeID string `validate:"required,uuid"`
}

func (e comment) Create(ctx context.Context, req CommentCreateRequest) (model.Comment, error) {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	if err := global.Validate().Struct(req); err != nil {
		return model.Comment{}, err
	}

	comment := model.Comment{
		ID:        uuid.New().String(),
		Text:      req.Text,
		CreatedAt: time.Now().Truncate(time.Millisecond).Local(),
		UserID:    req.UserID,
		JokeID:    req.JokeID,
	}

	if err := e.C().Comment.Create(ctx, comment); err != nil {
		return model.Comment{}, err
	}

	return comment, nil
}

type CommentUpdateRequest struct {
	Text *string `validate:"min=4,max=256"`
}

func (e comment) Update(ctx context.Context, commentID string, req CommentUpdateRequest) (model.Comment, error) {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	if err := global.Validate().Struct(req); err != nil {
		return model.Comment{}, err
	}

	comment, err := e.Q().Comment.GetByID(ctx, commentID)
	if err != nil {
		return model.Comment{}, err
	}

	comment.UpdatedAt = time.Now().Truncate(time.Millisecond).Local()

	if req.Text != nil {
		comment.Text = *req.Text
	}

	if err := e.C().Comment.Update(ctx, comment); err != nil {
		return model.Comment{}, err
	}

	return comment, nil
}

func (e comment) Delete(ctx context.Context, commentID string) error {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	return e.C().Comment.Delete(ctx, model.Comment{ID: commentID})
}
