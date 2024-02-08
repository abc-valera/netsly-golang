package domain

import (
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query"
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
