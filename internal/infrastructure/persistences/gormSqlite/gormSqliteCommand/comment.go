package gormSqliteCommand

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/core/global"
	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/gormSqlite/gormSqliteDto"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/gormSqlite/gormSqliteErrutil"
	"gorm.io/gorm"
)

type comment struct {
	db *gorm.DB
}

func NewComment(db *gorm.DB) command.IComment {
	return &comment{
		db: db,
	}
}

func (c comment) Create(ctx context.Context, req command.CommentCreateRequest) (model.Comment, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	comment := gormSqliteDto.Comment{
		ID:        req.Comment.ID,
		Text:      req.Comment.Text,
		CreatedAt: req.Comment.CreatedAt,
		UserID:    req.UserID,
		JokeID:    req.JokeID,
	}
	res := c.db.Create(&comment)
	return gormSqliteDto.NewDomainComment(comment), gormSqliteErrutil.HandleCommandResult(res)
}

func (c comment) Update(ctx context.Context, id string, req command.CommentUpdateRequest) (model.Comment, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	var comment gormSqliteDto.Comment
	queryRes := c.db.Where("id = ?", id).First(&comment)
	if err := gormSqliteErrutil.HandleQueryResult(queryRes); err != nil {
		return model.Comment{}, err
	}

	if req.Text != nil {
		comment.Text = *req.Text
	}

	updateRes := c.db.Save(&comment)
	return gormSqliteDto.NewDomainComment(comment), gormSqliteErrutil.HandleCommandResult(updateRes)
}

func (c comment) Delete(ctx context.Context, id string) error {
	_, span := global.NewSpan(ctx)
	defer span.End()

	comment := gormSqliteDto.Comment{
		ID: id,
	}
	res := c.db.Delete(&comment)
	return gormSqliteErrutil.HandleCommandResult(res)
}
