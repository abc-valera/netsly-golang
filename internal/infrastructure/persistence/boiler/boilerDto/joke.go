package boilerDto

import (
	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/boiler/errors"
)

func NewDomainJoke(joke *sqlboiler.Joke) model.Joke {
	if joke == nil {
		return model.Joke{}
	}

	return model.Joke{
		ID:          joke.ID,
		Text:        joke.Text,
		Explanation: joke.Explanation,
		CreatedAt:   joke.CreatedAt,
		UserID:      joke.UserID,
	}
}

func NewDomainJokeWithErrHandle(joke *sqlboiler.Joke, err error) (model.Joke, error) {
	return NewDomainJoke(joke), errors.HandleErr(err)
}

func NewDomainJokes(jokes sqlboiler.JokeSlice) model.Jokes {
	var domainJokes model.Jokes
	for _, joke := range jokes {
		domainJokes = append(domainJokes, NewDomainJoke(joke))
	}
	return domainJokes
}

func NewDomainJokesWithErrHandle(jokes sqlboiler.JokeSlice, err error) (model.Jokes, error) {
	return NewDomainJokes(jokes), errors.HandleErr(err)
}
