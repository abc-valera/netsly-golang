package model

import (
	"time"

	"github.com/abc-valera/netsly-api-golang/internal/domain/coderr"
)

var (
	ErrCommentNotFound = coderr.NewCodeMessage(coderr.CodeNotFound, "Comment not found")
)

type Comment struct {
	ID        string
	Text      string
	CreatedAt time.Time

	UserID string
	JokeID string
}

type Comments []Comment
