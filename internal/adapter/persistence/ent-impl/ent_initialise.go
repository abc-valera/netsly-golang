package entimpl

import (
	"context"
	"os"

	"github.com/abc-valera/netsly-api-golang/gen/ent"
	"github.com/abc-valera/netsly-api-golang/internal/domain/coderr"

	_ "github.com/lib/pq"
)

var (
	postgresUrl = os.Getenv("POSTGRES_URL")
)

func InitEntClient() (*ent.Client, error) {
	// Connect to the database
	client, err := ent.Open(
		"postgres",
		postgresUrl,
	)
	if err != nil {
		return nil, coderr.NewInternal(err)
	}

	// Run the auto migration tool
	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, coderr.NewInternal(err)
	}

	return client, nil
}
