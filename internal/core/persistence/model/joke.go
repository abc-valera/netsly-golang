package model

import (
	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/core/persistence/model/common"
)

var (
	ErrJokeNotFound = coderr.NewMessage(coderr.CodeNotFound, "Joke not found")
)

type Joke struct {
	common.BaseModel
	UserID      string
	Title       string
	Text        string
	Explanation string
}

type Jokes []Joke
