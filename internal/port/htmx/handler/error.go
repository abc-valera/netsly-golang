package handler

import (
	"io/fs"
	"net/http"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/coderr"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/service"
	"github.com/abc-valera/flugo-api-golang/internal/port/htmx/handler/common"
)

type Error struct {
	error401 common.ITemplate
	error403 common.ITemplate
	error404 common.ITemplate
	error500 common.ITemplate
}

func NewErrorHandler(templateFS fs.FS) (Error, error) {
	return Error{
		error401: coderr.Must[common.ITemplate](common.NewTemplate(templateFS, "error/401", "layout/base")),
		error403: coderr.Must[common.ITemplate](common.NewTemplate(templateFS, "error/403", "layout/base")),
		error404: coderr.Must[common.ITemplate](common.NewTemplate(templateFS, "error/404", "layout/base")),
		error500: coderr.Must[common.ITemplate](common.NewTemplate(templateFS, "error/500", "layout/base")),
	}, nil
}

func (h Error) Error401Get(w http.ResponseWriter, r *http.Request) error {
	return h.error401.Render(w, nil)
}

func (h Error) Error403Get(w http.ResponseWriter, r *http.Request) error {
	return h.error403.Render(w, nil)
}

func (h Error) Error404Get(w http.ResponseWriter, r *http.Request) error {
	return h.error404.Render(w, nil)
}

func (h Error) Error500Get(w http.ResponseWriter, r *http.Request) error {
	if err := h.error500.Render(w, nil); err != nil {
		service.Log.Error("Error500Get", "err", err.Error())
		w.Write([]byte("500 - Internal server error"))
	}
	return nil
}
