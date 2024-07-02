package boilerCommand

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/boiler/boilerDto"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/boiler/errutil"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type comment struct {
	executor boil.ContextExecutor
}

func NewComment(executor boil.ContextExecutor) command.IComment {
	return &comment{
		executor: executor,
	}
}

func (c comment) Create(ctx context.Context, userID, jokeID string, req model.Comment) (model.Comment, error) {
	comment := sqlboiler.Comment{
		ID:        req.ID,
		Text:      req.Text,
		CreatedAt: req.CreatedAt,

		UserID: userID,
		JokeID: jokeID,
	}
	err := comment.Insert(ctx, c.executor, boil.Infer())
	return boilerDto.NewDomainCommentWithErrHandle(&comment, err)
}

func (c comment) Update(ctx context.Context, commentID string, req command.CommentUpdate) (model.Comment, error) {
	comment, err := sqlboiler.FindComment(ctx, c.executor, commentID)
	if err != nil {
		return model.Comment{}, errutil.HandleErr(err)
	}
	if req.Text != nil {
		comment.Text = *req.Text
	}
	_, err = comment.Update(ctx, c.executor, boil.Infer())
	return boilerDto.NewDomainCommentWithErrHandle(comment, err)
}

func (c comment) Delete(ctx context.Context, id string) error {
	comment, err := sqlboiler.FindComment(ctx, c.executor, id)
	if err != nil {
		return errutil.HandleErr(err)
	}
	_, err = comment.Delete(ctx, c.executor)
	return errutil.HandleErr(err)
}
