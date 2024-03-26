package sqlboilerCommand

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/model"
	"github.com/abc-valera/netsly-api-golang/pkg/persistence/sqlboilerImpl/dto"
	"github.com/abc-valera/netsly-api-golang/pkg/persistence/sqlboilerImpl/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type comment struct {
	executor boil.ContextExecutor
}

func newComment(executor boil.ContextExecutor) command.IComment {
	return &comment{
		executor: executor,
	}
}

func (c comment) Create(ctx context.Context, req model.Comment) (model.Comment, error) {
	comment := sqlboiler.Comment{
		ID:        req.ID,
		Text:      req.Text,
		CreatedAt: req.CreatedAt,
		UserID:    req.UserID,
		JokeID:    req.JokeID,
	}
	err := comment.Insert(ctx, c.executor, boil.Infer())
	return dto.ToDomainCommentWithErrHandle(&comment, err)
}

func (c comment) Update(ctx context.Context, commentID string, req command.CommentUpdate) (model.Comment, error) {
	comment, err := sqlboiler.FindComment(ctx, c.executor, commentID)
	if err != nil {
		return model.Comment{}, errors.HandleErr(err)
	}
	if req.Text != nil {
		comment.Text = *req.Text
	}
	_, err = comment.Update(ctx, c.executor, boil.Infer())
	return dto.ToDomainCommentWithErrHandle(comment, err)
}

func (c comment) Delete(ctx context.Context, id string) error {
	comment, err := sqlboiler.FindComment(ctx, c.executor, id)
	if err != nil {
		return errors.HandleErr(err)
	}
	_, err = comment.Delete(ctx, c.executor)
	return errors.HandleErr(err)
}
