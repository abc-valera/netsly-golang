package model

import (
	"time"

	"github.com/abc-valera/netsly-api-golang/pkg/core/coderr"
)

var (
	ErrJokeNotFound = coderr.NewCodeMessage(coderr.CodeNotFound, "Joke not found")
)

type Joke struct {
	ID          string
	Title       string
	Text        string
	Explanation string
	CreatedAt   time.Time

	UserID string
}

type Jokes []Joke
