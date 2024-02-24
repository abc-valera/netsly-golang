// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/abc-valera/netsly-api-golang/gen/ent/comment"
	"github.com/abc-valera/netsly-api-golang/gen/ent/joke"
	"github.com/abc-valera/netsly-api-golang/gen/ent/room"
	"github.com/abc-valera/netsly-api-golang/gen/ent/roommessage"
	"github.com/abc-valera/netsly-api-golang/gen/ent/user"
	"github.com/abc-valera/netsly-api-golang/internal/adapter/persistence/ent-impl/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	commentFields := schema.Comment{}.Fields()
	_ = commentFields
	// commentDescText is the schema descriptor for text field.
	commentDescText := commentFields[1].Descriptor()
	// comment.TextValidator is a validator for the "text" field. It is called by the builders before save.
	comment.TextValidator = commentDescText.Validators[0].(func(string) error)
	// commentDescID is the schema descriptor for id field.
	commentDescID := commentFields[0].Descriptor()
	// comment.IDValidator is a validator for the "id" field. It is called by the builders before save.
	comment.IDValidator = commentDescID.Validators[0].(func(string) error)
	jokeFields := schema.Joke{}.Fields()
	_ = jokeFields
	// jokeDescTitle is the schema descriptor for title field.
	jokeDescTitle := jokeFields[1].Descriptor()
	// joke.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	joke.TitleValidator = jokeDescTitle.Validators[0].(func(string) error)
	// jokeDescText is the schema descriptor for text field.
	jokeDescText := jokeFields[2].Descriptor()
	// joke.TextValidator is a validator for the "text" field. It is called by the builders before save.
	joke.TextValidator = jokeDescText.Validators[0].(func(string) error)
	// jokeDescID is the schema descriptor for id field.
	jokeDescID := jokeFields[0].Descriptor()
	// joke.IDValidator is a validator for the "id" field. It is called by the builders before save.
	joke.IDValidator = jokeDescID.Validators[0].(func(string) error)
	roomFields := schema.Room{}.Fields()
	_ = roomFields
	// roomDescName is the schema descriptor for name field.
	roomDescName := roomFields[1].Descriptor()
	// room.NameValidator is a validator for the "name" field. It is called by the builders before save.
	room.NameValidator = roomDescName.Validators[0].(func(string) error)
	// roomDescID is the schema descriptor for id field.
	roomDescID := roomFields[0].Descriptor()
	// room.IDValidator is a validator for the "id" field. It is called by the builders before save.
	room.IDValidator = roomDescID.Validators[0].(func(string) error)
	roommessageFields := schema.RoomMessage{}.Fields()
	_ = roommessageFields
	// roommessageDescText is the schema descriptor for text field.
	roommessageDescText := roommessageFields[1].Descriptor()
	// roommessage.TextValidator is a validator for the "text" field. It is called by the builders before save.
	roommessage.TextValidator = roommessageDescText.Validators[0].(func(string) error)
	// roommessageDescID is the schema descriptor for id field.
	roommessageDescID := roommessageFields[0].Descriptor()
	// roommessage.IDValidator is a validator for the "id" field. It is called by the builders before save.
	roommessage.IDValidator = roommessageDescID.Validators[0].(func(string) error)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[1].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = userDescUsername.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[2].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescHashedPassword is the schema descriptor for hashed_password field.
	userDescHashedPassword := userFields[3].Descriptor()
	// user.HashedPasswordValidator is a validator for the "hashed_password" field. It is called by the builders before save.
	user.HashedPasswordValidator = userDescHashedPassword.Validators[0].(func(string) error)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.IDValidator is a validator for the "id" field. It is called by the builders before save.
	user.IDValidator = userDescID.Validators[0].(func(string) error)
}
