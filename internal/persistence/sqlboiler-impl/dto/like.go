package dto

import (
	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
	"github.com/abc-valera/netsly-api-golang/internal/persistence/sqlboiler-impl/errors"
)

func ToDomainLike(like *sqlboiler.Like) model.Like {
	if like == nil {
		return model.Like{}
	}

	return model.Like{
		CreatedAt: like.CreatedAt,
		UserID:    like.UserID,
		JokeID:    like.JokeID,
	}
}

func ToDomainLikeWithErrHandle(like *sqlboiler.Like, err error) (model.Like, error) {
	return ToDomainLike(like), errors.HandleErr(err)
}

func ToDomainLikes(likes sqlboiler.LikeSlice) model.Likes {
	var domainLikes model.Likes
	for _, like := range likes {
		domainLikes = append(domainLikes, ToDomainLike(like))
	}
	return domainLikes
}

func ToDomainLikesWithErrHandle(likes sqlboiler.LikeSlice, err error) (model.Likes, error) {
	return ToDomainLikes(likes), errors.HandleErr(err)
}
