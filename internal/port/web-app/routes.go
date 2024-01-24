package webapp

import (
	"net/http"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/coderr"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/service"
	"github.com/abc-valera/flugo-api-golang/internal/port/web-app/handler"
	"github.com/abc-valera/flugo-api-golang/internal/port/web-app/middleware"
	"github.com/go-chi/chi/v5"
)

func initRoutesMiddlewares(r *chi.Mux, handlers handler.Handlers) {
	// middlewares
	r.Use(middleware.Tracer)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// 404 handler
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/error/404", http.StatusMovedPermanently)
	})

	// static files
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("internal/port/web-app/static"))))

	r.Get("/error/401", newHandlerFunc(handlers.Error.Error401Get))
	r.Get("/error/403", newHandlerFunc(handlers.Error.Error403Get))
	r.Get("/error/404", newHandlerFunc(handlers.Error.Error404Get))
	r.Get("/error/500", newHandlerFunc(handlers.Error.Error500Get))

	r.Get("/sign", newHandlerFunc(handlers.Sign.SignGet))
	r.Post("/sign/up", newHandlerFunc(handlers.Sign.SignUpPost))
	r.Post("/sign/in", newHandlerFunc(handlers.Sign.SignInPost))

	r.Get("/home", newHandlerFunc(handlers.Home.HomeGet))
}

func newHandlerFunc(h handlerWithError) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	}
}

type handlerWithError func(w http.ResponseWriter, r *http.Request) error

func (h handlerWithError) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h(w, r)
	if err != nil {
		code := coderr.ErrorCode(err)

		// These errors should be handled in handlers
		// if code == codeerr.CodeInvalidArgument ||
		// 	code == codeerr.CodeNotFound ||
		// 	code == codeerr.CodeAlreadyExists {
		// 	return
		// }

		if code == coderr.CodeUnauthenticated {
			http.Redirect(w, r, "/error/401", http.StatusUnauthorized)
			return
		}

		if code == coderr.CodePermissionDenied {
			http.Redirect(w, r, "/error/403", http.StatusForbidden)
			return
		}

		service.Log.Error("REQUEST_ERROR", "err", err.Error())
		http.Redirect(w, r, "/error/500", http.StatusMovedPermanently)
	}
}
