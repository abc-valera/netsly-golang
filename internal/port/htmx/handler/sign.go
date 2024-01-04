package handler

import (
	"fmt"
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
		[]string{"sign/index_sign_up_modal", "layout/modal"},
		[]string{"sign/index_sign_in_modal", "layout/modal"},
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

func (h SignHandler) SignUpGet(w http.ResponseWriter, r *http.Request) error {
	return h.t.Render(w, "sign/index_sign_up_modal", nil)
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

func (h SignHandler) SignInGet(w http.ResponseWriter, r *http.Request) error {
	return h.t.Render(w, "sign/index_sign_in_modal", nil)
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

func (h SignHandler) SignAlreadyGet(w http.ResponseWriter, r *http.Request) error {
	return codeerr.NewInternal(fmt.Errorf("not implemented"))
}

func (h SignHandler) SignOutGet(w http.ResponseWriter, r *http.Request) error {
	return codeerr.NewInternal(fmt.Errorf("not implemented"))
}

func (h SignHandler) SignOutPost(w http.ResponseWriter, r *http.Request) error {
	return codeerr.NewInternal(fmt.Errorf("not implemented"))
}
