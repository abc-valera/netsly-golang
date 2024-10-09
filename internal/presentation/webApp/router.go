package webApp

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/abc-valera/netsly-golang/internal/domain/global"
	"github.com/abc-valera/netsly-golang/internal/domain/service"
	"github.com/abc-valera/netsly-golang/internal/domain/util/coderr"
	"github.com/abc-valera/netsly-golang/internal/presentation/webApp/handler"
	"github.com/abc-valera/netsly-golang/internal/presentation/webApp/middleware"
	"github.com/go-chi/chi/v5"
)

//go:embed static
var staticFiles embed.FS

func newRouter(
	r *chi.Mux,
	_ service.Services,
	handlers handler.Handlers,
) {
	// Static files (before middleware)
	r.Handle("/static", http.FileServer(http.FS(coderr.Must(fs.Sub(staticFiles, "static")))))

	r.Route("/", func(r chi.Router) {
		// Middleware
		r.Use(
			middleware.Logger,
			middleware.Recoverer,
			middleware.SessionMiddleware,
		)

		// 404 handler
		r.NotFound(func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(404)
			w.Header().Set("HX-Redirect", "/error/404")
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
			r.Get("/partial/jokes", newHandlerFunc(handlers.Home.HomePartialJokesGet))
		})

		// Jokes routes
		r.Route("/jokes", func(r chi.Router) {
			r.Post("/", newHandlerFunc(handlers.Joke.JokesPost))
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
	if err := h(w, r); err != nil {
		switch coderr.ErrorCode(err) {
		case coderr.CodeInvalidArgument, coderr.CodeNotFound, coderr.CodeAlreadyExists:
			return
		case coderr.CodeUnauthenticated:
			w.WriteHeader(401)
			w.Header().Set("HX-Redirect", "/error/401")
			return
		case coderr.CodePermissionDenied:
			w.WriteHeader(403)
			w.Header().Set("HX-Redirect", "/error/403")
			return
		default:
			global.Log().Error("REQUEST_ERROR", "err", err.Error())
			w.WriteHeader(500)
			w.Header().Set("HX-Redirect", "/error/500")
		}
	}
}
