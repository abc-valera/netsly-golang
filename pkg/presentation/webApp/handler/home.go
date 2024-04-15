package handler

import (
	"context"
	"io/fs"
	"net/http"

	"github.com/abc-valera/netsly-api-golang/pkg/core/coderr"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/entity"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/query/selectParams"
	"github.com/abc-valera/netsly-api-golang/pkg/presentation/webApp/handler/session"
	"github.com/abc-valera/netsly-api-golang/pkg/presentation/webApp/handler/tmpl"
)

type Home struct {
	homeIndex    tmpl.ITemplate
	partialJokes tmpl.ITemplate

	user entity.IUser
	joke entity.IJoke
}

func NewHome(
	templateFS fs.FS,
	user entity.IUser,
	joke entity.IJoke,
) Home {
	return Home{
		homeIndex:    coderr.Must(tmpl.NewTemplate(templateFS, "home/index/index", "home/layout", "layout")),
		partialJokes: coderr.Must(tmpl.NewTemplate(templateFS, "home/index/partial_jokes.html")),

		user: user,
		joke: joke,
	}
}

func (h Home) HomeGet(w http.ResponseWriter, r *http.Request) error {
	userID := r.Context().Value(session.UserIDKey).(string)

	user, err := h.user.GetByID(r.Context(), userID)
	if err != nil {
		return err
	}

	jokes, err := h.joke.GetAllByUserID(
		r.Context(),
		userID,
		selectParams.NewSelectParams("desc", 5, 0),
	)
	if err != nil {
		return err
	}

	return h.homeIndex.Render(w, map[string]any{
		"User":  user,
		"Jokes": jokes,
	})
}

func (h Home) HomePartialJokesGet(w http.ResponseWriter, r *http.Request) error {
	jokes, err := h.joke.GetAllByUserID(
		context.Background(),
		session.GetUserID(r),
		selectParams.NewSelectParams("desc", 5, 0),
	)
	if err != nil {
		return err
	}

	return h.partialJokes.Render(w, tmpl.Data{
		"Jokes": jokes,
	})
}
