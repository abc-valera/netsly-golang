package jsonApi

import (
	"embed"
	"io/fs"
	"net/http"
	"time"

	"github.com/abc-valera/netsly-golang/gen/ogen"
	"github.com/abc-valera/netsly-golang/internal/application"
	"github.com/abc-valera/netsly-golang/internal/domain/entity"
	"github.com/abc-valera/netsly-golang/internal/domain/global"
	"github.com/abc-valera/netsly-golang/internal/domain/service"
	"github.com/abc-valera/netsly-golang/internal/domain/util/coderr"
	"github.com/abc-valera/netsly-golang/internal/presentation/jsonApi/auth"
	"github.com/abc-valera/netsly-golang/internal/presentation/jsonApi/rest/handler"
	"github.com/abc-valera/netsly-golang/internal/presentation/jsonApi/rest/middleware"
	"github.com/abc-valera/netsly-golang/internal/presentation/jsonApi/rest/restErrors"
	"github.com/abc-valera/netsly-golang/internal/presentation/jsonApi/ws"
)

//go:embed static
var staticFiles embed.FS

func newHttpHandler(
	signKey string,
	accessTokenDuration time.Duration,
	refreshTokenDuration time.Duration,

	entities entity.Entities,
	_ service.Services,
	usecases application.Usecases,
) http.Handler {
	// Init router
	mux := http.NewServeMux()

	// TODO: register static fils AFTER MIDDLWARES
	// Static files
	mux.Handle("/", http.FileServer(http.FS(coderr.Must(fs.Sub(staticFiles, "static")))))

	// Init auth manager
	authManager := auth.NewManager(signKey, accessTokenDuration, refreshTokenDuration)

	// Ogen routes
	{
		// Init handlers (implementations of ogen interfaces)
		ogenHandler := &struct {
			restErrors.Handler
			handler.SignHandler
			handler.MeHandler
			handler.JokesHandler
		}{
			Handler:      restErrors.NewHandler(),
			SignHandler:  handler.NewSignHandler(authManager, usecases.SignUsecase),
			MeHandler:    handler.NewMeHandler(entities.User),
			JokesHandler: handler.NewJokesHandler(entities.Joke),
		}

		// Init security handler
		securityHandler := auth.NewHandler(authManager)

		// Init ogen server
		ogenServer := coderr.Must(
			ogen.NewServer(
				ogenHandler,
				securityHandler,
			),
		)

		// Ogen routes
		mux.Handle("/api/v1/", http.StripPrefix("/api/v1", ogenServer))
	}

	// WS routes
	{
		// Init Open Telemetry handlerTracer for WS handlers
		// handlerTracer := global.TraceProvider().Tracer("ws")

		// Init ws manager
		wsManager := ws.NewManager(
			authManager,
		)

		// WS routes
		mux.HandleFunc("/ws/v1/chat", wsManager.ServeWS)
	}

	// Init middlewares
	withTracer := middleware.NewTracer(global.Tracer())
	withRecoverer := middleware.NewRecoverer()
	withLogger := middleware.NewLogger()

	return withTracer(withRecoverer(withLogger(mux)))
}
