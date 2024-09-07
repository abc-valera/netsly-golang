package entity_test

import (
	"os"
	"testing"

	"github.com/abc-valera/netsly-golang/internal/domain/entity"
	"github.com/abc-valera/netsly-golang/internal/domain/util/coderr"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/globals"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/services"
	"github.com/joho/godotenv"
)

var entities entity.Entities

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
	coderr.NoErr(godotenv.Load(".env.test"))

	// Set the Mode
	// Note, that we always want to run the tests in production mode
	os.Setenv("APP_MODE", "production")

	// Set the folder paths for the infrastructure.
	os.Setenv("GORM_SQLITE_FOLDER_PATH", testsFilesFolder+"/gormSqlite")
	os.Setenv("BUN_SQLITE_FOLDER_PATH", testsFilesFolder+"/bunSqlite")
	os.Setenv("LOCAL_FILE_SAVER_FOLDER_PATH", testsFilesFolder+"/localFileSaver")
	os.Setenv("SLOG_LOGGER_FOLDER_PATH", testsFilesFolder+"/slogLogger")

	// Init Global
	globals.New("testing")

	// Init Services
	services := services.NewServices()

	// Init DB and check if Commands and Queries are valid
	db := persistences.NewDB()

	// Init Entities
	entities = entity.NewEntities(entity.NewDependency(db, services))

	// Seed the persistence layer
	coderr.NoErr(Seed(entities))

	// Run the tests
	code := m.Run()

	// Run the cleanup
	os.RemoveAll(testsFilesFolder)

	// End the test suite
	os.Exit(code)
}
