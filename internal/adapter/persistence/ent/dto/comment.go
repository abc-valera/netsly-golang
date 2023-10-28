package dto

import (
	"github.com/abc-valera/flugo-api-golang/gen/ent"
	"github.com/abc-valera/flugo-api-golang/internal/adapter/persistence/ent/dto/common"
	"github.com/abc-valera/flugo-api-golang/internal/domain/entity"
)

func FromEntCommentToComment(entComment *ent.Comment) *entity.Comment {
	if entComment == nil {
		return nil
	}
	return &entity.Comment{
		BaseEntity: common.FromEntToBaseEntity(entComment.ID, entComment.CreatedAt),
		UserID:     entComment.UserID,
		JokeID:     entComment.JokeID,
		Text:       entComment.Text,
	}
}

func FromEntCommentsToComments(entComments []*ent.Comment) entity.Comments {
	comments := make(entity.Comments, len(entComments))
	for i, entComment := range entComments {
		comments[i] = FromEntCommentToComment(entComment)
	}
	return comments
}
