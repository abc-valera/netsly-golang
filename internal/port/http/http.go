package http

import (
	"net/http"

	"github.com/abc-valera/flugo-api-golang/gen/ogen"
	"github.com/abc-valera/flugo-api-golang/internal/core/application"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/coderr"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/domain"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/query"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/service"
	"github.com/abc-valera/flugo-api-golang/internal/port/http/handler"
	"github.com/abc-valera/flugo-api-golang/internal/port/http/middleware"
	"github.com/go-chi/chi/v5"
)

// RunServer runs HTTP server
func RunServer(
	port string,
	docsPath string,
	queries query.Queries,
	domains domain.Domains,
	services service.Services,
	usecases application.UseCases,
) error {
	// Init handlers (ogenHandler implements ogen.Server interface)
	ogenHandler := &struct {
		handler.ErrorHandler
		handler.SignHandler
		handler.MeHandler
		handler.MeJokesHandler
		handler.MeCommentsHandler
		handler.MeLikesHandler
		handler.CommentsHandler
		handler.LikesHandler
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
	securityHandler := handler.NewSecurityHandler(services.TokenMaker)

	// Init ogen server
	server, err := ogen.NewServer(ogenHandler, securityHandler)
	if err != nil {
		return coderr.NewInternal(err)
	}
	// Init middlewares
	loggingMiddleware := middleware.NewLoggingMiddleware()

	// Init chi router
	r := chi.NewRouter()
	// Host documentation (docs are located in docs/http/index.html)
	r.Mount("/docs/", http.StripPrefix("/docs/", http.FileServer(http.Dir(docsPath))))
	// Register middlewares
	httpHandler := loggingMiddleware(server)
	// Register routes
	r.Mount("/", httpHandler)

	// Start HTTP server
	service.Log.Info("Starting HTTP server on " + port)
	if err := http.ListenAndServe(port, r); err != nil {
		return coderr.NewInternal(err)
	}

	return nil
}
