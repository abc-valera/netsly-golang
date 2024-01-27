package core

import (
	"errors"
	"log"

	"github.com/abc-valera/flugo-api-golang/internal/core/coderr"
	"github.com/abc-valera/flugo-api-golang/internal/core/persistence/command"
	"github.com/abc-valera/flugo-api-golang/internal/core/persistence/query"
)

type Commands struct {
	User        command.IUser
	Joke        command.IJoke
	Like        command.ILike
	Comment     command.IComment
	ChatRoom    command.IChatRoom
	ChatMember  command.IChatMember
	ChatMessage command.IChatMessage
}

func NewCommands(
	user command.IUser,
	joke command.IJoke,
	like command.ILike,
	comment command.IComment,
	chatRoom command.IChatRoom,
	chatMember command.IChatMember,
	chatMessage command.IChatMessage,
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

type Queries struct {
	User        query.IUser
	Joke        query.IJoke
	Like        query.ILike
	Comment     query.IComment
	ChatRoom    query.IChatRoom
	ChatMember  query.IChatMember
	ChatMessage query.IChatMessage
}

func NewQueries(
	user query.IUser,
	joke query.IJoke,
	like query.ILike,
	comment query.IComment,
	chatRoom query.IChatRoom,
	chatMember query.IChatMember,
	chatMessage query.IChatMessage,
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
