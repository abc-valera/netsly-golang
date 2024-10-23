package model

import (
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/util/coderr"
)

var ErrLikeNotFound = coderr.NewCodeMessage(coderr.CodeNotFound, "Like not found")

type Like struct {
	CreatedAt time.Time
	DeletedAt time.Time

	UserID string
	JokeID string
}
