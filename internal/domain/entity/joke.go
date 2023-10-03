package entity

import (
	"time"

	"github.com/abc-valera/flugo-api-golang/internal/domain/codeerr"
	"github.com/google/uuid"
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
		return nil, codeerr.NewErrWithMsg(codeerr.CodeInvalidArgument, "Provided invalid user ID")
	}
	if title == "" {
		return nil, codeerr.NewErrWithMsg(codeerr.CodeInvalidArgument, "Provided invalid title")
	}
	if text == "" {
		return nil, codeerr.NewErrWithMsg(codeerr.CodeInvalidArgument, "Provided invalid text")
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
