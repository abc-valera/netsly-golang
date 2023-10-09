package common

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/gen/ent"
)

type BaseRepository struct {
	Client *ent.Client
}

func NewBaseRepository(client *ent.Client) BaseRepository {
	return BaseRepository{
		Client: client,
	}
}

func (r *BaseRepository) PerformTX(ctx context.Context, txFunc func(ctx context.Context) error) error {
	tx, err := r.Client.Tx(ctx)
	if err != nil {
		return HandleErr(err)
	}

	oldClient := r.Client
	defer func() {
		r.Client = oldClient
	}()

	r.Client = tx.Client()
	if err := txFunc(ctx); err != nil {
		if err := tx.Rollback(); err != nil {
			return HandleErr(err)
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return HandleErr(err)
	}
	return nil
}
