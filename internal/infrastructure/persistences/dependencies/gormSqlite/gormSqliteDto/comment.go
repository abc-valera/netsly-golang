package gormSqliteDto

import (
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
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

func NewComments(domainComments []model.Comment) Comments {
	var comments Comments
	for _, domainComment := range domainComments {
		comments = append(comments, NewComment(domainComment))
	}
	return comments
}

func (dtos Comments) ToDomain() []model.Comment {
	var domainComments []model.Comment
	for _, dto := range dtos {
		domainComments = append(domainComments, dto.ToDomain())
	}
	return domainComments
}
