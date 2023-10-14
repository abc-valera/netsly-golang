package entity

import (
	"time"

	"github.com/abc-valera/flugo-api-golang/internal/domain/codeerr"
	"github.com/abc-valera/flugo-api-golang/internal/domain/repository/spec"
	"github.com/google/uuid"
)

var (
	ErrCommentsOrderByNotSupported = codeerr.NewMsgErr(codeerr.CodeInvalidArgument, "OrderBy is supported only for 'created_at' field")
)

type Comment struct {
	ID        string
	UserID    string
	JokeID    string
	Text      string
	CreatedAt time.Time
}

func NewComment(userID, jokeID, text string) (*Comment, error) {
	if userID == "" {
		return nil, codeerr.NewMsgErr(codeerr.CodeInvalidArgument, "Provided invalid user ID")
	}
	if jokeID == "" {
		return nil, codeerr.NewMsgErr(codeerr.CodeInvalidArgument, "Provided invalid joke ID")
	}
	if text == "" {
		return nil, codeerr.NewMsgErr(codeerr.CodeInvalidArgument, "Provided invalid text")
	}

	return &Comment{
		ID:        uuid.NewString(),
		UserID:    userID,
		JokeID:    jokeID,
		Text:      text,
		CreatedAt: time.Now(),
	}, nil
}

type Comments []*Comment

func ValidateCommentSelectParams(params spec.SelectParams) error {
	if params.OrderBy != "" && params.OrderBy != "created_at" {
		return ErrCommentsOrderByNotSupported
	}
	return nil
}
