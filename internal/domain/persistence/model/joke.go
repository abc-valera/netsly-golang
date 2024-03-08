package model

import (
	"github.com/abc-valera/netsly-api-golang/internal/domain/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model/common"
)

var (
	ErrJokeNotFound = coderr.NewCodeMessage(coderr.CodeNotFound, "Joke not found")
)

type Joke struct {
	common.BaseEntity
	Title       string
	Text        string
	Explanation string

	UserID string
}

type Jokes []Joke
