package command

import (
	"errors"
	"log"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/coderr"
)

type Commands struct {
	User        IUser
	Joke        IJoke
	Like        ILike
	Comment     IComment
	ChatRoom    IChatRoom
	ChatMember  IChatMember
	ChatMessage IChatMessage
}

func NewCommands(
	user IUser,
	joke IJoke,
	like ILike,
	comment IComment,
	chatRoom IChatRoom,
	chatMember IChatMember,
	chatMessage IChatMessage,
) Commands {
	if user == nil {
		log.Fatal(coderr.NewInternal(errors.New("user command is nil")))
	}
	if joke == nil {
		log.Fatal(coderr.NewInternal(errors.New("joke command is nil")))
	}
	if like == nil {
		log.Fatal(coderr.NewInternal(errors.New("like command is nil")))
	}
	if comment == nil {
		log.Fatal(coderr.NewInternal(errors.New("comment command is nil")))
	}
	if chatRoom == nil {
		log.Fatal(coderr.NewInternal(errors.New("chatRoom ommand is nil")))
	}
	if chatMember == nil {
		log.Fatal(coderr.NewInternal(errors.New("chatMember command is nil")))
	}
	if chatMessage == nil {
		log.Fatal(coderr.NewInternal(errors.New("chatMessage command is nil")))
	}

	return Commands{
		User:        user,
		Joke:        joke,
		Like:        like,
		Comment:     comment,
		ChatRoom:    chatRoom,
		ChatMember:  chatMember,
		ChatMessage: chatMessage,
	}
}
