package http

import (
	"net/http"

	"github.com/abc-valera/flugo-api-golang/gen/ogen"
	"github.com/abc-valera/flugo-api-golang/internal/core/application"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/codeerr"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/service"
	"github.com/abc-valera/flugo-api-golang/internal/port/http/handler/comments"
	"github.com/abc-valera/flugo-api-golang/internal/port/http/handler/likes"
	"github.com/abc-valera/flugo-api-golang/internal/port/http/handler/me"
	"github.com/abc-valera/flugo-api-golang/internal/port/http/handler/other"
	"github.com/abc-valera/flugo-api-golang/internal/port/http/handler/sign"
	"github.com/abc-valera/flugo-api-golang/internal/port/http/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func RunServer(
	port string,
	docsPath string,
	repos repository.Repositories,
	services service.Services,
	usecases application.UseCases,
) error {
	// Init handlers (ogenHandler implements ogen.Server interface)
	ogenHandler := &struct {
		other.ErrorHandler
		sign.SignHandler
		me.MeHandler
		me.MeJokesHandler
		me.MeCommentsHandler
		me.MeLikesHandler
		comments.CommentsHandler
		likes.LikesHandler
	}{
		ErrorHandler:      handler.NewErrorHandler(),
		SignHandler:       handler.NewSignHandler(usecases.SignUseCase),
		MeHandler:         handler.NewMeHandler(queries.User, domains.User),
		MeJokesHandler:    handler.NewMeJokesHandler(queries.Joke, domains.Joke),
		MeCommentsHandler: handler.NewMeCommentsHandler(queries.Comment, domains.Comment),
		MeLikesHandler:    handler.NewMeLikesHandler(queries.Like, domains.Like),
		CommentsHandler:   handler.NewCommentsHandler(queries.Comment, domains.Comment),
		LikesHandler:      handler.NewLikesHandler(queries.Like),
	}
	// Init security handler
	securityHandler := other.NewSecurityHandler(services.TokenMaker)

	// Init ogen server
	server, err := ogen.NewServer(ogenHandler, securityHandler)
	if err != nil {
		return codeerr.NewInternal("newHTTPServer", err)
	}
	// Init middlewares
	loggingMiddleware := middleware.NewLoggingMiddleware()

	// Init chi router
	r := chi.NewRouter()
	// Basic CORS
	r.Use(cors.Handler(cors.Options{
		AllowOriginFunc: func(r *http.Request, origin string) bool { return true },
	}))
	// Host documentation (docs are located in docs/http/index.html)
	r.Mount("/docs/http/", http.StripPrefix("/docs/http/", http.FileServer(http.Dir(docsPath))))
	// Register middlewares
	httpHandler := loggingMiddleware(server)
	// Register routes
	r.Mount("/", httpHandler)

	// Start HTTP server
	service.Log.Info("Starting HTTP server on " + port)
	if err := http.ListenAndServe(port, r); err != nil {
		return codeerr.NewInternal("RunServer", err)
	}

	return nil
}
