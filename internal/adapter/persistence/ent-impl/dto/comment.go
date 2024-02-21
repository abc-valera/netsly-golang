package dto

import (
	"github.com/abc-valera/netsly-api-golang/gen/ent"
	errhandler "github.com/abc-valera/netsly-api-golang/internal/adapter/persistence/ent-impl/errors"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model/common"
)

func FromEntComment(entComment *ent.Comment) model.Comment {
	if entComment == nil {
		return model.Comment{}
	}
	return model.Comment{
		BaseEntity: common.BaseEntity{
			ID:        entComment.ID,
			CreatedAt: entComment.CreatedAt,
		},
		UserID: entComment.UserID,
		JokeID: entComment.JokeID,
		Text:   entComment.Text,
	}
}

func FromEntCommentToCommentWithErrHandle(entComment *ent.Comment, err error) (model.Comment, error) {
	return FromEntComment(entComment), errhandler.HandleErr(err)
}

func FromEntCommentsToComments(entComments []*ent.Comment) model.Comments {
	comments := make(model.Comments, len(entComments))
	for i, entComment := range entComments {
		comments[i] = FromEntComment(entComment)
	}
	return comments
}

func FromEntCommentsToCommentsWithErrHandle(entComments []*ent.Comment, err error) (model.Comments, error) {
	return FromEntCommentsToComments(entComments), errhandler.HandleErr(err)
}
