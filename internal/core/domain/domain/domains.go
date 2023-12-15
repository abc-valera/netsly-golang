package domain

import (
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/command"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/query"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/service"
)

type Domains struct {
	User        UserDomain
	Joke        JokeDomain
	Like        LikeDomain
	Comment     CommentDomain
	ChatRoom    ChatRoomDomain
	ChatMember  ChatMemberDomain
	ChatMessage ChatMessageDomain
}

func NewDomains(
	commands command.Commands,
	queries query.Queries,
	services service.Services,
) Domains {
	return Domains{
		User:        NewUserDomain(commands.User, queries.User, services.PasswordMaker),
		Joke:        NewJokeDomain(commands.Joke),
		Like:        NewLikeDomain(commands.Like),
		Comment:     NewCommentDomain(commands.Comment),
		ChatRoom:    NewChatRoomDomain(commands.ChatRoom),
		ChatMember:  NewChatMemberDomain(commands.ChatMember),
		ChatMessage: NewChatMessageDomain(commands.ChatMessage),
	}
}
