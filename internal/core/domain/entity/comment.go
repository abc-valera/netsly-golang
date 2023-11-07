package entity

import (
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/codeerr"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/entity/common"
)

var (
	ErrCommentUserIDInvalid = codeerr.NewMsgErr(codeerr.CodeInvalidArgument, "Provided invalid user ID")
	ErrCommentJokeIDInvalid = codeerr.NewMsgErr(codeerr.CodeInvalidArgument, "Provided invalid joke ID")
	ErrCommentTextInvalid   = codeerr.NewMsgErr(codeerr.CodeInvalidArgument, "Provided invalid text")
)

type Comment struct {
	common.BaseEntity
	UserID string
	JokeID string
	Text   string
}

func NewComment(userID, jokeID, text string) (*Comment, error) {
	if userID == "" {
		return nil, ErrCommentUserIDInvalid
	}
	if jokeID == "" {
		return nil, ErrCommentJokeIDInvalid
	}
	if text == "" {
		return nil, ErrCommentTextInvalid
	}

	return &Comment{
		BaseEntity: common.NewBaseEntity(),
		UserID:     userID,
		JokeID:     jokeID,
		Text:       text,
	}, nil
}

type Comments []*Comment
