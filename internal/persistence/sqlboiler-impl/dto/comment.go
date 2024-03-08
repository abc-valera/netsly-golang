package dto

import (
	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model/common"
	"github.com/abc-valera/netsly-api-golang/internal/persistence/sqlboiler-impl/errors"
)

func ToDomainComment(comment *sqlboiler.Comment) model.Comment {
	if comment == nil {
		return model.Comment{}
	}

	return model.Comment{
		BaseEntity: common.BaseEntity{
			ID:        comment.ID,
			CreatedAt: comment.CreatedAt,
		},
		Text:   comment.Text,
		UserID: comment.UserID,
	}
}

func ToDomainCommentWithErrHandle(comment *sqlboiler.Comment, err error) (model.Comment, error) {
	return ToDomainComment(comment), errors.HandleErr(err)
}

func ToDomainComments(comments sqlboiler.CommentSlice) model.Comments {
	var domainComments model.Comments
	for _, comment := range comments {
		domainComments = append(domainComments, ToDomainComment(comment))
	}
	return domainComments
}

func ToDomainCommentsWithErrHandle(comments sqlboiler.CommentSlice, err error) (model.Comments, error) {
	return ToDomainComments(comments), errors.HandleErr(err)
}
