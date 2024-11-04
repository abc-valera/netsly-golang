package restHandler

import (
	"embed"
	"net/http"

	"github.com/abc-valera/netsly-golang/gen/ogen"
	"github.com/abc-valera/netsly-golang/internal/application"
	"github.com/abc-valera/netsly-golang/internal/domain/entity"
	"github.com/abc-valera/netsly-golang/internal/domain/service"
	"github.com/abc-valera/netsly-golang/internal/domain/util/coderr"
	"github.com/abc-valera/netsly-golang/internal/presentation/jsonApi/auth"
	"github.com/abc-valera/netsly-golang/internal/presentation/jsonApi/rest/restErrors"
)

type Handlers struct {
	Debug Debug
	Docs  Docs
	Ogen  http.Handler
}

type ogenHandlers struct {
	Sign
	Me
	Jokes

	restErrors.Handler
}

// Make sure that Handlers implements ogen.Handler
var _ ogen.Handler = ogenHandlers{}

func New(
	staticFiles embed.FS,
	openapiFile []byte,

	authManager auth.Manager,

	entities entity.Entities,
	services service.Services,
	usecases application.Usecases,
) Handlers {
	ogenHandlers := ogenHandlers{
		Sign:  newSign(authManager, usecases.SignUsecase),
		Me:    newMe(entities.User),
		Jokes: newJokes(entities.Joke),

		Handler: restErrors.NewHandler(),
	}

	return Handlers{
		Debug: NewDebug(),
		Docs:  NewDocs(openapiFile),
		Ogen:  coderr.Must(ogen.NewServer(ogenHandlers, auth.NewHandler(authManager))),
	}
}
