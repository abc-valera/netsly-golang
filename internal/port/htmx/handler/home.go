package handler

import (
	"io/fs"
	"net/http"

	"github.com/abc-valera/flugo-api-golang/internal/port/htmx/handler/common"
	"github.com/abc-valera/flugo-api-golang/internal/port/htmx/handler/cookie"
)

type Home struct {
	t common.Templates
}

func NewHome(
	templateFS fs.FS,
) (Home, error) {
	t, err := common.NewTemplates(false, templateFS,
		[]string{"home/index", "layout/base"},
	)
	if err != nil {
		return Home{}, err
	}

	return Home{
		t: t,
	}, nil
}

func (h Home) HomeGet(w http.ResponseWriter, r *http.Request) error {
	access, err := cookie.Get(r, cookie.AccessTokenKey)
	if err != nil {
		return err
	}

	return h.t.Render(w, "home/index", map[string]any{
		"AccessToken": access,
	})
}
