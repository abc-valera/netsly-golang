package query

type Queries struct {
	User        IUser
	Joke        IJoke
	Like        ILike
	Comment     IComment
	Room        IRoom
	RoomMember  IRoomMember
	RoomMessage IRoomMessage
	FileInfo    IFileInfo
	FileContent IFileContent
}
