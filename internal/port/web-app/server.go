package webapp

import (
	"log"
	"net/http"
	"os"

	"github.com/abc-valera/flugo-api-golang/internal/core/application"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain"
	"github.com/abc-valera/flugo-api-golang/internal/port/web-app/handler"
	"github.com/go-chi/chi/v5"
)

// NewServer returns HTTP server
func NewServer(
	port string,
	templatePath string,
	queries domain.Queries,
	domains domain.Domains,
	services domain.Services,
	usecases application.UseCases,
) http.Server {
	if port == "" {
		log.Fatal("port is not set")
	}
	if templatePath == "" {
		log.Fatal("template path is not set")
	}

	// Init handlers
	handlers := handler.NewHandlers(
		os.DirFS(templatePath),
		usecases,
	)

	// Init router
	r := chi.NewRouter()
	initRoutesMiddlewares(r, handlers)

	// Init server
	return http.Server{
		Addr:    port,
		Handler: r,
	}
}
