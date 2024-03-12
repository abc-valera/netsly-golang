package dto

import (
	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
	"github.com/abc-valera/netsly-api-golang/internal/persistence/sqlboilerImpl/errors"
)

func ToDomainJoke(joke *sqlboiler.Joke) model.Joke {
	if joke == nil {
		return model.Joke{}
	}

	return model.Joke{
		ID:          joke.ID,
		Text:        joke.Text,
		Explanation: joke.Explanation.String,
		CreatedAt:   joke.CreatedAt,
		UserID:      joke.UserID,
	}
}

func ToDomainJokeWithErrHandle(joke *sqlboiler.Joke, err error) (model.Joke, error) {
	return ToDomainJoke(joke), errors.HandleErr(err)
}

func ToDomainJokes(jokes sqlboiler.JokeSlice) model.Jokes {
	var domainJokes model.Jokes
	for _, joke := range jokes {
		domainJokes = append(domainJokes, ToDomainJoke(joke))
	}
	return domainJokes
}

func ToDomainJokesWithErrHandle(jokes sqlboiler.JokeSlice, err error) (model.Jokes, error) {
	return ToDomainJokes(jokes), errors.HandleErr(err)
}
