package model

import (
	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/core/persistence/model/common"
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