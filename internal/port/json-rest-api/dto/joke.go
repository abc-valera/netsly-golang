package dto

import (
	"github.com/abc-valera/flugo-api-golang/gen/ogen"
	"github.com/abc-valera/flugo-api-golang/internal/core/persistence/model"
)

func NewJokeResponse(joke model.Joke) *ogen.Joke {
	return &ogen.Joke{
		ID:          joke.ID,
		UserID:      joke.UserID,
		Title:       joke.Title,
		Text:        joke.Text,
		Explanation: NewOptString(joke.Explanation),
		CreatedAt:   joke.CreatedAt,
	}
}

func NewJokesResponse(jokes model.Jokes) *ogen.Jokes {
	var res []ogen.Joke
	for _, joke := range jokes {
		res = append(res, *NewJokeResponse(joke))
	}
	return &ogen.Jokes{Jokes: res}
}
