package dto

import (
	"github.com/abc-valera/netsly-api-golang/gen/ent"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
)

func FromEntLike(entLike *ent.Like) model.Like {
	return model.Like{
		UserID:    entLike.UserID,
		JokeID:    entLike.JokeID,
		CreatedAt: entLike.CreatedAt,
	}
}
