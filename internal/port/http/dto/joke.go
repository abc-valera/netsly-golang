package dto

import (
	"github.com/abc-valera/flugo-api-golang/gen/ogen"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/entity"
	"github.com/abc-valera/flugo-api-golang/internal/port/http/dto/common"
)

func NewJokeResponse(joke *entity.Joke) *ogen.Joke {
	if joke == nil {
		return &ogen.Joke{}
	}
	return &ogen.Joke{
		ID:          joke.ID,
		UserID:      joke.UserID,
		Title:       joke.Title,
		Text:        joke.Text,
		Explanation: common.NewOptString(joke.Explanation),
		CreatedAt:   joke.CreatedAt,
	}
}

func NewJokesResponse(jokes entity.Jokes) *ogen.Jokes {
	var res []ogen.Joke
	for _, joke := range jokes {
		res = append(res, *NewJokeResponse(joke))
	}
	return &ogen.Jokes{Jokes: res}
}
