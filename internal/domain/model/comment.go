package model

import (
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/util/coderr"
)

var ErrCommentNotFound = coderr.NewCodeMessage(coderr.CodeNotFound, "Comment not found")

type Comment struct {
	ID        string
	Text      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time

	UserID string
	JokeID string
}
