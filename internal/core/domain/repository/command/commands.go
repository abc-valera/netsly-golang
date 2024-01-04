package command

type Commands struct {
	User        IUserCommand
	Joke        IJokeCommand
	Like        ILikeCommand
	Comment     ICommentCommand
	ChatRoom    IChatRoomCommand
	ChatMember  IChatMemberCommand
	ChatMessage IChatMessageCommand
}
