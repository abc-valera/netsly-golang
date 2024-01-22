package query

import (
	"errors"
	"log"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/coderr"
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
) Queries {
	if user == nil {
		log.Fatal(coderr.NewInternal(errors.New("user query is nil")))
	}
	if joke == nil {
		log.Fatal(coderr.NewInternal(errors.New("joke query is nil")))
	}
	if like == nil {
		log.Fatal(coderr.NewInternal(errors.New("like query is nil")))
	}
	if comment == nil {
		log.Fatal(coderr.NewInternal(errors.New("comment query is nil")))
	}
	if chatRoom == nil {
		log.Fatal(coderr.NewInternal(errors.New("chatRoom query is nil")))
	}
	if chatMember == nil {
		log.Fatal(coderr.NewInternal(errors.New("chatMember query is nil")))
	}
	if chatMessage == nil {
		log.Fatal(coderr.NewInternal(errors.New("chatMessage query is nil")))
	}

	return Queries{
		User:        user,
		Joke:        joke,
		Like:        like,
		Comment:     comment,
		ChatRoom:    chatRoom,
		ChatMember:  chatMember,
		ChatMessage: chatMessage,
	}
}
