package boilerSqliteDto

import (
	"github.com/abc-valera/netsly-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-golang/internal/domain/model"
)

func NewDomainLike(like *sqlboiler.Like) model.Like {
	if like == nil {
		return model.Like{}
	}

	return model.Like{
		CreatedAt: like.CreatedAt,
		DeletedAt: like.DeletedAt,
	}
}

func NewDomainLikes(likes sqlboiler.LikeSlice) model.Likes {
	var domainLikes model.Likes
	for _, like := range likes {
		domainLikes = append(domainLikes, NewDomainLike(like))
	}
	return domainLikes
}
