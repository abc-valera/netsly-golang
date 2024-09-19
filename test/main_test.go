package entity_test

import (
	"bytes"
	"context"
	"os"
	"testing"
	"text/template"

	"github.com/abc-valera/netsly-golang/internal/domain/entity"
	"github.com/abc-valera/netsly-golang/internal/domain/util/coderr"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/globals"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/services"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

// NewTest initializes the context, assertations and entities for the test and
// should be called at the beginning of each test function.
//
// Note, that each test function runs in a separate transaction.
var NewTest func(t *testing.T) (context.Context, *require.Assertions, entity.Entities)

// testsFilesFolder is the folder where the files generated during the tests are stored.
const testsFilesFolder = "./tmp"

// TestMain is the entry point for the test suite.
// It is responsible for setting up the infrastructure for the test environment.
// You can change the infrastructure you want to test.
func TestMain(m *testing.M) {
	// Set the Global Environment and init the Global Layer
	os.Setenv("APP_MODE", "production")
	os.Setenv("OTEL_TRACE_EXPORTER", "nop")
	os.Setenv("LOGGER_SERVICE", "nop")

	globals.New("testing")

	// Set the environment for the Persistence and Service layers.
	// For that read the environment variables from the .env file,
	// then create the directory for the temporary test infrastructure,
	// and finally initialize the Persistence and Service layers.

	// Create a text template with the content of the .env file
	var buf bytes.Buffer
	coderr.NoErr(coderr.Must(template.New(".env").ParseFiles(".env")).Execute(&buf, map[string]string{
		"testInfraFolder": testsFilesFolder,
	}))
	// Parse and set the environment variables
	for key, value := range coderr.Must(godotenv.Parse(&buf)) {
		os.Setenv(key, value)
	}

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

	// Init Services
	services := services.NewServices()

	// Init DB and check if Commands and Queries are valid
	db := persistences.NewDB()

	// Init Entities
	entities := entity.NewEntities(entity.NewDependency(db, services))

	// Seed the persistence layer
	coderr.NoErr(Seed(entities))

	// Initialize the newTest function.
	// Note, that we want to run each test in a separate transaction,
	// and rollback that transaction at the end of the test.
	// That is done to ensure that the tests are isolated from each other.
	NewTest = func(t *testing.T) (context.Context, *require.Assertions, entity.Entities) {
		tx, err := db.BeginTX(context.Background())
		require.NoError(t, err)

		t.Cleanup(func() {
			tx.Rollback()
		})

		return context.Background(), require.New(t), entity.NewEntities(entity.NewDependency(tx, services))
	}

	// Run the tests
	code := m.Run()

	// Run the cleanup
	os.RemoveAll(testsFilesFolder)

	// End the test suite
	os.Exit(code)
}
