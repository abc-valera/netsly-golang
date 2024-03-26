package webApp

import (
	"context"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/abc-valera/netsly-api-golang/pkg/application"
	"github.com/abc-valera/netsly-api-golang/pkg/core/global"
	"github.com/abc-valera/netsly-api-golang/pkg/domain"
	"github.com/abc-valera/netsly-api-golang/pkg/presentation/webApp/handler"
	"github.com/go-chi/chi/v5"
)

// NewServer returns HTTP server
func NewServer(
	port string,
	templatePath string,
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
		Handler: NewHandler(templatePath, staticPath, queries, entities, services, usecases),
	}

	return func() {
			global.Log().Info("webApp is running", "port", port)
			if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
				global.Log().Fatal("webApp server error: ", err)
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
	templatePath string,
	staticPath string,

	queries domain.Queries,
	entities domain.Entities,
	services domain.Services,
	usecases application.UseCases,
) http.Handler {
	// Init handlers
	handlers := handler.NewHandlers(
		os.DirFS(templatePath),
		queries,
		entities,
		usecases,
	)

	// Init router
	r := chi.NewRouter()
	initRoutes(r, staticPath, services, handlers)

	return r
}
