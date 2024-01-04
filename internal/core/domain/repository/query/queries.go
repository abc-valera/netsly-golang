package query

type Queries struct {
	User        IUserQuery
	Joke        IJokeQuery
	Like        ILikeQuery
	Comment     ICommentQuery
	ChatRoom    IChatRoomQuery
	ChatMember  IChatMemberQuery
	ChatMessage IChatMessageQuery
}
