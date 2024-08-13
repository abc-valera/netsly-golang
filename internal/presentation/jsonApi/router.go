package jsonApi

import (
	"embed"
	"io/fs"
	"net/http"
	"time"

	"github.com/abc-valera/netsly-golang/gen/ogen"
	"github.com/abc-valera/netsly-golang/internal/application"
	"github.com/abc-valera/netsly-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-golang/internal/core/global"
	"github.com/abc-valera/netsly-golang/internal/domain"
	"github.com/abc-valera/netsly-golang/internal/presentation/jsonApi/auth"
	"github.com/abc-valera/netsly-golang/internal/presentation/jsonApi/rest/errutil"
	"github.com/abc-valera/netsly-golang/internal/presentation/jsonApi/rest/handler"
	"github.com/abc-valera/netsly-golang/internal/presentation/jsonApi/rest/middleware"
	"github.com/abc-valera/netsly-golang/internal/presentation/jsonApi/ws"
)

//go:embed static
var staticFiles embed.FS

func newHttpHandler(
	signKey string,
	accessTokenDuration time.Duration,
	refreshTokenDuration time.Duration,

	entities domain.Entities,
	services domain.Services,
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
			errutil.Handler
			handler.SignHandler
			handler.MeHandler
			handler.MeJokesHandler
			handler.MeCommentsHandler
			handler.MeLikesHandler
			handler.CommentsHandler
			handler.LikesHandler
			handler.MeRooms
			handler.Rooms
		}{
			Handler:           errutil.NewHandler(),
			SignHandler:       handler.NewSignHandler(authManager, usecases.SignUsecase),
			MeHandler:         handler.NewMeHandler(entities.User),
			MeJokesHandler:    handler.NewMeJokesHandler(entities.Joke),
			MeCommentsHandler: handler.NewMeCommentsHandler(entities.Comment),
			MeLikesHandler:    handler.NewMeLikesHandler(entities.Like),
			MeRooms:           handler.NewMeRooms(entities.Room, entities.RoomMember),
			CommentsHandler:   handler.NewCommentsHandler(entities.Comment),
			LikesHandler:      handler.NewLikesHandler(entities.Like),
			Rooms:             handler.NewRooms(entities.RoomMessage),
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
			services,
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
