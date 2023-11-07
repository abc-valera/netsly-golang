package entity

import (
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/codeerr"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/entity/common"
)

var (
	ErrJokeUserIDInvalid = codeerr.NewMsgErr(codeerr.CodeInvalidArgument, "Provided invalid user ID")
	ErrJokeTitleInvalid  = codeerr.NewMsgErr(codeerr.CodeInvalidArgument, "Provided invalid title")
	ErrJokeTextInvalid   = codeerr.NewMsgErr(codeerr.CodeInvalidArgument, "Provided invalid text")
)

type Joke struct {
	common.BaseEntity
	UserID      string
	Title       string
	Text        string
	Explanation string
}

func NewJoke(userID, title, text, explanation string) (*Joke, error) {
	if userID == "" {
		return nil, ErrJokeUserIDInvalid
	}
	if title == "" {
		return nil, ErrJokeTitleInvalid
	}
	if text == "" {
		return nil, ErrJokeTextInvalid
	}

	return &Joke{
		BaseEntity:  common.NewBaseEntity(),
		UserID:      userID,
		Title:       title,
		Text:        text,
		Explanation: explanation,
	}, nil
}

type Jokes []*Joke
