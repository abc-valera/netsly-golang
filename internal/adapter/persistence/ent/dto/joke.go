package dto

import (
	"github.com/abc-valera/flugo-api-golang/gen/ent"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/model"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/model/common"
)

func FromEntJokeToJoke(entJoke *ent.Joke) *model.Joke {
	if entJoke == nil {
		return nil
	}
	return &model.Joke{
		BaseModel: common.BaseModel{
			ID:        entJoke.ID,
			CreatedAt: entJoke.CreatedAt,
		},
		UserID:      entJoke.UserID,
		Title:       entJoke.Title,
		Text:        entJoke.Text,
		Explanation: entJoke.Explanation,
	}
}

func FromEntJokesToJokes(entJokes []*ent.Joke) model.Jokes {
	jokes := make(model.Jokes, len(entJokes))
	for i, entJoke := range entJokes {
		jokes[i] = FromEntJokeToJoke(entJoke)
	}
	return jokes
}
