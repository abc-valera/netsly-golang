package entity

import (
	"time"

	"github.com/abc-valera/flugo-api-golang/internal/domain/codeerr"
	"github.com/abc-valera/flugo-api-golang/internal/domain/repository/spec"
	"github.com/google/uuid"
)

var (
	ErrJokesOrderByNotSupported = codeerr.NewMsgErr(codeerr.CodeInvalidArgument, "OrderBy is supported only for 'title' and 'created_at' field")
)

type Joke struct {
	ID          string
	UserID      string
	Title       string
	Text        string
	Explanation string
	CreatedAt   time.Time
}

func NewJoke(userID, title, text, explanation string) (*Joke, error) {
	if userID == "" {
		return nil, codeerr.NewMsgErr(codeerr.CodeInvalidArgument, "Provided invalid user ID")
	}
	if title == "" {
		return nil, codeerr.NewMsgErr(codeerr.CodeInvalidArgument, "Provided invalid title")
	}
	if text == "" {
		return nil, codeerr.NewMsgErr(codeerr.CodeInvalidArgument, "Provided invalid text")
	}

	return &Joke{
		ID:          uuid.NewString(),
		UserID:      userID,
		Title:       title,
		Text:        text,
		Explanation: explanation,
		CreatedAt:   time.Now(),
	}, nil
}

type Jokes []*Joke

func ValidateJokeSelectParams(params spec.SelectParams) error {
	if params.OrderBy != "" && params.OrderBy != "title" && params.OrderBy != "created_at" {
		return ErrJokesOrderByNotSupported
	}
	return nil
}
