package jsonApi

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/abc-valera/netsly-api-golang/gen/ogen"
	"github.com/abc-valera/netsly-api-golang/pkg/application"
	"github.com/abc-valera/netsly-api-golang/pkg/core/coderr"
	"github.com/abc-valera/netsly-api-golang/pkg/core/global"
	"github.com/abc-valera/netsly-api-golang/pkg/domain"
	"github.com/abc-valera/netsly-api-golang/pkg/presentation/jsonApi/rest/handler"
	"github.com/abc-valera/netsly-api-golang/pkg/presentation/jsonApi/rest/middleware"
	"github.com/abc-valera/netsly-api-golang/pkg/presentation/jsonApi/ws"
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
	// Init server
	server := http.Server{
		Addr:    port,
		Handler: NewHandler(staticPath, queries, entities, services, usecases),
	}

	return func() {
			global.Log().Info("jsonApi is running", "port", port)
			if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
				global.Log().Fatal("jsonApi server error: ", err)
			}
		}, func() {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			if err := server.Shutdown(ctx); err != nil {
				global.Log().Fatal("Shutdown server error: ", err)
			}
		}
}

func NewHandler(
	staticPath string,

	queries domain.Queries,
	entities domain.Entities,
	services domain.Services,
	usecases application.UseCases,
) http.Handler {
	// Init chi router
	r := chi.NewRouter()

	// Static files
	r.Handle("/*", http.FileServer(http.Dir(staticPath)))

	// Ogen routes
	r.Route("/api/v1", func(r chi.Router) {
		// Regiter middlewares
		r.Use(middleware.Recoverer)
		r.Use(middleware.Logger)

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
			handler.MeRooms
			handler.Rooms
		}{
			ErrorHandler:      handler.NewErrorHandler(),
			SignHandler:       handler.NewSignHandler(usecases.SignUseCase),
			MeHandler:         handler.NewMeHandler(queries.User, entities.User),
			MeJokesHandler:    handler.NewMeJokesHandler(queries.Joke, entities.Joke),
			MeCommentsHandler: handler.NewMeCommentsHandler(queries.Comment, entities.Comment),
			MeLikesHandler:    handler.NewMeLikesHandler(queries.Like, entities.Like),
			MeRooms:           handler.NewMeRooms(queries.Room, entities.Room, entities.RoomMember),
			CommentsHandler:   handler.NewCommentsHandler(queries.Comment, entities.Comment),
			LikesHandler:      handler.NewLikesHandler(queries.Like),
			Rooms:             handler.NewRooms(queries.RoomMessage),
		}
		// Init security handler
		securityHandler := handler.NewSecurityHandler(services.TokenMaker)
		// Init ogen server
		ogenServer := coderr.Must(ogen.NewServer(ogenHandler, securityHandler))
		// Register ogen routes
		r.Mount("/", http.StripPrefix("/api/v1", ogenServer))
	})

	// WS routes
	r.Route("/ws/v1", func(r chi.Router) {
		// Init ws manager
		wsManager := ws.NewManager(
			services,
		)
		// Register ws routes
		r.HandleFunc("/chat", wsManager.ServeWS)
	})

	return r
}
