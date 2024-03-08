package model

import (
	"github.com/abc-valera/netsly-api-golang/internal/domain/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model/common"
)

var (
	ErrCommentNotFound = coderr.NewCodeMessage(coderr.CodeNotFound, "Comment not found")
)

type Comment struct {
	common.BaseEntity
	Text string

	UserID string
	JokeID string
}

type Comments []Comment
