package entimpl

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/ent"
	"github.com/abc-valera/netsly-api-golang/internal/domain/coderr"

	_ "github.com/lib/pq"
)

func InitEntClient(postgresUrl string) (*ent.Client, error) {
	// Connect to the database
	client, err := ent.Open(
		"postgres",
		postgresUrl,
	)
	if err != nil {
		return nil, coderr.NewInternalErr(err)
	}

	// Run the auto migration tool
	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, coderr.NewInternalErr(err)
	}

	return client, nil
}
