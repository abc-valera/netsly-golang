package handler

import (
	"io/fs"
	"net/http"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/coderr"
	"github.com/abc-valera/flugo-api-golang/internal/port/htmx/handler/common"
	"github.com/abc-valera/flugo-api-golang/internal/port/htmx/handler/cookie"
)

type Home struct {
	homeIndex common.ITemplate
}

func NewHome(
	templateFS fs.FS,
) (Home, error) {
	return Home{
		homeIndex: coderr.Must[common.ITemplate](common.NewTemplate(templateFS, "home/index", "layout/base")),
	}, nil
}

func (h Home) HomeGet(w http.ResponseWriter, r *http.Request) error {
	access, err := cookie.Get(r, cookie.AccessTokenKey)
	if err != nil {
		return err
	}

	return h.homeIndex.Render(w, map[string]any{
		"AccessToken": access,
	})
}
