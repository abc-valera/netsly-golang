package model

import (
	"time"

	"github.com/abc-valera/netsly-golang/internal/core/coderr"
)

var ErrCommentNotFound = coderr.NewCodeMessage(coderr.CodeNotFound, "Comment not found")

type Comment struct {
	ID        string
	Text      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Comments []Comment
