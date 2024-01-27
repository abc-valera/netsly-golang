package domain

import (
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/domainval"
)

type Domains struct {
	User        domainval.User
	Joke        domainval.Joke
	Like        domainval.Like
	Comment     domainval.Comment
	ChatRoom    domainval.ChatRoom
	ChatMember  domainval.ChatMember
	ChatMessage domainval.ChatMessage
}

func NewDomains(
	commands Commands,
	queries Queries,
	services Services,
) Domains {
	return Domains{
		User:        domainval.NewUser(commands.User, queries.User, services.PasswordMaker),
		Joke:        domainval.NewJoke(commands.Joke),
		Like:        domainval.NewLike(commands.Like),
		Comment:     domainval.NewComment(commands.Comment),
		ChatRoom:    domainval.NewChatRoom(commands.ChatRoom),
		ChatMember:  domainval.NewChatMember(commands.ChatMember),
		ChatMessage: domainval.NewChatMessage(commands.ChatMessage),
	}
}
