package model

import (
	"github.com/abc-valera/flugo-api-golang/internal/core/coderr"
	"github.com/abc-valera/flugo-api-golang/internal/core/model/common"
)

var (
	ErrJokeNotFound = coderr.NewMessage(coderr.CodeNotFound, "Joke not found")

	ErrJokeOwnerTitleAlreadyExists = coderr.NewMessage(coderr.CodeAlreadyExists, "Joke with such title already exists by such user")
)

type Joke struct {
	common.BaseModel
	UserID      string
	Title       string
	Text        string
	Explanation string
}

type Jokes []Joke
