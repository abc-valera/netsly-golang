package restDto

import (
	"github.com/abc-valera/netsly-golang/gen/ogen"
	"github.com/abc-valera/netsly-golang/internal/domain/model"
)

func NewJoke(joke model.Joke) *ogen.Joke {
	return &ogen.Joke{
		ID:          joke.ID,
		Title:       joke.Title,
		Text:        joke.Text,
		Explanation: ogen.NewOptString(joke.Explanation),
		CreatedAt:   joke.CreatedAt,
		UserID:      joke.UserID,
	}
}

func NewJokes(domainJokes []model.Joke) ogen.Jokes {
	var jokes ogen.Jokes
	for _, joke := range domainJokes {
		jokes = append(jokes, *NewJoke(joke))
	}
	return jokes
}
