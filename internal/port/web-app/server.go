package webapp

import (
	"context"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/abc-valera/netsly-api-golang/internal/application"
	"github.com/abc-valera/netsly-api-golang/internal/domain"
	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/port/web-app/handler"
	"github.com/go-chi/chi/v5"
)

var (
	port         = os.Getenv("WEB_APP_PORT")
	templatePath = os.Getenv("WEB_APP_TEMPLATE_PATH")
	staticPath   = os.Getenv("WEB_APP_STATIC_PATH")
)

// NewServer returns HTTP server
func NewServer(
	queries domain.Queries,
	entities domain.Entities,
	services domain.Services,
	usecases application.UseCases,
) (
	serverStart func(),
	serverGracefulStop func(),
) {
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

	// Init server
	server := http.Server{
		Addr:    port,
		Handler: r,
	}

	return func() {
			global.Log().Info("web-app is running on port ", "port", port)
			if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
				global.Log().Fatal("web-app server error: ", err)
			}
		}, func() {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			if err := server.Shutdown(ctx); err != nil {
				global.Log().Fatal("Shutdown server error: ", err)
			}
		}
}
