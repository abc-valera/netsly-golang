package model

import (
	"time"

	"github.com/abc-valera/netsly-api-golang/internal/domain/coderr"
)

var (
	ErrLikeNotFound = coderr.NewMessage(coderr.CodeNotFound, "Like not found")
)

// Like represents a like entity.
// Technically, it's a many-to-many relationship between users and jokes.
type Like struct {
	UserID    string
	JokeID    string
	CreatedAt time.Time
}

type Likes []Like