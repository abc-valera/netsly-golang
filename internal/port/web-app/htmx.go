package webapp

import (
	"net/http"
	"os"

	"github.com/abc-valera/flugo-api-golang/internal/core/application"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/domain"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/query"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/service"
	"github.com/abc-valera/flugo-api-golang/internal/port/web-app/handler"
	"github.com/go-chi/chi/v5"
)

// NewServer returns HTTP server
func NewServer(
	port string,
	templatePath string,
	queries query.Queries,
	domains domain.Domains,
	services service.Services,
	usecases application.UseCases,
) http.Server {
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
