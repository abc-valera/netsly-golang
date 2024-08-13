package restDto

import (
	"github.com/abc-valera/netsly-golang/gen/ogen"
	"github.com/abc-valera/netsly-golang/internal/domain/model"
)

func NewComment(comment model.Comment) *ogen.Comment {
	return &ogen.Comment{
		ID:        comment.ID,
		Text:      comment.Text,
		CreatedAt: comment.CreatedAt,
	}
}

func NewComments(comments []model.Comment) *ogen.Comments {
	res := make([]ogen.Comment, 0, len(comments))
	for _, comment := range comments {
		res = append(res, *NewComment(comment))
	}
	return &ogen.Comments{Comments: res}
}
