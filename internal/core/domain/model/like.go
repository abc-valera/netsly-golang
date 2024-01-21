package model

import (
	"time"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/coderr"
)

var (
	ErrLikeNotFound = coderr.NewMessage(coderr.CodeNotFound, "Like not found")

	ErrLikeAlreadyExists = coderr.NewMessage(coderr.CodeAlreadyExists, "Like already exists")
)

// Like represents a like entity.
// Technically, it's a many-to-many relationship between users and jokes.
type Like struct {
	UserID    string
	JokeID    string
	CreatedAt time.Time
}

type Likes []Like
