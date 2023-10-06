package impl

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/gen/ent"
)

type baseRepository struct {
	client *ent.Client
}

func NewBaseRepository(client *ent.Client) baseRepository {
	return baseRepository{
		client: client,
	}
}

func (r *baseRepository) PerformTX(ctx context.Context, txFunc func(ctx context.Context) error) error {
	tx, err := r.client.Tx(ctx)
	if err != nil {
		return handleErr(err)
	}

	oldClient := r.client
	defer func() {
		r.client = oldClient
	}()

	r.client = tx.Client()
	if err := txFunc(ctx); err != nil {
		if err := tx.Rollback(); err != nil {
			return handleErr(err)
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return handleErr(err)
	}
	return nil
}
