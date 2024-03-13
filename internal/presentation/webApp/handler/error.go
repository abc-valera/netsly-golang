package handler

import (
	"io/fs"
	"net/http"

	"github.com/abc-valera/netsly-api-golang/internal/domain/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/presentation/webApp/handler/tmpl"
)

type Error struct {
	error401 tmpl.ITemplate
	error403 tmpl.ITemplate
	error404 tmpl.ITemplate
	error500 tmpl.ITemplate
}

func NewErrorHandler(templateFS fs.FS) Error {
	return Error{
		error401: coderr.MustWithVal(tmpl.NewTemplate(templateFS, "error/401", "layout")),
		error403: coderr.MustWithVal(tmpl.NewTemplate(templateFS, "error/403", "layout")),
		error404: coderr.MustWithVal(tmpl.NewTemplate(templateFS, "error/404", "layout")),
		error500: coderr.MustWithVal(tmpl.NewTemplate(templateFS, "error/500", "layout")),
	}
}

func (h Error) Error401Get(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusUnauthorized)
	return h.error401.Render(w, nil)
}

func (h Error) Error403Get(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusForbidden)
	return h.error403.Render(w, nil)
}

func (h Error) Error404Get(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusNotFound)
	return h.error404.Render(w, nil)
}

func (h Error) Error500Get(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusInternalServerError)
	if err := h.error500.Render(w, nil); err != nil {
		global.Log().Error("Error500Get", "err", err.Error())
		w.Write([]byte("500 - Internal server error"))
	}
	return nil
}
