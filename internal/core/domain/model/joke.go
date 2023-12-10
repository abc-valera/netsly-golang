package model

import (
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/codeerr"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/model/common"
)

var (
	ErrJokeNotFound = codeerr.NewMessageErr(codeerr.CodeNotFound, "Joke not found")

	ErrJokeOwnerTitleAlreadyExists = codeerr.NewMessageErr(codeerr.CodeAlreadyExists, "Joke with such title already exists by such user")
)

type Joke struct {
	common.BaseModel
	UserID      string
	Title       string
	Text        string
	Explanation string
}

type JokeUpdate struct {
	Title       *string
	Text        *string
	Explanation *string
}

type Jokes []*Joke
