package dto

import (
	"github.com/abc-valera/flugo-api-golang/gen/ent"
	errhandler "github.com/abc-valera/flugo-api-golang/internal/adapter/persistence/ent/err-handler"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/model"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/model/common"
)

func FromEntCommentToComment(entComment *ent.Comment) model.Comment {
	if entComment == nil {
		return model.Comment{}
	}
	return model.Comment{
		BaseModel: common.BaseModel{
			ID:        entComment.ID,
			CreatedAt: entComment.CreatedAt,
		},
		UserID: entComment.UserID,
		JokeID: entComment.JokeID,
		Text:   entComment.Text,
	}
}

func FromEntCommentToCommentWithErrHandle(entComment *ent.Comment, err error) (model.Comment, error) {
	return FromEntCommentToComment(entComment), errhandler.HandleErr(err)
}

func FromEntCommentsToComments(entComments []*ent.Comment) model.Comments {
	comments := make(model.Comments, len(entComments))
	for i, entComment := range entComments {
		comments[i] = FromEntCommentToComment(entComment)
	}
	return comments
}

func FromEntCommentsToCommentsWithErrHandle(entComments []*ent.Comment, err error) (model.Comments, error) {
	return FromEntCommentsToComments(entComments), errhandler.HandleErr(err)
}
