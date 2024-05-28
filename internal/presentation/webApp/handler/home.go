package handler

import (
	"context"
	"io/fs"
	"net/http"

	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/entity"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query/selector"
	"github.com/abc-valera/netsly-api-golang/internal/presentation/webApp/handler/session"
	"github.com/abc-valera/netsly-api-golang/internal/presentation/webApp/handler/templates"
)

type Home struct {
	homeIndex    templates.ITemplate
	partialJokes templates.ITemplate

	user entity.IUser
	joke entity.IJoke
}

func NewHome(
	templateFS fs.FS,
	user entity.IUser,
	joke entity.IJoke,
) Home {
	return Home{
		homeIndex:    coderr.Must(templates.NewTemplate(templateFS, "home/index/index", "home/layout", "layout")),
		partialJokes: coderr.Must(templates.NewTemplate(templateFS, "home/index/partial_jokes.html")),

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
		selector.Selector{
			Order:  "desc",
			Limit:  5,
			Offset: 0,
		},
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
		selector.Selector{
			Order:  "desc",
			Limit:  5,
			Offset: 0,
		},
	)
	if err != nil {
		return err
	}

	return h.partialJokes.Render(w, templates.Data{
		"Jokes": jokes,
	})
}
