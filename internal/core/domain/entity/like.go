package entity

import (
	"time"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/codeerr"
)

// Like represents a like entity.
// Technically, it's a many-to-many relationship between users and jokes.
type Like struct {
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
		UserID:    userID,
		JokeID:    jokeID,
		CreatedAt: time.Now(),
	}, nil
}

type Likes []*Like
