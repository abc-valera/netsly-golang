package boilerDto

import (
	"github.com/abc-valera/netsly-api-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/boiler/errutil"
)

func NewDomainComment(comment *sqlboiler.Comment) model.Comment {
	if comment == nil {
		return model.Comment{}
	}

	return model.Comment{
		ID:        comment.ID,
		Text:      comment.Text,
		CreatedAt: comment.CreatedAt,
	}
}

func NewDomainCommentWithErrHandle(comment *sqlboiler.Comment, err error) (model.Comment, error) {
	return NewDomainComment(comment), errutil.HandleErr(err)
}

func NewDomainComments(comments sqlboiler.CommentSlice) model.Comments {
	var domainComments model.Comments
	for _, comment := range comments {
		domainComments = append(domainComments, NewDomainComment(comment))
	}
	return domainComments
}

func NewDomainCommentsWithErrHandle(comments sqlboiler.CommentSlice, err error) (model.Comments, error) {
	return NewDomainComments(comments), errutil.HandleErr(err)
}
