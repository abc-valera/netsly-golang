package handler

import (
	"io/fs"
	"net/http"

	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/port/web-app/handler/common"
	"github.com/abc-valera/netsly-api-golang/internal/port/web-app/handler/cookie"
)

type Home struct {
	homeIndex common.ITemplate
}

func NewHome(
	templateFS fs.FS,
) Home {
	return Home{
		homeIndex: coderr.Must[common.ITemplate](common.NewTemplate(templateFS, "home/index", "layout/base")),
	}
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