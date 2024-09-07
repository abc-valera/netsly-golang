package gormSqliteDto

import (
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
)

type Like struct {
	CreatedAt time.Time `gorm:"not null"`
	DeletedAt time.Time `gorm:"not null"`
	UserID    string    `gorm:"primaryKey;not null"`
	JokeID    string    `gorm:"primaryKey;not null"`
}

func NewLike(like model.Like) Like {
	return Like{
		CreatedAt: like.CreatedAt,
		DeletedAt: like.DeletedAt,
		UserID:    like.UserID,
		JokeID:    like.JokeID,
	}
}

func (dto Like) ToDomain() model.Like {
	return model.Like{
		CreatedAt: dto.CreatedAt,
		DeletedAt: dto.DeletedAt,
		UserID:    dto.UserID,
		JokeID:    dto.JokeID,
	}
}

type Likes []Like

func NewLikes(domainLikes model.Likes) Likes {
	var likes Likes
	for _, domainLike := range domainLikes {
		likes = append(likes, NewLike(domainLike))
	}
	return likes
}

func (dtos Likes) ToDomain() model.Likes {
	var domainLikes model.Likes
	for _, dto := range dtos {
		domainLikes = append(domainLikes, dto.ToDomain())
	}
	return domainLikes
}
