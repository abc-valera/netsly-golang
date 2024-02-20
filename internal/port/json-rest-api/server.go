package jsonrestapi

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/abc-valera/netsly-api-golang/gen/ogen"
	"github.com/abc-valera/netsly-api-golang/internal/application"
	"github.com/abc-valera/netsly-api-golang/internal/domain"
	"github.com/abc-valera/netsly-api-golang/internal/domain/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
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
) (
	serverStart func(),
	serverGracefulStop func(),
) {
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
	ogenServer := coderr.Must[*ogen.Server](ogen.NewServer(ogenHandler, securityHandler))

	// Init chi router
	r := chi.NewRouter()

	// Regiter middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Host static files (docs are hosted in /static/docs/index.html)
	r.Mount("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(staticPath))))

	// Register routes
	r.Mount("/", ogenServer)

	// Init server
	server := http.Server{
		Addr:    port,
		Handler: r,
	}

	return func() {
			global.Log().Info("json-rest-api is running", "port", port)
			if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
				global.Log().Fatal("json-rest-api server error: ", err)
			}
		}, func() {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			if err := server.Shutdown(ctx); err != nil {
				global.Log().Fatal("Shutdown server error: ", err)
			}
		}
}
