package gormSqliteDto

import (
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/command"
)

type Comment struct {
	ID        string    `gorm:"primaryKey;not null"`
	Text      string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
	DeletedAt time.Time `gorm:"not null"`

	UserID string `gorm:"not null"`
	JokeID string `gorm:"not null"`
}

func NewComment(comment model.Comment) Comment {
	return Comment{
		ID:        comment.ID,
		Text:      comment.Text,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
		DeletedAt: comment.DeletedAt,
		UserID:    comment.UserID,
		JokeID:    comment.JokeID,
	}
}

func NewCommentUpdate(comment Comment, req command.CommentUpdateRequest) Comment {
	comment.UpdatedAt = req.UpdatedAt

	if req.Text != nil {
		comment.Text = *req.Text
	}

	return comment
}

func (dto Comment) ToDomain() model.Comment {
	return model.Comment{
		ID:        dto.ID,
		Text:      dto.Text,
		CreatedAt: dto.CreatedAt,
		UpdatedAt: dto.UpdatedAt,
		DeletedAt: dto.DeletedAt,
		UserID:    dto.UserID,
		JokeID:    dto.JokeID,
	}
}

type Comments []Comment

func NewComments(domainComments model.Comments) Comments {
	var comments Comments
	for _, domainComment := range domainComments {
		comments = append(comments, NewComment(domainComment))
	}
	return comments
}

func (dtos Comments) ToDomain() model.Comments {
	var domainComments model.Comments
	for _, dto := range dtos {
		domainComments = append(domainComments, dto.ToDomain())
	}
	return domainComments
}
