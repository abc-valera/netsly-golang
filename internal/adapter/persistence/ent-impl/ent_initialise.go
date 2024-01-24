package entimpl

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/gen/ent"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/coderr"

	_ "github.com/lib/pq"
)

func InitEntClient(databaseURL string) (*ent.Client, error) {
	// Connect to the database
	client, err := ent.Open(
		"postgres",
		databaseURL,
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