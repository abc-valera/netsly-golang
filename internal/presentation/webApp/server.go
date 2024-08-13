package webApp

import (
	"context"
	"embed"
	"errors"
	"io/fs"
	"net/http"
	"os"
	"time"

	"github.com/abc-valera/netsly-golang/internal/application"
	"github.com/abc-valera/netsly-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-golang/internal/core/global"
	"github.com/abc-valera/netsly-golang/internal/domain"
	"github.com/abc-valera/netsly-golang/internal/presentation/webApp/handler"
	"github.com/go-chi/chi/v5"
)

//go:embed template
var templateEmbedFS embed.FS

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
	var templateFS fs.FS
	if global.IsProduction() {
		templateFS = templateEmbedFS
	} else {
		templateFS = os.DirFS(templatePath)
	}

	// Init handlers
	handlers := handler.NewHandlers(
		templateFS,
		entities,
		usecases,
	)

	// Init router
	r := chi.NewRouter()
	newRouter(r, services, handlers)

	// Init server
	server := http.Server{
		Addr:    port,
		Handler: r,
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
