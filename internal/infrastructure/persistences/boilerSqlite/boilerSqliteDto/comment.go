package boilerSqliteDto

import (
	"github.com/abc-valera/netsly-golang/gen/sqlboiler"
	"github.com/abc-valera/netsly-golang/internal/domain/model"
)

func NewDomainComment(comment *sqlboiler.Comment) model.Comment {
	if comment == nil {
		return model.Comment{}
	}

	return model.Comment{
		ID:        comment.ID,
		Text:      comment.Text,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
		DeletedAt: comment.DeletedAt,
	}
}

func NewDomainComments(comments sqlboiler.CommentSlice) model.Comments {
	var domainComments model.Comments
	for _, comment := range comments {
		domainComments = append(domainComments, NewDomainComment(comment))
	}
	return domainComments
}
