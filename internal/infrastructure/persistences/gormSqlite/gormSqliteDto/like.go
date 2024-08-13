package gormSqliteDto

import (
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
)

type Like struct {
	CreatedAt time.Time `gorm:"not null"`
	DeletedAt time.Time `gorm:"not null"`

	UserID string `gorm:"primaryKey;not null"`
	JokeID string `gorm:"primaryKey;not null"`
}

func NewDomainLike(like Like) model.Like {
	return model.Like{
		CreatedAt: like.CreatedAt,
		DeletedAt: like.DeletedAt,
	}
}

type Likes []Like

func NewDomainLikes(likes Likes) model.Likes {
	var domainLikes model.Likes
	for _, like := range likes {
		domainLikes = append(domainLikes, NewDomainLike(like))
	}
	return domainLikes
}
