package entity

import (
	"github.com/abc-valera/flugo-api-golang/internal/domain/codeerr"
	"github.com/abc-valera/flugo-api-golang/internal/domain/entity/common"
	"github.com/abc-valera/flugo-api-golang/internal/domain/repository/spec"
)

var (
	ErrCommentsOrderByNotSupported = codeerr.NewMsgErr(codeerr.CodeInvalidArgument, "OrderBy is supported only for 'created_at' field")
)

type Comment struct {
	common.BaseEntity
	UserID string
	JokeID string
	Text   string
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
		BaseEntity: common.NewBaseEntity(),
		UserID:     userID,
		JokeID:     jokeID,
		Text:       text,
	}, nil
}

type Comments []*Comment

func ValidateCommentSelectParams(params spec.SelectParams) error {
	if params.OrderBy != "" && params.OrderBy != "created_at" {
		return ErrCommentsOrderByNotSupported
	}
	return nil
}
