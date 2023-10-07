package http

import (
	"net/http"

	"github.com/abc-valera/flugo-api-golang/gen/ogen"
	"github.com/abc-valera/flugo-api-golang/internal/application"
	"github.com/abc-valera/flugo-api-golang/internal/domain/codeerr"
	"github.com/abc-valera/flugo-api-golang/internal/domain/repository"
	"github.com/abc-valera/flugo-api-golang/internal/domain/service"
	"github.com/abc-valera/flugo-api-golang/internal/infrastructure/port/http/handler"
	"github.com/abc-valera/flugo-api-golang/internal/infrastructure/port/http/middlewares"
	"github.com/go-chi/chi/v5"
)

func RunServer(
	port string,
	repos repository.Repositories,
	services service.Services,
	usecases application.UseCases,
) error {
	// Init handlers (ogenHandler implements ogen.Server interface)
	ogenHandler := &struct {
		handler.SignHandler
		handler.ErrorHandler
	}{
		SignHandler:  handler.NewSignHandler(repos.UserRepo, usecases.SignUseCase),
		ErrorHandler: handler.NewErrorHandler(services.Logger),
	}

	// Init ogen server
	server, err := ogen.NewServer(ogenHandler)
	if err != nil {
		return codeerr.NewInternal("newHTTPServer", err)
	}
	// Init middlewares
	loggingMiddleware := middlewares.NewLoggingMiddleware(services.Logger)

	// Init chi router
	r := chi.NewRouter()
	// Host documentation
	r.Mount("/docs", http.StripPrefix("/docs/", http.FileServer(http.Dir("./docs/http"))))
	// Register middlewares
	httpHandler := loggingMiddleware(server)
	// Register routes
	r.Mount("/", httpHandler)

	// Start HTTP server
	services.Logger.Info("Starting HTTP server on " + port)
	if err := http.ListenAndServe(port, r); err != nil {
		return codeerr.NewInternal("RunServer", err)
	}

	return nil
}
