package model

import (
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/coderr"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/model/common"
)

var (
	ErrCommentNotFound = coderr.NewMessage(coderr.CodeNotFound, "Comment not found")
)

type Comment struct {
	common.BaseModel
	UserID string
	JokeID string
	Text   string
}

type Comments []Comment
