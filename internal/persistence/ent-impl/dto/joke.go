package dto

import (
	"github.com/abc-valera/netsly-api-golang/gen/ent"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model/common"
	errhandler "github.com/abc-valera/netsly-api-golang/internal/persistence/ent-impl/errors"
)

func FromEntJoke(entJoke *ent.Joke) model.Joke {
	if entJoke == nil {
		return model.Joke{}
	}
	return model.Joke{
		BaseEntity: common.BaseEntity{
			ID:        entJoke.ID,
			CreatedAt: entJoke.CreatedAt,
		},
		UserID:      entJoke.Edges.User.ID,
		Title:       entJoke.Title,
		Text:        entJoke.Text,
		Explanation: entJoke.Explanation,
	}
}

func FromEntJokeToJokeWithErrHandle(entJoke *ent.Joke, err error) (model.Joke, error) {
	return FromEntJoke(entJoke), errhandler.HandleErr(err)
}

func FromEntJokesToJokes(entJokes []*ent.Joke) model.Jokes {
	jokes := make(model.Jokes, len(entJokes))
	for i, entJoke := range entJokes {
		jokes[i] = FromEntJoke(entJoke)
	}
	return jokes
}

func FromEntJokesToJokesWithErrHandle(entJokes []*ent.Joke, err error) (model.Jokes, error) {
	return FromEntJokesToJokes(entJokes), errhandler.HandleErr(err)
}
