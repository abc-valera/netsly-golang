package command

import (
	"github.com/abc-valera/netsly-golang/internal/domain/model"
)

type Commands struct {
	User         ICreateUpdateDelete[model.User]
	Joke         ICreateUpdateDelete[model.Joke]
	Like         ICreateDelete[model.Like]
	Comment      ICreateUpdateDelete[model.Comment]
	Room         ICreateUpdateDelete[model.Room]
	RoomMember   ICreateDelete[model.RoomMember]
	RoomMessage  ICreateUpdateDelete[model.RoomMessage]
	FileInfo     ICreateUpdateDelete[model.FileInfo]
	FileInfoJoke ICreate[model.FileInfoJoke]
	FileInfoRoom ICreate[model.FileInfoRoom]
	FileContent  ICreateUpdateDelete[model.FileContent]
}

// TODO: maybe define errors that can be returned by the commands here
