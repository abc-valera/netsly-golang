package webApp

import (
	"context"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/abc-valera/netsly-api-golang/internal/application"
	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain"
	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/presentation/webApp/handler"
	"github.com/go-chi/chi/v5"
)

// NewServer returns HTTP server
func NewServer(
	port string,
	templatePath string,
	staticPath string,

	services domain.Services,
	entities domain.Entities,
	usecases application.Usecases,
) (
	serverStart func(),
	serverGracefulStop func(),
) {
	// Init server
	server := http.Server{
		Addr:    port,
		Handler: NewHandler(templatePath, staticPath, entities, services, usecases),
	}

	return func() {
			global.Log().Info("webApp is running", "port", port)
			if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
				coderr.Fatal("webApp server error: ", err)
			}
		}, func() {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			if err := server.Shutdown(ctx); err != nil {
				coderr.Fatal("Shutdown server error: ", err)
			}
		}
}

func NewHandler(
	templatePath string,
	staticPath string,

	entities domain.Entities,
	services domain.Services,
	usecases application.Usecases,
) http.Handler {
	// Init handlers
	handlers := handler.NewHandlers(
		os.DirFS(templatePath),
		entities,
		usecases,
	)

	// Init router
	r := chi.NewRouter()
	initRoutes(r, staticPath, services, handlers)

	return r
}
