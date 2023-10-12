package dto

import (
	"github.com/abc-valera/flugo-api-golang/gen/ogen"
	"github.com/abc-valera/flugo-api-golang/internal/domain/entity"
)

func NewCommentResponse(comment *entity.Comment) *ogen.Comment {
	return &ogen.Comment{
		ID:        comment.ID,
		JokeID:    comment.JokeID,
		UserID:    comment.UserID,
		Text:      comment.Text,
		CreatedAt: comment.CreatedAt,
	}
}

func NewCommentsResponse(comments []*entity.Comment) *ogen.Comments {
	res := make([]ogen.Comment, 0, len(comments))
	for _, comment := range comments {
		res = append(res, *NewCommentResponse(comment))
	}
	return &ogen.Comments{Comments: res}
}
