package bunSqliteCommand

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/bunSqlite/bunSqliteDto"
	bunSqlitErrutil "github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/bunSqlite/bunSqliteErrutil"
	"github.com/uptrace/bun"
)

type comment struct {
	db bun.IDB
}

func NewComment(db bun.IDB) command.IComment {
	return &comment{
		db: db,
	}
}

func (c comment) Create(ctx context.Context, req command.CommentCreateRequest) (model.Comment, error) {
	comment := bunSqliteDto.Comment{
		ID:        req.Comment.ID,
		Text:      req.Comment.Text,
		CreatedAt: req.Comment.CreatedAt,
		UpdatedAt: req.Comment.UpdatedAt,
		DeletedAt: req.Comment.DeletedAt,

		UserID: req.UserID,
		JokeID: req.JokeID,
	}

	res, err := c.db.NewInsert().Model(&comment).Exec(ctx)
	return comment.ToDomain(), bunSqlitErrutil.HandleCommandResult(res, err)
}

func (c comment) Update(ctx context.Context, id string, req command.CommentUpdateRequest) (model.Comment, error) {
	comment := bunSqliteDto.Comment{
		ID: id,
	}
	var columns []string

	if req.Text != nil {
		comment.Text = *req.Text
		columns = append(columns, "text")
	}

	if len(columns) == 0 {
		return model.Comment{}, nil
	}

	res, err := c.db.NewUpdate().Model(&comment).Column(columns...).WherePK().Exec(ctx)
	return comment.ToDomain(), bunSqlitErrutil.HandleCommandResult(res, err)
}

func (c comment) Delete(ctx context.Context, id string) error {
	comment := bunSqliteDto.Comment{
		ID: id,
	}
	res, err := c.db.NewDelete().Model(&comment).WherePK().Exec(ctx)
	return bunSqlitErrutil.HandleCommandResult(res, err)
}
