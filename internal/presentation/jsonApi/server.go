package jsonApi

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/abc-valera/netsly-api-golang/internal/application"
	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain"
	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
)

// NewServer returns HTTP server
func NewServer(
	port string,

	signKey string,
	accessTokenDuration time.Duration,
	refreshTokenDuration time.Duration,

	entities domain.Entities,
	services domain.Services,
	usecases application.Usecases,
) (
	serverStart func(),
	serverGracefulStop func(),
) {
	// Init server
	server := http.Server{
		Addr: port,
		Handler: newHttpHandler(
			signKey,
			accessTokenDuration,
			refreshTokenDuration,
			entities,
			services,
			usecases,
		),
	}

	return func() {
			global.Log().Info("jsonApi is running", "port", port)
			if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
				coderr.Fatal("jsonApi server error", err)
			}
		}, func() {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			if err := server.Shutdown(ctx); err != nil {
				coderr.Fatal("Shutdown server error", err)
			}
		}
}
