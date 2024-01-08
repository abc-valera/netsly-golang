package query

import (
	"errors"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/codeerr"
)

type Queries struct {
	User        IUser
	Joke        IJoke
	Like        ILike
	Comment     IComment
	ChatRoom    IChatRoom
	ChatMember  IChatMember
	ChatMessage IChatMessage
}

func NewQueries(
	user IUser,
	joke IJoke,
	like ILike,
	comment IComment,
	chatRoom IChatRoom,
	chatMember IChatMember,
	chatMessage IChatMessage,
) (Queries, error) {
	if user == nil {
		return Queries{}, codeerr.NewInternal(errors.New("user query is nil"))
	}
	if joke == nil {
		return Queries{}, codeerr.NewInternal(errors.New("joke query is nil"))
	}
	if like == nil {
		return Queries{}, codeerr.NewInternal(errors.New("like query is nil"))
	}
	if comment == nil {
		return Queries{}, codeerr.NewInternal(errors.New("comment query is nil"))
	}
	if chatRoom == nil {
		return Queries{}, codeerr.NewInternal(errors.New("chatRoom query is nil"))
	}
	if chatMember == nil {
		return Queries{}, codeerr.NewInternal(errors.New("chatMember query is nil"))
	}
	if chatMessage == nil {
		return Queries{}, codeerr.NewInternal(errors.New("chatMessage query is nil"))
	}
	return Queries{
		User:        user,
		Joke:        joke,
		Like:        like,
		Comment:     comment,
		ChatRoom:    chatRoom,
		ChatMember:  chatMember,
		ChatMessage: chatMessage,
	}, nil
}
