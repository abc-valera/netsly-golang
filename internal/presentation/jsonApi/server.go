package jsonApi

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/abc-valera/netsly-api-golang/gen/ogen"
	"github.com/abc-valera/netsly-api-golang/internal/application"
	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/core/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain"
	"github.com/abc-valera/netsly-api-golang/internal/presentation/jsonApi/rest/handler"
	"github.com/abc-valera/netsly-api-golang/internal/presentation/jsonApi/rest/middleware"
	"github.com/abc-valera/netsly-api-golang/internal/presentation/jsonApi/ws"
	"github.com/go-chi/chi/v5"
)

// NewServer returns HTTP server
func NewServer(
	port string,
	staticPath string,

	services domain.Services,
	entities domain.Entities,
	usecases application.UseCases,
) (
	serverStart func(),
	serverGracefulStop func(),
) {
	// Init server
	server := http.Server{
		Addr:    port,
		Handler: NewHandler(staticPath, entities, services, usecases),
	}

	return func() {
			global.Log().Info("jsonApi is running", "port", port)
			if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
				coderr.Fatal("jsonApi server error", err)
			}
		}, func() {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			if err := server.Shutdown(ctx); err != nil {
				coderr.Fatal("Shutdown server error", err)
			}
		}
}

func NewHandler(
	staticPath string,

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
			MeHandler:         handler.NewMeHandler(entities.User),
			MeJokesHandler:    handler.NewMeJokesHandler(entities.Joke),
			MeCommentsHandler: handler.NewMeCommentsHandler(entities.Comment),
			MeLikesHandler:    handler.NewMeLikesHandler(entities.Like),
			MeRooms:           handler.NewMeRooms(entities.Room, entities.RoomMember),
			CommentsHandler:   handler.NewCommentsHandler(entities.Comment),
			LikesHandler:      handler.NewLikesHandler(entities.Like),
			Rooms:             handler.NewRooms(entities.RoomMessage),
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
