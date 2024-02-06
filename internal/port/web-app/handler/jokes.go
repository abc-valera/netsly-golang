package handler

import (
	"net/http"

	"github.com/abc-valera/netsly-api-golang/internal/domain/entity"
	"github.com/abc-valera/netsly-api-golang/internal/port/web-app/handler/session"
)

type Joke struct {
	jokeEntity entity.Joke
}

func NewJoke(jokeEntity entity.Joke) Joke {
	return Joke{
		jokeEntity: jokeEntity,
	}
}

func (h Joke) JokePost(w http.ResponseWriter, r *http.Request) error {
	return h.jokeEntity.Create(r.Context(), entity.JokeCreateRequest{
		UserID:      session.GetUserID(r),
		Title:       r.FormValue("title"),
		Text:        r.FormValue("text"),
		Explanation: r.FormValue("explanation"),
	})
}
