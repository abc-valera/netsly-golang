package jsonrestapi

import (
	"log"
	"net/http"

	"github.com/abc-valera/netsly-api-golang/gen/ogen"
	"github.com/abc-valera/netsly-api-golang/internal/application"
	"github.com/abc-valera/netsly-api-golang/internal/domain"
	"github.com/abc-valera/netsly-api-golang/internal/domain/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/port/json-rest-api/handler"
	"github.com/abc-valera/netsly-api-golang/internal/port/json-rest-api/middleware"
	"github.com/go-chi/chi/v5"
)

// NewServer returns HTTP server
func NewServer(
	port string,
	staticPath string,
	queries domain.Queries,
	entities domain.Entities,
	services domain.Services,
	usecases application.UseCases,
) http.Server {
	if port == "" {
		log.Fatal("port is not set")
	}
	if staticPath == "" {
		log.Fatal("static path is not set")
	}

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
		MeHandler:         handler.NewMeHandler(queries.User, entities.User),
		MeJokesHandler:    handler.NewMeJokesHandler(queries.Joke, entities.Joke),
		MeCommentsHandler: handler.NewMeCommentsHandler(queries.Comment, entities.Comment),
		MeLikesHandler:    handler.NewMeLikesHandler(queries.Like, entities.Like),
		CommentsHandler:   handler.NewCommentsHandler(queries.Comment, entities.Comment),
		LikesHandler:      handler.NewLikesHandler(queries.Like),
	}
	// Init security handler
	securityHandler := handler.NewSecurityHandler(services.TokenMaker)

	// Init ogen server
	server := coderr.Must[*ogen.Server](ogen.NewServer(ogenHandler, securityHandler))

	// Init chi router
	r := chi.NewRouter()

	// Regiter middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Host static files (docs are hosted in /static/docs/index.html)
	r.Mount("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(staticPath))))

	// Register routes
	r.Mount("/", server)

	// Init server
	return http.Server{
		Addr:    port,
		Handler: r,
	}
}
