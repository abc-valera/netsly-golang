package handler

import (
	"io/fs"
	"net/http"

	"github.com/abc-valera/netsly-api-golang/internal/application"
	"github.com/abc-valera/netsly-api-golang/internal/domain/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/port/web-app/cookie"
	"github.com/abc-valera/netsly-api-golang/internal/port/web-app/handler/tmpl"
)

type Sign struct {
	signIndex tmpl.ITemplate

	application.SignUseCase
}

func NewSign(templateFS fs.FS, signUseCase application.SignUseCase) Sign {
	return Sign{
		signIndex: coderr.Must[tmpl.ITemplate](tmpl.NewTemplate(templateFS, "sign/index", "layout")),

		SignUseCase: signUseCase,
	}
}

func (h Sign) SignGet(w http.ResponseWriter, r *http.Request) error {
	return h.signIndex.Render(w, nil)
}

func (h Sign) SignUpPost(w http.ResponseWriter, r *http.Request) error {
	err := r.ParseForm()
	if err != nil {
		return coderr.NewInternal(err)
	}

	if err := h.SignUseCase.SignUp(r.Context(), application.SignUpRequest{
		Username: r.FormValue("username"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}); err != nil {
		return err
	}

	resp, err := h.SignUseCase.SignIn(r.Context(), application.SignInRequest{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	})
	if err != nil {
		return err
	}

	cookie.Set(w, cookie.AccessTokenKey, resp.AccessToken)
	cookie.Set(w, cookie.RefreshTokenKey, resp.RefreshToken)

	http.Redirect(w, r, "/home", http.StatusMovedPermanently)
	return nil
}

func (h Sign) SignInPost(w http.ResponseWriter, r *http.Request) error {
	err := r.ParseForm()
	if err != nil {
		return coderr.NewInternal(err)
	}

	resp, err := h.SignUseCase.SignIn(r.Context(), application.SignInRequest{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	})
	if err != nil {
		return err
	}

	cookie.Set(w, cookie.AccessTokenKey, resp.AccessToken)
	cookie.Set(w, cookie.RefreshTokenKey, resp.RefreshToken)

	http.Redirect(w, r, "/home", http.StatusMovedPermanently)
	return nil
}
