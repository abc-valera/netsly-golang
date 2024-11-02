package jsonApi

import (
	"net/http"
	"time"

	"github.com/abc-valera/netsly-golang/internal/application"
	"github.com/abc-valera/netsly-golang/internal/domain/entity"
	"github.com/abc-valera/netsly-golang/internal/domain/service"
)

// NewServer returns HTTP server
func NewServer(
	signKey string,
	accessTokenDuration time.Duration,
	refreshTokenDuration time.Duration,

	entities entity.Entities,
	services service.Services,
	usecases application.Usecases,
) http.Handler {
	return newHttpHandler(
		signKey,
		accessTokenDuration,
		refreshTokenDuration,
		entities,
		services,
		usecases,
	)
}
