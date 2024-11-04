package jsonApi

import (
	"embed"
	"net/http"
	"time"

	"github.com/abc-valera/netsly-golang/internal/application"
	"github.com/abc-valera/netsly-golang/internal/domain/entity"
	"github.com/abc-valera/netsly-golang/internal/domain/global"
	"github.com/abc-valera/netsly-golang/internal/domain/service"
	"github.com/abc-valera/netsly-golang/internal/presentation/jsonApi/auth"
	"github.com/abc-valera/netsly-golang/internal/presentation/jsonApi/rest/middleware"
	"github.com/abc-valera/netsly-golang/internal/presentation/jsonApi/rest/restHandler"
	"github.com/abc-valera/netsly-golang/internal/presentation/jsonApi/ws"
)

func NewHandler(
	staticFiles embed.FS,
	openapiFile []byte,

	signKey string,
	accessTokenDuration time.Duration,
	refreshTokenDuration time.Duration,

	entities entity.Entities,
	services service.Services,
	usecases application.Usecases,
) http.Handler {
	// Init auth manager
	authManager := auth.NewManager(signKey, accessTokenDuration, refreshTokenDuration)

	// REST ROUTER
	restMux := http.NewServeMux()
	{
		handlers := restHandler.New(
			staticFiles,
			openapiFile,
			authManager,
			entities,
			services,
			usecases,
		)

		restMux.HandleFunc("/ping", handlers.Debug.Ping)
		restMux.HandleFunc("/openapi", handlers.Docs.OpenApiFile)
		restMux.HandleFunc("/docs", handlers.Docs.ScalarDocs)

		restMux.Handle("/", handlers.Ogen)
	}

	// WEBSOCKET ROUTER
	wsMux := http.NewServeMux()
	{
		// TODO:
		// Init Open Telemetry handlerTracer for WS handlers
		// handlerTracer := global.TraceProvider().Tracer("ws")

		// Init ws manager
		wsManager := ws.NewManager(
			authManager,
		)

		// WS routes
		wsMux.HandleFunc("/chat", wsManager.ServeWS)
	}

	// GENERAL ROUTER
	generalMux := http.NewServeMux()
	{
		generalMux.Handle("/v1/", http.StripPrefix("/v1", restMux))
		generalMux.Handle("/v1/ws/", http.StripPrefix("/v1/ws", wsMux))
	}

	// MIDDLEWARES
	withTracer := middleware.NewTracer(global.Tracer())
	withRecoverer := middleware.NewRecoverer()
	withLogger := middleware.NewLogger()

	return withTracer(withRecoverer(withLogger(generalMux)))
}
