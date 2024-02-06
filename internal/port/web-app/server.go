package webapp

import (
	"log"
	"net/http"
	"os"

	"github.com/abc-valera/netsly-api-golang/internal/application"
	"github.com/abc-valera/netsly-api-golang/internal/domain"
	"github.com/abc-valera/netsly-api-golang/internal/port/web-app/handler"
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
		queries,
		entities,
		usecases,
	)

	// Init router
	r := chi.NewRouter()
	initRoutes(r, staticPath, services, handlers)

	// Init server
	return http.Server{
		Addr:    port,
		Handler: r,
	}
}
