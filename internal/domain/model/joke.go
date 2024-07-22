package model

import (
	"time"

	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
)

var ErrJokeNotFound = coderr.NewCodeMessage(coderr.CodeNotFound, "Joke not found")

type Joke struct {
	ID          string
	Title       string
	Text        string
	Explanation string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type Jokes []Joke
