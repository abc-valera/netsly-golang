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

func NewDomainComment(comment Comment) model.Comment {
	return model.Comment{
		ID:        comment.ID,
		Text:      comment.Text,
		CreatedAt: comment.CreatedAt,
	}
}

type Comments []Comment

func NewDomainComments(comments Comments) model.Comments {
	var domainComments model.Comments
	for _, comment := range comments {
		domainComments = append(domainComments, NewDomainComment(comment))
	}
	return domainComments
}
