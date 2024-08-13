package boilerSqliteDto

import (
	"github.com/abc-valera/netsly-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-golang/internal/domain/model"
)

func NewDomainJoke(joke *sqlboiler.Joke) model.Joke {
	if joke == nil {
		return model.Joke{}
	}

	return model.Joke{
		ID:          joke.ID,
		Title:       joke.Title,
		Text:        joke.Text,
		Explanation: joke.Explanation,
		CreatedAt:   joke.CreatedAt,
		UpdatedAt:   joke.UpdatedAt,
		DeletedAt:   joke.DeletedAt,
	}
}

func NewDomainJokes(jokes sqlboiler.JokeSlice) model.Jokes {
	var domainJokes model.Jokes
	for _, joke := range jokes {
		domainJokes = append(domainJokes, NewDomainJoke(joke))
	}
	return domainJokes
}
