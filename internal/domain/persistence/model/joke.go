package model

import (
	"github.com/abc-valera/netsly-api-golang/internal/domain/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model/common"
)

var (
	ErrJokeNotFound = coderr.NewMessage(coderr.CodeNotFound, "Joke not found")
)

type Joke struct {
	common.BaseEntity
	UserID      string
	Title       string
	Text        string
	Explanation string
}

type Jokes []Joke
