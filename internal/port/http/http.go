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
		ErrorHandler:      other.NewErrorHandler(services.Logger),
		SignHandler:       sign.NewSignHandler(repos.UserRepo, usecases.SignUseCase),
		MeHandler:         me.NewMeHandler(repos.UserRepo),
		MeJokesHandler:    me.NewMeJokesHandler(repos.JokeRepo),
		MeCommentsHandler: me.NewMeCommentsHandler(repos.CommentRepo),
		MeLikesHandler:    me.NewMeLikesHandler(repos.LikeRepo),
		CommentsHandler:   comments.NewCommentsHandler(repos.CommentRepo),
		LikesHandler:      likes.NewLikesHandler(repos.LikeRepo),
	}
	// Init security handler
	securityHandler := other.NewSecurityHandler(services.TokenMaker)

	// Init ogen server
	server, err := ogen.NewServer(ogenHandler, securityHandler)
	if err != nil {
		return codeerr.NewInternal("newHTTPServer", err)
	}
	// Init middlewares
	loggingMiddleware := middleware.NewLoggingMiddleware(services.Logger)

	// Init chi router
	r := chi.NewRouter()
	// Host documentation (docs are located in docs/http/index.html)
	r.Mount("/docs/http/", http.StripPrefix("/docs/http/", http.FileServer(http.Dir(docsPath))))
	// Register middlewares
	httpHandler := loggingMiddleware(server)
	// Register routes
	r.Mount("/", httpHandler)

	// Start HTTP server
	services.Logger.Info("Starting HTTP server on " + port)
	if err := http.ListenAndServe(port, r); err != nil {
		return codeerr.NewInternal("RunServer", err)
	}

	return nil
}
