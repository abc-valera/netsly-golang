package handler

import (
	"io/fs"
	"net/http"
	"strings"

	"github.com/abc-valera/netsly-golang/internal/application"
	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/util/coderr"
	"github.com/abc-valera/netsly-golang/internal/presentation/webApp/cookie"
	"github.com/abc-valera/netsly-golang/internal/presentation/webApp/handler/templates"
)

type Sign struct {
	signIndex templates.ITemplate

	application.ISignUsecase
}

func NewSign(templateFS fs.FS, signUsecase application.ISignUsecase) Sign {
	return Sign{
		signIndex: coderr.Must(templates.NewTemplate(templateFS, "sign/index", "layout")),

		ISignUsecase: signUsecase,
	}
}

func (h Sign) SignGet(w http.ResponseWriter, _ *http.Request) error {
	return h.signIndex.Render(w, nil)
}

func (h Sign) SignUpPost(w http.ResponseWriter, r *http.Request) error {
	err := r.ParseForm()
	if err != nil {
		return coderr.NewInternalErr(err)
	}

	if _, err := h.ISignUsecase.SignUp(r.Context(), application.SignUpRequest{
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

	user, err := h.ISignUsecase.SignIn(r.Context(), application.SignInRequest{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	})
	if err != nil {
		return err
	}

	cookie.Set(w, cookie.UserIDKey, user.ID)

	w.Header().Set("HX-Redirect", "/home")
	return nil
}

func (h Sign) SignInPost(w http.ResponseWriter, r *http.Request) error {
	err := r.ParseForm()
	if err != nil {
		return coderr.NewInternalErr(err)
	}

	user, err := h.ISignUsecase.SignIn(r.Context(), application.SignInRequest{
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

	cookie.Set(w, cookie.UserIDKey, user.ID)

	w.Header().Set("HX-Redirect", "/home")
	return nil
}
