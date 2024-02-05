package model

import (
	"github.com/abc-valera/netsly-api-golang/internal/domain/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model/common"
)

var (
	ErrCommentNotFound = coderr.NewMessage(coderr.CodeNotFound, "Comment not found")
)

type Comment struct {
	common.BaseEntity
	UserID string
	JokeID string
	Text   string
}

type Comments []Comment
