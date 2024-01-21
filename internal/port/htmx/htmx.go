package htmx

import (
	"net/http"
	"os"

	"github.com/abc-valera/flugo-api-golang/internal/core/application"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/coderr"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/domain"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/query"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/service"
	"github.com/abc-valera/flugo-api-golang/internal/port/htmx/handler"
	"github.com/go-chi/chi/v5"
)

func RunServer(
	port string,
	templatePath string,
	queries query.Queries,
	domains domain.Domains,
	services service.Services,
	usecases application.UseCases,
) error {
	// Init handlers
	handlers, err := handler.NewHandlers(
		os.DirFS(templatePath),
		usecases,
	)
	if err != nil {
		return err
	}

	// Init router
	r := chi.NewRouter()

	initRoutesMiddlewares(r, handlers)

	// Start server
	service.Log.Info("Starting HTMX server on " + port)
	if err := http.ListenAndServe(port, r); err != nil {
		return coderr.NewInternal(err)
	}

	return nil
}
