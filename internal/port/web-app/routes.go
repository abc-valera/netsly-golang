package webapp

import (
	"net/http"

	"github.com/abc-valera/netsly-api-golang/internal/domain"
	"github.com/abc-valera/netsly-api-golang/internal/domain/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/port/web-app/handler"
	"github.com/abc-valera/netsly-api-golang/internal/port/web-app/middleware"
	"github.com/go-chi/chi/v5"
)

func initRoutes(r *chi.Mux, staticPath string, services domain.Services, handlers handler.Handlers) {
	// Static files (before middleware)
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir(staticPath))))

	r.Route("/", func(r chi.Router) {
		// Middleware
		r.Use(
			middleware.Logger,
			middleware.Recoverer,
			middleware.NewSessionMiddleware(services.TokenMaker),
		)

		// 404 handler
		r.NotFound(func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "/error/404", http.StatusMovedPermanently)
		})

		// Error routes
		r.Route("/error", func(r chi.Router) {
			r.Get("/401", newHandlerFunc(handlers.Error.Error401Get))
			r.Get("/403", newHandlerFunc(handlers.Error.Error403Get))
			r.Get("/404", newHandlerFunc(handlers.Error.Error404Get))
			r.Get("/500", newHandlerFunc(handlers.Error.Error500Get))
		})

		// Sign routes
		r.Route("/sign", func(r chi.Router) {
			r.Get("/", newHandlerFunc(handlers.Sign.SignGet))
			r.Post("/up", newHandlerFunc(handlers.Sign.SignUpPost))
			r.Post("/in", newHandlerFunc(handlers.Sign.SignInPost))
		})

		// Home routes
		r.Route("/home", func(r chi.Router) {
			r.Get("/", newHandlerFunc(handlers.Home.HomeGet))
			r.Get("/partial/jokes", newHandlerFunc(handlers.Home.HomePartialJokes))
		})

		// Jokes routes
		r.Route("/jokes", func(r chi.Router) {
			r.Post("/", newHandlerFunc(handlers.Joke.JokePost))
		})
	})
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

		if code == coderr.CodeInvalidArgument || code == coderr.CodeNotFound || code == coderr.CodeAlreadyExists {
			msg := coderr.ErrorMessage(err)
			if msg == "" {
				msg = "Something wrong with your input!"
			}
			w.Write([]byte(msg))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if code == coderr.CodeUnauthenticated {
			http.Redirect(w, r, "/error/401", http.StatusUnauthorized)
			return
		}

		if code == coderr.CodePermissionDenied {
			http.Redirect(w, r, "/error/403", http.StatusForbidden)
			return
		}

		global.Log.Error("REQUEST_ERROR", "err", err.Error())
		http.Redirect(w, r, "/error/500", http.StatusMovedPermanently)
	}
}
