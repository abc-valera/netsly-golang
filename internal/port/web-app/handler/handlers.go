package handler

import (
	"io/fs"

	"github.com/abc-valera/netsly-api-golang/internal/application"
	"github.com/abc-valera/netsly-api-golang/internal/domain"
)

type Handlers struct {
	Error Error

	Sign Sign
	Home Home

	Joke Joke
}

func NewHandlers(
	templateFS fs.FS,
	queries domain.Queries,
	entities domain.Entities,
	usecases application.UseCases,
) Handlers {
	return Handlers{
		Error: NewErrorHandler(templateFS),

		Sign: NewSign(templateFS, usecases.SignUseCase),
		Home: NewHome(templateFS, queries.User, queries.Joke),

		Joke: NewJoke(entities.Joke),
	}
}
