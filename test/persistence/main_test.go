package persistence_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	domainPersistence "github.com/abc-valera/netsly-api-golang/internal/domain/persistence"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/boiler"
	"github.com/abc-valera/netsly-api-golang/test"
	"github.com/docker/docker/pkg/ioutils"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

// initTest initializes environment for tests
// and should be called in every test function from this package.
var initTest func(t *testing.T) (context.Context, *require.Assertions, domainPersistence.Commands, domainPersistence.Queries)

func TestMain(m *testing.M) {
	test.InitTestMain()

	ctx, ctxCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer ctxCancel()

	// Disable testcontainers logs
	testcontainers.Logger = log.New(&ioutils.NopWriter{}, "", 0)
	// Init Postgres testcontainer
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Env: map[string]string{
				"POSTGRES_USER":     "test",
				"POSTGRES_PASSWORD": "test",
				"POSTGRES_DB":       "test",
			},
			ExposedPorts: []string{"5432/tcp"},
			Image:        "postgres:15-alpine",
			WaitingFor: wait.
				ForExec([]string{"pg_isready"}).
				WithStartupTimeout(5 * time.Second).
				WithPollInterval(250 * time.Millisecond).
				WithExitCodeMatcher(func(exitCode int) bool { return exitCode == 0 }),
		},
		Started: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer container.Terminate(ctx)

	endpoint, err := container.Endpoint(ctx, "")
	if err != nil {
		log.Fatal(err)
	}

	// Init DB
	db, err := boiler.New(fmt.Sprintf("postgres://test:test@%s/test?sslmode=disable", endpoint))
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	initTest = func(t *testing.T) (context.Context, *require.Assertions, domainPersistence.Commands, domainPersistence.Queries) {
		// Make sure tests run in parallel
		t.Parallel()

		// Init require assertion
		r := require.New(t)

		// Start transaction
		tx, err := db.BeginTx(ctx, nil)
		r.NoError(err)
		t.Cleanup(func() { r.NoError(tx.Rollback()) })

		// Write fixtures to the transaction
		fixtures, err := os.ReadFile("fixtures/fixtures.sql")
		r.NoError(err)
		res, err := tx.Exec(string(fixtures))
		r.NoError(err)
		_ = res

		return ctx, r,
			implementation.NewCommands(tx, ""),
			implementation.NewQueries(tx, "")
	}

	m.Run()
}
