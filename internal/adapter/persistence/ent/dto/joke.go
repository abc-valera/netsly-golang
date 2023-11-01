package dto

import (
	"github.com/abc-valera/flugo-api-golang/gen/ent"
	"github.com/abc-valera/flugo-api-golang/internal/adapter/persistence/ent/dto/common"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/entity"
)

func FromEntJokeToJoke(entJoke *ent.Joke) *entity.Joke {
	if entJoke == nil {
		return nil
	}
	return &entity.Joke{
		BaseEntity:  common.FromEntToBaseEntity(entJoke.ID, entJoke.CreatedAt),
		UserID:      entJoke.UserID,
		Title:       entJoke.Title,
		Text:        entJoke.Text,
		Explanation: entJoke.Explanation,
	}
}

func FromEntJokesToJokes(entJokes []*ent.Joke) entity.Jokes {
	jokes := make(entity.Jokes, len(entJokes))
	for i, entJoke := range entJokes {
		jokes[i] = FromEntJokeToJoke(entJoke)
	}
	return jokes
}
