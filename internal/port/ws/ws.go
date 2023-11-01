package ws

import (
	"net/http"

	"github.com/abc-valera/flugo-api-golang/internal/core/application"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/codeerr"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/service"
	"github.com/abc-valera/flugo-api-golang/internal/port/ws/manager"
	"github.com/go-chi/chi/v5"
)

func RunServer(
	port string,
	repos repository.Repositories,
	services service.Services,
	usecases application.UseCases,
) error {
	manager := manager.NewManager()

	r := chi.NewRouter()
	r.HandleFunc("/ws", manager.ServeWS)

	services.Logger.Info("Starting WebSocket server on " + port)
	if err := http.ListenAndServe(port, r); err != nil {
		return codeerr.NewInternal("RunServer", err)
	}

	return nil
}
