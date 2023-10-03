package http

import (
	"net/http"

	"github.com/abc-valera/flugo-api-golang/internal/application"
	"github.com/abc-valera/flugo-api-golang/internal/domain/codeerr"
	"github.com/abc-valera/flugo-api-golang/internal/domain/repository"
	"github.com/abc-valera/flugo-api-golang/internal/domain/service"
	"github.com/abc-valera/flugo-api-golang/internal/infrastructure/port/http/handlers"
	"github.com/abc-valera/flugo-api-golang/internal/infrastructure/port/http/middlewares"
	"github.com/abc-valera/flugo-api-golang/tools/ogen"
)

type serverHandler struct {
	handlers.SignHandler
	handlers.ErrorHandler
}

func newServerHandler(
	repos repository.Repositories,
	services service.Services,
	usecases application.UseCases,
) ogen.Handler {
	return &serverHandler{
		SignHandler: handlers.NewSignHandler(
			repos.UserRepo,
			usecases.SignUseCase,
		),
		ErrorHandler: handlers.NewErrorHandler(services.Logger),
	}
}

func RunServer(
	port string,
	repos repository.Repositories,
	services service.Services,
	usecases application.UseCases,
) error {
	// Init handler and server
	handler := newServerHandler(repos, services, usecases)
	server, err := ogen.NewServer(handler)
	if err != nil {
		return codeerr.NewInternal("RunServer", err)
	}

	// Init middlewares
	loggingMiddleware := middlewares.NewLoggingMiddleware(services.Logger)
	loggingHandler := loggingMiddleware(server)

	// Start HTTP server
	services.Logger.Info("Starting HTTP server on " + port)
	if err := http.ListenAndServe(port, loggingHandler); err != nil {
		return codeerr.NewInternal("RunServer", err)
	}

	return nil
}
