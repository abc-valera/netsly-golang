package persistence

import (
	"fmt"
	"reflect"

	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
)

type Commands struct {
	User        command.IUser
	Joke        command.IJoke
	Like        command.ILike
	Comment     command.IComment
	Room        command.IRoom
	RoomMember  command.IRoomMember
	RoomMessage command.IRoomMessage
	FileInfo    command.IFileInfo
	FileContent command.IFileContent
}

func ValidateCommands(c Commands) error {
	defer func() {
		if r := recover(); r != nil {
			coderr.NewInternalString(fmt.Sprintf("Recovered: %v", r))
		}
	}()

	reflectVal := reflect.ValueOf(c)

	for i := range reflectVal.NumField() {
		field := reflectVal.Field(i)

		if field.Interface() == nil {
			return coderr.NewInternalString(reflectVal.Type().Field(i).Name + "Command is nil")
		} else {
			continue
		}
	}

	return nil
}
