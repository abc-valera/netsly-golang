package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/abc-valera/netsly-api-golang/internal/core/mode"
	"github.com/abc-valera/netsly-api-golang/internal/domain"
	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/commandTransactor"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/gormSqlite"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/localFileSaver"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/service/emailSender/dummyEmailSender"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/service/logger/slogLogger"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/service/passworder"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/service/taskQueuer/dummyTaskQueuer"
	"go.opentelemetry.io/otel/trace/noop"
)

// TestMain is the entry point for the test suite.
// It is responsible for setting up the infrastructure for the test environment.
// You can change the infrastructure you want to test.
//
// Note, that you shouldn't use os.Exit() in your tests, because it will terminate the cleanup process.
// Also, make sure the test won't panic, because it will also terminate the cleanup process.
// If the panic is expected, then recover from it inside the test.
func TestMain(m *testing.M) {
	if err := Initiliaze(); err != nil {
		fmt.Println("Initialization failed: ", err)
		CleanupExit(1)
	}

	CleanupExit(m.Run())
}

var entities domain.Entities

const tmpDir = "./tmp"

// Initiliaze initializes the test environment.
func Initiliaze() error {
	// Create a directory for temporary test files.
	// If the directory already exists, delete it and create a new one.
	if err := os.Mkdir(tmpDir, 0o755); err != nil {
		if os.IsExist(err) {
			if err := os.RemoveAll(tmpDir); err != nil {
				return err
			}
			if err := os.Mkdir(tmpDir, 0o755); err != nil {
				return err
			}
		} else {
			return err
		}
	}

	// Init services
	var services domain.Services

	// services.Logger = nopLogger.New()
	services.Logger = slogLogger.New(tmpDir)

	services.EmailSender = dummyEmailSender.New()

	services.Passworder = passworder.New()

	services.TaskQueuer = dummyTaskQueuer.New(services.EmailSender)

	// Init global
	global.Init(
		mode.Mode(mode.Production),
		noop.NewTracerProvider().Tracer("noop"),
		services.Logger,
	)

	// Init persistence
	gormSqliteDependency, err := gormSqlite.New(tmpDir + "/gorm_sqlite.db")
	if err != nil {
		return err
	}
	// boilerSqliteDependency, err := boilerSqlite.New(tmpDir + "/boiler_sqlite.db")
	// if err != nil {
	// 	return err
	// }
	localFileSaver, err := localFileSaver.New(tmpDir)
	if err != nil {
		return err
	}

	commands, queries, err := implementation.NewCommandsAndQueries(implementation.CommandsAndQueriesDependencies{
		GormSqlite: gormSqliteDependency,
		// BoilerSqlite: boilerSqliteDependency,
		LocalFileSaver: localFileSaver,
	})
	if err != nil {
		return err
	}

	// Init command transactor
	commandTransactor := commandTransactor.New(commandTransactor.Dependencies{
		GormSqlite: gormSqliteDependency,
		// BoilerSqlite: boilerSqliteDependency,
		LocalFileSaver: localFileSaver,
	})

	// Init entities
	entities = domain.NewEntities(commands, commandTransactor, queries, services)

	// Seed the persistence layer
	if err := Seed(entities); err != nil {
		return err
	}

	return nil
}

// CleanupExit cleans up the test environment and exits with the given code.
func CleanupExit(code int) {
	if err := os.RemoveAll(tmpDir); err != nil {
		fmt.Println("Cleanup failed: ", err)
	}

	os.Exit(code)
}
