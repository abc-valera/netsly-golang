package model

import (
	"time"

	"github.com/abc-valera/netsly-golang/internal/core/coderr"
)

var ErrLikeNotFound = coderr.NewCodeMessage(coderr.CodeNotFound, "Like not found")

type Like struct {
	CreatedAt time.Time
	DeletedAt time.Time
}

type Likes []Like
