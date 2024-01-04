package model

import (
	"time"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/codeerr"
)

var (
	ErrLikeNotFound = codeerr.NewMessage(codeerr.CodeNotFound, "Like not found")

	ErrLikeAlreadyExists = codeerr.NewMessage(codeerr.CodeAlreadyExists, "Like already exists")
)

// Like represents a like entity.
// Technically, it's a many-to-many relationship between users and jokes.
type Like struct {
	UserID    string
	JokeID    string
	CreatedAt time.Time
}

type Likes []*Like
