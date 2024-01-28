package dto

import (
	"github.com/abc-valera/flugo-api-golang/gen/ent"
	errhandler "github.com/abc-valera/flugo-api-golang/internal/adapter/persistence/ent-impl/errors"
	"github.com/abc-valera/flugo-api-golang/internal/core/persistence/model"
	"github.com/abc-valera/flugo-api-golang/internal/core/persistence/model/common"
)

func FromEntJokeToJoke(entJoke *ent.Joke) model.Joke {
	if entJoke == nil {
		return model.Joke{}
	}
	return model.Joke{
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

func FromEntJokeToJokeWithErrHandle(entJoke *ent.Joke, err error) (model.Joke, error) {
	return FromEntJokeToJoke(entJoke), errhandler.HandleErr(err)
}

func FromEntJokesToJokes(entJokes []*ent.Joke) model.Jokes {
	jokes := make(model.Jokes, len(entJokes))
	for i, entJoke := range entJokes {
		jokes[i] = FromEntJokeToJoke(entJoke)
	}
	return jokes
}

func FromEntJokesToJokesWithErrHandle(entJokes []*ent.Joke, err error) (model.Jokes, error) {
	return FromEntJokesToJokes(entJokes), errhandler.HandleErr(err)
}
