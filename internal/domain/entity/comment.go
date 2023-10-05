package entity

import (
	"time"

	"github.com/abc-valera/flugo-api-golang/internal/domain/codeerr"
	"github.com/google/uuid"
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
