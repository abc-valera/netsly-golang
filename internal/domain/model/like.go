package model

import (
	"time"

	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
)

var ErrLikeNotFound = coderr.NewCodeMessage(coderr.CodeNotFound, "Like not found")

type Like struct {
	CreatedAt time.Time
	UserID    string
	JokeID    string
}

type Likes []Like
