package handler

import (
	"context"
	"io/fs"
	"net/http"

	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query/spec"
	"github.com/abc-valera/netsly-api-golang/internal/presentation/webApp/handler/session"
	"github.com/abc-valera/netsly-api-golang/internal/presentation/webApp/handler/tmpl"
)

type Home struct {
	homeIndex    tmpl.ITemplate
	partialJokes tmpl.ITemplate

	userQuery query.IUser
	jokeQuery query.IJoke
}

func NewHome(
	templateFS fs.FS,
	userQuery query.IUser,
	jokeQuery query.IJoke,
) Home {
	return Home{
		homeIndex:    coderr.Must(tmpl.NewTemplate(templateFS, "home/index/index", "home/layout", "layout")),
		partialJokes: coderr.Must(tmpl.NewTemplate(templateFS, "home/index/partial_jokes.html")),

		userQuery: userQuery,
		jokeQuery: jokeQuery,
	}
}

func (h Home) HomeGet(w http.ResponseWriter, r *http.Request) error {
	userID := r.Context().Value(session.UserIDKey).(string)

	user, err := h.userQuery.GetByID(r.Context(), userID)
	if err != nil {
		return err
	}

	jokes, err := h.jokeQuery.GetAllByUserID(
		r.Context(),
		userID,
		spec.NewSelectParams("desc", 5, 0),
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
	jokes, err := h.jokeQuery.GetAllByUserID(
		context.Background(),
		session.GetUserID(r),
		spec.NewSelectParams("desc", 5, 0),
	)
	if err != nil {
		return err
	}

	return h.partialJokes.Render(w, tmpl.Data{
		"Jokes": jokes,
	})
}
