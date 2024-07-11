package boilerDto

import (
	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/boiler/errutil"
)

func NewDomainLike(like *sqlboiler.Like) model.Like {
	if like == nil {
		return model.Like{}
	}

	return model.Like{
		CreatedAt: like.CreatedAt,
		UserID:    like.UserID,
		JokeID:    like.JokeID,
	}
}

func NewDomainLikeWithErrHandle(like *sqlboiler.Like, err error) (model.Like, error) {
	return NewDomainLike(like), errutil.HandleErr(err)
}

func NewDomainLikes(likes sqlboiler.LikeSlice) model.Likes {
	var domainLikes model.Likes
	for _, like := range likes {
		domainLikes = append(domainLikes, NewDomainLike(like))
	}
	return domainLikes
}

func NewDomainLikesWithErrHandle(likes sqlboiler.LikeSlice, err error) (model.Likes, error) {
	return NewDomainLikes(likes), errutil.HandleErr(err)
}
