package webApp

import (
	"embed"
	"io/fs"
	"net/http"
	"os"

	"github.com/abc-valera/netsly-golang/internal/application"
	"github.com/abc-valera/netsly-golang/internal/domain/entity"
	"github.com/abc-valera/netsly-golang/internal/domain/global"
	"github.com/abc-valera/netsly-golang/internal/domain/service"
	"github.com/abc-valera/netsly-golang/internal/presentation/webApp/handler"
	"github.com/go-chi/chi/v5"
)

//go:embed template
var templateEmbedFS embed.FS

// NewServer returns HTTP server
func NewServer(
	templatePath string,

	services service.Services,
	entities entity.Entities,
	usecases application.Usecases,
) http.Handler {
	var templateFS fs.FS
	if global.IsProduction() {
		templateFS = templateEmbedFS
	} else {
		templateFS = os.DirFS(templatePath)
	}

	// Init handlers
	handlers := handler.NewHandlers(
		templateFS,
		entities,
		usecases,
	)

	// Init router
	r := chi.NewRouter()
	newRouter(r, services, handlers)

	return r
}
