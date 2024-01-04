package handler

import (
	"io/fs"
	"net/http"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/service"
	"github.com/abc-valera/flugo-api-golang/internal/port/htmx/handler/common"
)

type ErrorHandler struct {
	t common.Templates
}

func NewErrorHandler(templateFS fs.FS) (ErrorHandler, error) {
	t, err := common.NewTemplates(false, templateFS,
		[]string{"error/401", "layout/base"},
		[]string{"error/403", "layout/base"},
		[]string{"error/404", "layout/base"},
		[]string{"error/500", "layout/base"},
	)
	if err != nil {
		return ErrorHandler{}, err
	}
	return ErrorHandler{
		t: t,
	}, nil
}

func (h ErrorHandler) Error401Get(w http.ResponseWriter, r *http.Request) error {
	return h.t.Render(w, "error/401", nil)
}

func (h ErrorHandler) Error403Get(w http.ResponseWriter, r *http.Request) error {
	return h.t.Render(w, "error/403", nil)
}

func (h ErrorHandler) Error404Get(w http.ResponseWriter, r *http.Request) error {
	return h.t.Render(w, "error/404", nil)
}

func (h ErrorHandler) Error500Get(w http.ResponseWriter, r *http.Request) error {
	if err := h.t.Render(w, "error/500", nil); err != nil {
		service.Log.Error("Error500Get", "err", err.Error())
		w.Write([]byte("500 - Internal server error"))
	}
	return nil
}
