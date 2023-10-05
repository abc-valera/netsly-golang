package entity

import (
	"time"

	"github.com/abc-valera/flugo-api-golang/internal/domain/codeerr"
	"github.com/google/uuid"
)

type Like struct {
	ID        string
	UserID    string
	JokeID    string
	CreatedAt time.Time
}

func NewLike(userID, jokeID string) (*Like, error) {
	if userID == "" {
		return nil, codeerr.NewMsgErr(codeerr.CodeInvalidArgument, "Provided invalid user ID")
	}
	if jokeID == "" {
		return nil, codeerr.NewMsgErr(codeerr.CodeInvalidArgument, "Provided invalid joke ID")
	}

	return &Like{
		ID:        uuid.NewString(),
		UserID:    userID,
		JokeID:    jokeID,
		CreatedAt: time.Now(),
	}, nil
}

type Likes []*Like
