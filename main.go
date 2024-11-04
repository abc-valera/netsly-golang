package main

import (
	"embed"

	"github.com/abc-valera/netsly-golang/internal/application"
	"github.com/abc-valera/netsly-golang/internal/domain/entity"
	"github.com/abc-valera/netsly-golang/internal/domain/util/coderr"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/globals"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/services"
	"github.com/abc-valera/netsly-golang/internal/presentation"
)

// Embedded files

//go:embed internal/presentation/jsonApi/static
var jsonApiStaticFiles embed.FS

//go:embed gen/openapi/openapi.yaml
var jsonApiOpenapiFile []byte

func main() {
	// Init Globals
	globals.New()

	// Init Services
	services := services.NewServices()

	// Init DB
	db := persistences.NewDB()

	// Init Entities
	entities := entity.NewEntities(entity.NewDependency(db, services))

	// Init usecases
	usecases := application.NewUsecases(application.NewDependency(db, services))

	// Check if all layers are initialized correctly
	coderr.NoErr(coderr.CheckIfStructHasEmptyFields(db.Commands()))
	coderr.NoErr(coderr.CheckIfStructHasEmptyFields(db.Queries()))
	coderr.NoErr(coderr.CheckIfStructHasEmptyFields(services))
	coderr.NoErr(coderr.CheckIfStructHasEmptyFields(entities))
	coderr.NoErr(coderr.CheckIfStructHasEmptyFields(usecases))

	// Start the server
	presentation.StartServer(
		jsonApiStaticFiles,
		jsonApiOpenapiFile,

		services,
		entities,
		usecases,
	)
}
