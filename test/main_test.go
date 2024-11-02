package test

import (
	"context"
	"os"
	"testing"

	"github.com/abc-valera/netsly-golang/internal/domain/entity"
	"github.com/abc-valera/netsly-golang/internal/domain/global"
	"github.com/abc-valera/netsly-golang/internal/domain/util/coderr"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/globals/logger/loggerNop"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/services"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/trace/noop"
)

// NewTest initializes the test environment for a single test.
//
// The NewTest function should be called at the beginning of every test function,
// shouldn't be used unside the subtests.
//
// Note, that we want to run each test in a separate transaction,
// and rollback that transaction at the end of the test.
// That is done to ensure that the tests are isolated from each other.
var NewTest func(t *testing.T) (context.Context, *require.Assertions, entity.Entities)

// TestMain is the entry point for the test suite.
// It is responsible for setting up the infrastructure for the test environment.
// You can change the infrastructure you want to test.
func TestMain(m *testing.M) {
	// Init global variables. Note, that generally we don't want to mock it,
	// just use noop variants and make sure it's not null.
	global.Init(
		global.ModeProduction,
		noop.NewTracerProvider().Tracer("testing"),
		loggerNop.New(),
		"example.com",
		"web.",
		"api.",
	)

	// Get the path to the directory where the test files will be stored
	testFilesFolderPath := os.Getenv("TEST_FILES_FOLDER_PATH")

	// Create a directory for temporary test files.
	// If the directory already exists, delete it and create a new one.
	if err := os.Mkdir(testFilesFolderPath, 0o755); err != nil {
		if os.IsExist(err) {
			coderr.NoErr(os.RemoveAll(testFilesFolderPath))
			coderr.NoErr(os.Mkdir(testFilesFolderPath, 0o755))
		} else {
			coderr.Fatal(err)
		}
	}
	// Defer the cleanup
	defer os.RemoveAll(testFilesFolderPath)

	// Init Services
	services := services.NewServices()

	// Init DB and check if Commands and Queries are valid
	db := persistences.NewDB()

	// Initialize the newTest function.
	//
	// The NewTest function should be called at the beginning of every test function,
	// shouldn't be used unside the subtests.
	//
	// Note, that we want to run each test in a separate transaction,
	// and rollback that transaction at the end of the test.
	// That is done to ensure that the tests are isolated from each other.
	NewTest = func(t *testing.T) (context.Context, *require.Assertions, entity.Entities) {
		// Make sure to run tests in parallel
		t.Parallel()

		tx, err := db.BeginTX(context.Background())
		require.NoError(t, err)

		t.Cleanup(func() {
			tx.Rollback()
		})

		return context.Background(), require.New(t), entity.NewEntities(entity.NewDependency(tx, services))
	}

	// Run the tests
	m.Run()
}
