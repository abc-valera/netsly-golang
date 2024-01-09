package handler

import (
	"io/fs"
	"net/http"

	"github.com/abc-valera/flugo-api-golang/internal/core/application"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/codeerr"
	"github.com/abc-valera/flugo-api-golang/internal/port/htmx/handler/common"
)

type SignHandler struct {
	t common.Templates
	application.SignUseCase
}

func NewSignHandler(templateFS fs.FS, signUseCase application.SignUseCase) (SignHandler, error) {
	t, err := common.NewTemplates(false, templateFS,
		[]string{"sign/index", "layout/base"},
	)
	if err != nil {
		return SignHandler{}, err
	}
	return SignHandler{
		t:           t,
		SignUseCase: signUseCase,
	}, nil
}

func (h SignHandler) SignGet(w http.ResponseWriter, r *http.Request) error {
	return h.t.Render(w, "sign/index", nil)
}

func (h SignHandler) SignUpPost(w http.ResponseWriter, r *http.Request) error {
	err := r.ParseForm()
	if err != nil {
		return codeerr.NewInternal(err)
	}

	if err := h.SignUseCase.SignUp(r.Context(), application.SignUpRequest{
		Username: r.FormValue("username"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}); err != nil {
		return err
	}

	return nil
}

func (h SignHandler) SignInPost(w http.ResponseWriter, r *http.Request) error {
	err := r.ParseForm()
	if err != nil {
		return codeerr.NewInternal(err)
	}

	resp, err := h.SignUseCase.SignIn(r.Context(), application.SignInRequest{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	})
	if err != nil {
		return err
	}
	_ = resp

	return nil
}
