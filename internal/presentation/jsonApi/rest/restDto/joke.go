package restDto

import (
	"github.com/abc-valera/netsly-api-golang/gen/ogen"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
)

func NewJoke(joke model.Joke) *ogen.Joke {
	return &ogen.Joke{
		ID:          joke.ID,
		Title:       joke.Title,
		Text:        joke.Text,
		Explanation: ogen.NewOptString(joke.Explanation),
		CreatedAt:   joke.CreatedAt,
	}
}

func NewJokes(jokes model.Jokes) *ogen.Jokes {
	var res []ogen.Joke
	for _, joke := range jokes {
		res = append(res, *NewJoke(joke))
	}
	return &ogen.Jokes{Jokes: res}
}
