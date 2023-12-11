package dto

import (
	"github.com/abc-valera/flugo-api-golang/gen/ent"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/model"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/model/common"
)

func FromEntCommentToComment(entComment *ent.Comment) *model.Comment {
	if entComment == nil {
		return nil
	}
	return &model.Comment{
		BaseModel: common.BaseModel{
			ID:        entComment.ID,
			CreatedAt: entComment.CreatedAt,
		},
		UserID: entComment.UserID,
		JokeID: entComment.JokeID,
		Text:   entComment.Text,
	}
}

func FromEntCommentsToComments(entComments []*ent.Comment) model.Comments {
	comments := make(model.Comments, len(entComments))
	for i, entComment := range entComments {
		comments[i] = FromEntCommentToComment(entComment)
	}
	return comments
}
