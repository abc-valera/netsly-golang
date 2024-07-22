package boilerSqliteCommand

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/boilerSqlite/boilerSqliteDto"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/boilerSqlite/boilerSqliteErrutil"
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

func (c comment) Create(ctx context.Context, req command.CommentCreateRequest) (model.Comment, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	comment := sqlboiler.Comment{
		ID:        req.Comment.ID,
		Text:      req.Comment.Text,
		CreatedAt: req.Comment.CreatedAt,

		UserID: req.UserID,
		JokeID: req.JokeID,
	}
	err := comment.Insert(ctx, c.executor, boil.Infer())
	return boilerSqliteDto.NewDomainComment(&comment), boilerSqliteErrutil.HandleErr(err)
}

func (c comment) Update(ctx context.Context, commentID string, req command.CommentUpdateRequest) (model.Comment, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	comment, err := sqlboiler.FindComment(ctx, c.executor, commentID)
	if err != nil {
		return model.Comment{}, boilerSqliteErrutil.HandleErr(err)
	}
	if req.Text != nil {
		comment.Text = *req.Text
	}
	_, err = comment.Update(ctx, c.executor, boil.Infer())
	return boilerSqliteDto.NewDomainComment(comment), boilerSqliteErrutil.HandleErr(err)
}

func (c comment) Delete(ctx context.Context, id string) error {
	_, span := global.NewSpan(ctx)
	defer span.End()

	comment, err := sqlboiler.FindComment(ctx, c.executor, id)
	if err != nil {
		return boilerSqliteErrutil.HandleErr(err)
	}
	_, err = comment.Delete(ctx, c.executor)
	return boilerSqliteErrutil.HandleErr(err)
}
