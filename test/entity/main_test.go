package entity_test

import (
	"os"
	"testing"

	"github.com/abc-valera/netsly-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-golang/internal/core/global"
	"github.com/abc-valera/netsly-golang/internal/domain"
	"github.com/abc-valera/netsly-golang/internal/infrastructure"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/services/logger/loggerNop"
	"github.com/joho/godotenv"
	"go.opentelemetry.io/otel/trace/noop"
)

var entities domain.Entities

// testsFilesFolder is the folder where the files generated during the tests are stored.
const testsFilesFolder = "./tmp"

// TestMain is the entry point for the test suite.
// It is responsible for setting up the infrastructure for the test environment.
// You can change the infrastructure you want to test.
func TestMain(m *testing.M) {
	// Create a directory for temporary test files.
	// If the directory already exists, delete it and create a new one.
	if err := os.Mkdir(testsFilesFolder, 0o755); err != nil {
		if os.IsExist(err) {
			coderr.NoErr(os.RemoveAll(testsFilesFolder))
			coderr.NoErr(os.Mkdir(testsFilesFolder, 0o755))
		} else {
			coderr.Fatal(err)
		}
	}

	// Load the environment from the test.env file.
	coderr.NoErr(godotenv.Load("../env/test.env"))

	// Set the Mode
	os.Setenv("MODE", "production")

	// Set the folder paths for the infrastructure.
	os.Setenv("GORM_SQLITE_FOLDER_PATH", testsFilesFolder+"/gormSqlite")
	os.Setenv("BOILER_SQLITE_FOLDER_PATH", testsFilesFolder+"/boilerSqlite")
	os.Setenv("LOCAL_FILE_SAVER_FOLDER_PATH", testsFilesFolder+"/localFileSaver")
	os.Setenv("SLOG_LOGGER_FOLDER_PATH", testsFilesFolder+"/slogLogger")

	// Init Mode first
	global.InitMode()

	// Init services
	services := infrastructure.NewServices()

	// Init other global variables
	global.InitLog(loggerNop.New())
	global.InitTracer(noop.NewTracerProvider().Tracer("testing"))

	// Init persistence
	commands, queries, commandTransactor, _ := infrastructure.NewPersistences(services)

	// Init entities
	entities = domain.NewEntities(commands, commandTransactor, queries, services)

	// Seed the persistence layer
	coderr.NoErr(Seed(entities))

	// Run the tests
	code := m.Run()

	// Run the cleanup
	os.RemoveAll(testsFilesFolder)

	// End the test suite
	os.Exit(code)
}
