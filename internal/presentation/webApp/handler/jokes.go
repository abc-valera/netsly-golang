package handler

import (
	"net/http"
	"strings"

	"github.com/abc-valera/netsly-golang/internal/domain/entity"
	"github.com/abc-valera/netsly-golang/internal/presentation/webApp/handler/session"
)

type Joke struct {
	jokeEntity entity.IJoke
}

func NewJoke(jokeEntity entity.IJoke) Joke {
	return Joke{
		jokeEntity: jokeEntity,
	}
}

func (h Joke) JokesPost(w http.ResponseWriter, r *http.Request) error {
	if _, err := h.jokeEntity.Create(r.Context(), entity.JokeCreateRequest{
		UserID:      session.GetUserID(r),
		Title:       r.FormValue("title"),
		Text:        r.FormValue("text"),
		Explanation: r.FormValue("explanation"),
	}); err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "title") {
			w.WriteHeader(491)
			w.Write([]byte(err.Error()))
			return nil
		}
		if strings.Contains(strings.ToLower(err.Error()), "text") {
			w.WriteHeader(492)
			w.Write([]byte(err.Error()))
			return nil
		}
		if strings.Contains(strings.ToLower(err.Error()), "explanation") {
			w.WriteHeader(493)
			w.Write([]byte(err.Error()))
			return nil
		}
		return err
	}

	w.Header().Add("HX-Reswap", "none")
	w.Header().Add("HX-Trigger", "newJoke")
	return nil
}
