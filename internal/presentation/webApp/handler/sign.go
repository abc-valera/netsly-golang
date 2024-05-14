package handler

import (
	"io/fs"
	"net/http"
	"strings"

	"github.com/abc-valera/netsly-api-golang/internal/application"
	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/presentation/webApp/cookie"
	"github.com/abc-valera/netsly-api-golang/internal/presentation/webApp/handler/tmpl"
)

type Sign struct {
	signIndex tmpl.ITemplate

	application.ISignUseCase
}

func NewSign(templateFS fs.FS, signUseCase application.ISignUseCase) Sign {
	return Sign{
		signIndex: coderr.Must(tmpl.NewTemplate(templateFS, "sign/index", "layout")),

		ISignUseCase: signUseCase,
	}
}

func (h Sign) SignGet(w http.ResponseWriter, r *http.Request) error {
	return h.signIndex.Render(w, nil)
}

func (h Sign) SignUpPost(w http.ResponseWriter, r *http.Request) error {
	err := r.ParseForm()
	if err != nil {
		return coderr.NewInternalErr(err)
	}

	if err := h.ISignUseCase.SignUp(r.Context(), application.SignUpRequest{
		Username: r.FormValue("username"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}); err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "username") {
			w.WriteHeader(491)
			w.Write([]byte(err.Error()))
			return nil
		}
		if strings.Contains(strings.ToLower(err.Error()), "email") {
			w.WriteHeader(492)
			w.Write([]byte(err.Error()))
			return nil
		}
		if strings.Contains(strings.ToLower(err.Error()), "password") {
			w.WriteHeader(493)
			w.Write([]byte(err.Error()))
			return nil
		}
		return err
	}

	resp, err := h.ISignUseCase.SignIn(r.Context(), application.SignInRequest{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	})
	if err != nil {
		return err
	}

	cookie.Set(w, cookie.AccessTokenKey, resp.AccessToken)
	cookie.Set(w, cookie.RefreshTokenKey, resp.RefreshToken)

	w.Header().Set("HX-Redirect", "/home")
	return nil
}

func (h Sign) SignInPost(w http.ResponseWriter, r *http.Request) error {
	err := r.ParseForm()
	if err != nil {
		return coderr.NewInternalErr(err)
	}

	resp, err := h.ISignUseCase.SignIn(r.Context(), application.SignInRequest{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	})
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "email") {
			w.WriteHeader(491)
			w.Write([]byte(err.Error()))
			return nil
		}
		if err == model.ErrUserNotFound {
			w.WriteHeader(491)
			w.Write([]byte(err.Error()))
			return nil
		}
		if strings.Contains(strings.ToLower(err.Error()), "password") {
			w.WriteHeader(492)
			w.Write([]byte(err.Error()))
			return nil
		}
		return err
	}

	cookie.Set(w, cookie.AccessTokenKey, resp.AccessToken)
	cookie.Set(w, cookie.RefreshTokenKey, resp.RefreshToken)

	w.Header().Set("HX-Redirect", "/home")
	return nil
}
