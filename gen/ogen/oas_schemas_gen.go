// Code generated by ogen, DO NOT EDIT.

package ogen

import (
	"fmt"
	"time"

	"github.com/go-faster/errors"
)

func (s *CodeErrorStatusCode) Error() string {
	return fmt.Sprintf("code %d: %+v", s.StatusCode, s.Response)
}

type BearerAuth struct {
	Token string
}

// GetToken returns the value of Token.
func (s *BearerAuth) GetToken() string {
	return s.Token
}

// SetToken sets the value of Token.
func (s *BearerAuth) SetToken(val string) {
	s.Token = val
}

// Ref: #/components/schemas/code_error
type CodeError struct {
	Code         CodeErrorCode `json:"code"`
	ErrorMessage string        `json:"error_message"`
}

// GetCode returns the value of Code.
func (s *CodeError) GetCode() CodeErrorCode {
	return s.Code
}

// GetErrorMessage returns the value of ErrorMessage.
func (s *CodeError) GetErrorMessage() string {
	return s.ErrorMessage
}

// SetCode sets the value of Code.
func (s *CodeError) SetCode(val CodeErrorCode) {
	s.Code = val
}

// SetErrorMessage sets the value of ErrorMessage.
func (s *CodeError) SetErrorMessage(val string) {
	s.ErrorMessage = val
}

type CodeErrorCode string

const (
	CodeErrorCodeUnauthenticated  CodeErrorCode = "unauthenticated"
	CodeErrorCodeInvalidArgument  CodeErrorCode = "invalid_argument"
	CodeErrorCodeNotFound         CodeErrorCode = "not_found"
	CodeErrorCodeAlreadyExists    CodeErrorCode = "already_exists"
	CodeErrorCodePermissionDenied CodeErrorCode = "permission_denied"
	CodeErrorCodeInternal         CodeErrorCode = "internal"
)

// AllValues returns all CodeErrorCode values.
func (CodeErrorCode) AllValues() []CodeErrorCode {
	return []CodeErrorCode{
		CodeErrorCodeUnauthenticated,
		CodeErrorCodeInvalidArgument,
		CodeErrorCodeNotFound,
		CodeErrorCodeAlreadyExists,
		CodeErrorCodePermissionDenied,
		CodeErrorCodeInternal,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s CodeErrorCode) MarshalText() ([]byte, error) {
	switch s {
	case CodeErrorCodeUnauthenticated:
		return []byte(s), nil
	case CodeErrorCodeInvalidArgument:
		return []byte(s), nil
	case CodeErrorCodeNotFound:
		return []byte(s), nil
	case CodeErrorCodeAlreadyExists:
		return []byte(s), nil
	case CodeErrorCodePermissionDenied:
		return []byte(s), nil
	case CodeErrorCodeInternal:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *CodeErrorCode) UnmarshalText(data []byte) error {
	switch CodeErrorCode(data) {
	case CodeErrorCodeUnauthenticated:
		*s = CodeErrorCodeUnauthenticated
		return nil
	case CodeErrorCodeInvalidArgument:
		*s = CodeErrorCodeInvalidArgument
		return nil
	case CodeErrorCodeNotFound:
		*s = CodeErrorCodeNotFound
		return nil
	case CodeErrorCodeAlreadyExists:
		*s = CodeErrorCodeAlreadyExists
		return nil
	case CodeErrorCodePermissionDenied:
		*s = CodeErrorCodePermissionDenied
		return nil
	case CodeErrorCodeInternal:
		*s = CodeErrorCodeInternal
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// CodeErrorStatusCode wraps CodeError with StatusCode.
type CodeErrorStatusCode struct {
	StatusCode int
	Response   CodeError
}

// GetStatusCode returns the value of StatusCode.
func (s *CodeErrorStatusCode) GetStatusCode() int {
	return s.StatusCode
}

// GetResponse returns the value of Response.
func (s *CodeErrorStatusCode) GetResponse() CodeError {
	return s.Response
}

// SetStatusCode sets the value of StatusCode.
func (s *CodeErrorStatusCode) SetStatusCode(val int) {
	s.StatusCode = val
}

// SetResponse sets the value of Response.
func (s *CodeErrorStatusCode) SetResponse(val CodeError) {
	s.Response = val
}

// Ref: #/components/schemas/comment
type Comment struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	JokeID    string    `json:"joke_id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
}

// GetID returns the value of ID.
func (s *Comment) GetID() string {
	return s.ID
}

// GetUserID returns the value of UserID.
func (s *Comment) GetUserID() string {
	return s.UserID
}

// GetJokeID returns the value of JokeID.
func (s *Comment) GetJokeID() string {
	return s.JokeID
}

// GetText returns the value of Text.
func (s *Comment) GetText() string {
	return s.Text
}

// GetCreatedAt returns the value of CreatedAt.
func (s *Comment) GetCreatedAt() time.Time {
	return s.CreatedAt
}

// SetID sets the value of ID.
func (s *Comment) SetID(val string) {
	s.ID = val
}

// SetUserID sets the value of UserID.
func (s *Comment) SetUserID(val string) {
	s.UserID = val
}

// SetJokeID sets the value of JokeID.
func (s *Comment) SetJokeID(val string) {
	s.JokeID = val
}

// SetText sets the value of Text.
func (s *Comment) SetText(val string) {
	s.Text = val
}

// SetCreatedAt sets the value of CreatedAt.
func (s *Comment) SetCreatedAt(val time.Time) {
	s.CreatedAt = val
}

type Comments struct {
	Comments []Comment `json:"comments"`
}

// GetComments returns the value of Comments.
func (s *Comments) GetComments() []Comment {
	return s.Comments
}

// SetComments sets the value of Comments.
func (s *Comments) SetComments(val []Comment) {
	s.Comments = val
}

type CommentsByJokeIDGetSelector struct {
	Order  OptOrder `json:"order"`
	Limit  OptInt   `json:"limit"`
	Offset OptInt   `json:"offset"`
}

// GetOrder returns the value of Order.
func (s *CommentsByJokeIDGetSelector) GetOrder() OptOrder {
	return s.Order
}

// GetLimit returns the value of Limit.
func (s *CommentsByJokeIDGetSelector) GetLimit() OptInt {
	return s.Limit
}

// GetOffset returns the value of Offset.
func (s *CommentsByJokeIDGetSelector) GetOffset() OptInt {
	return s.Offset
}

// SetOrder sets the value of Order.
func (s *CommentsByJokeIDGetSelector) SetOrder(val OptOrder) {
	s.Order = val
}

// SetLimit sets the value of Limit.
func (s *CommentsByJokeIDGetSelector) SetLimit(val OptInt) {
	s.Limit = val
}

// SetOffset sets the value of Offset.
func (s *CommentsByJokeIDGetSelector) SetOffset(val OptInt) {
	s.Offset = val
}

// Ref: #/components/schemas/joke
type Joke struct {
	ID          string    `json:"id"`
	UserID      string    `json:"userID"`
	Title       string    `json:"title"`
	Text        string    `json:"text"`
	Explanation OptString `json:"explanation"`
	CreatedAt   time.Time `json:"createdAt"`
}

// GetID returns the value of ID.
func (s *Joke) GetID() string {
	return s.ID
}

// GetUserID returns the value of UserID.
func (s *Joke) GetUserID() string {
	return s.UserID
}

// GetTitle returns the value of Title.
func (s *Joke) GetTitle() string {
	return s.Title
}

// GetText returns the value of Text.
func (s *Joke) GetText() string {
	return s.Text
}

// GetExplanation returns the value of Explanation.
func (s *Joke) GetExplanation() OptString {
	return s.Explanation
}

// GetCreatedAt returns the value of CreatedAt.
func (s *Joke) GetCreatedAt() time.Time {
	return s.CreatedAt
}

// SetID sets the value of ID.
func (s *Joke) SetID(val string) {
	s.ID = val
}

// SetUserID sets the value of UserID.
func (s *Joke) SetUserID(val string) {
	s.UserID = val
}

// SetTitle sets the value of Title.
func (s *Joke) SetTitle(val string) {
	s.Title = val
}

// SetText sets the value of Text.
func (s *Joke) SetText(val string) {
	s.Text = val
}

// SetExplanation sets the value of Explanation.
func (s *Joke) SetExplanation(val OptString) {
	s.Explanation = val
}

// SetCreatedAt sets the value of CreatedAt.
func (s *Joke) SetCreatedAt(val time.Time) {
	s.CreatedAt = val
}

type Jokes struct {
	Jokes []Joke `json:"jokes"`
}

// GetJokes returns the value of Jokes.
func (s *Jokes) GetJokes() []Joke {
	return s.Jokes
}

// SetJokes sets the value of Jokes.
func (s *Jokes) SetJokes(val []Joke) {
	s.Jokes = val
}

// MeChatRoomsJoinPostCreated is response for MeChatRoomsJoinPost operation.
type MeChatRoomsJoinPostCreated struct{}

type MeChatRoomsJoinPostReq struct {
	ID string `json:"id"`
}

// GetID returns the value of ID.
func (s *MeChatRoomsJoinPostReq) GetID() string {
	return s.ID
}

// SetID sets the value of ID.
func (s *MeChatRoomsJoinPostReq) SetID(val string) {
	s.ID = val
}

// MeCommentsDelNoContent is response for MeCommentsDel operation.
type MeCommentsDelNoContent struct{}

type MeCommentsDelReq struct {
	CommentID string `json:"comment_id"`
}

// GetCommentID returns the value of CommentID.
func (s *MeCommentsDelReq) GetCommentID() string {
	return s.CommentID
}

// SetCommentID sets the value of CommentID.
func (s *MeCommentsDelReq) SetCommentID(val string) {
	s.CommentID = val
}

type MeCommentsPostReq struct {
	JokeID string `json:"joke_id"`
	Text   string `json:"text"`
}

// GetJokeID returns the value of JokeID.
func (s *MeCommentsPostReq) GetJokeID() string {
	return s.JokeID
}

// GetText returns the value of Text.
func (s *MeCommentsPostReq) GetText() string {
	return s.Text
}

// SetJokeID sets the value of JokeID.
func (s *MeCommentsPostReq) SetJokeID(val string) {
	s.JokeID = val
}

// SetText sets the value of Text.
func (s *MeCommentsPostReq) SetText(val string) {
	s.Text = val
}

type MeCommentsPutReq struct {
	CommentID string    `json:"comment_id"`
	Text      OptString `json:"text"`
}

// GetCommentID returns the value of CommentID.
func (s *MeCommentsPutReq) GetCommentID() string {
	return s.CommentID
}

// GetText returns the value of Text.
func (s *MeCommentsPutReq) GetText() OptString {
	return s.Text
}

// SetCommentID sets the value of CommentID.
func (s *MeCommentsPutReq) SetCommentID(val string) {
	s.CommentID = val
}

// SetText sets the value of Text.
func (s *MeCommentsPutReq) SetText(val OptString) {
	s.Text = val
}

// MeDelNoContent is response for MeDel operation.
type MeDelNoContent struct{}

type MeDelReq struct {
	Password string `json:"password"`
}

// GetPassword returns the value of Password.
func (s *MeDelReq) GetPassword() string {
	return s.Password
}

// SetPassword sets the value of Password.
func (s *MeDelReq) SetPassword(val string) {
	s.Password = val
}

// MeJokesDelNoContent is response for MeJokesDel operation.
type MeJokesDelNoContent struct{}

type MeJokesDelReq struct {
	JokeID string `json:"joke_id"`
}

// GetJokeID returns the value of JokeID.
func (s *MeJokesDelReq) GetJokeID() string {
	return s.JokeID
}

// SetJokeID sets the value of JokeID.
func (s *MeJokesDelReq) SetJokeID(val string) {
	s.JokeID = val
}

type MeJokesGetSelector struct {
	Order  OptOrder `json:"order"`
	Limit  OptInt   `json:"limit"`
	Offset OptInt   `json:"offset"`
}

// GetOrder returns the value of Order.
func (s *MeJokesGetSelector) GetOrder() OptOrder {
	return s.Order
}

// GetLimit returns the value of Limit.
func (s *MeJokesGetSelector) GetLimit() OptInt {
	return s.Limit
}

// GetOffset returns the value of Offset.
func (s *MeJokesGetSelector) GetOffset() OptInt {
	return s.Offset
}

// SetOrder sets the value of Order.
func (s *MeJokesGetSelector) SetOrder(val OptOrder) {
	s.Order = val
}

// SetLimit sets the value of Limit.
func (s *MeJokesGetSelector) SetLimit(val OptInt) {
	s.Limit = val
}

// SetOffset sets the value of Offset.
func (s *MeJokesGetSelector) SetOffset(val OptInt) {
	s.Offset = val
}

type MeJokesPostReq struct {
	Title       string    `json:"title"`
	Text        string    `json:"text"`
	Explanation OptString `json:"explanation"`
}

// GetTitle returns the value of Title.
func (s *MeJokesPostReq) GetTitle() string {
	return s.Title
}

// GetText returns the value of Text.
func (s *MeJokesPostReq) GetText() string {
	return s.Text
}

// GetExplanation returns the value of Explanation.
func (s *MeJokesPostReq) GetExplanation() OptString {
	return s.Explanation
}

// SetTitle sets the value of Title.
func (s *MeJokesPostReq) SetTitle(val string) {
	s.Title = val
}

// SetText sets the value of Text.
func (s *MeJokesPostReq) SetText(val string) {
	s.Text = val
}

// SetExplanation sets the value of Explanation.
func (s *MeJokesPostReq) SetExplanation(val OptString) {
	s.Explanation = val
}

type MeJokesPutReq struct {
	JokeID      string    `json:"joke_id"`
	Title       OptString `json:"title"`
	Text        OptString `json:"text"`
	Explanation OptString `json:"explanation"`
}

// GetJokeID returns the value of JokeID.
func (s *MeJokesPutReq) GetJokeID() string {
	return s.JokeID
}

// GetTitle returns the value of Title.
func (s *MeJokesPutReq) GetTitle() OptString {
	return s.Title
}

// GetText returns the value of Text.
func (s *MeJokesPutReq) GetText() OptString {
	return s.Text
}

// GetExplanation returns the value of Explanation.
func (s *MeJokesPutReq) GetExplanation() OptString {
	return s.Explanation
}

// SetJokeID sets the value of JokeID.
func (s *MeJokesPutReq) SetJokeID(val string) {
	s.JokeID = val
}

// SetTitle sets the value of Title.
func (s *MeJokesPutReq) SetTitle(val OptString) {
	s.Title = val
}

// SetText sets the value of Text.
func (s *MeJokesPutReq) SetText(val OptString) {
	s.Text = val
}

// SetExplanation sets the value of Explanation.
func (s *MeJokesPutReq) SetExplanation(val OptString) {
	s.Explanation = val
}

// MeLikesDelNoContent is response for MeLikesDel operation.
type MeLikesDelNoContent struct{}

type MeLikesDelReq struct {
	JokeID string `json:"joke_id"`
}

// GetJokeID returns the value of JokeID.
func (s *MeLikesDelReq) GetJokeID() string {
	return s.JokeID
}

// SetJokeID sets the value of JokeID.
func (s *MeLikesDelReq) SetJokeID(val string) {
	s.JokeID = val
}

// MeLikesPostCreated is response for MeLikesPost operation.
type MeLikesPostCreated struct{}

type MeLikesPostReq struct {
	JokeID string `json:"joke_id"`
}

// GetJokeID returns the value of JokeID.
func (s *MeLikesPostReq) GetJokeID() string {
	return s.JokeID
}

// SetJokeID sets the value of JokeID.
func (s *MeLikesPostReq) SetJokeID(val string) {
	s.JokeID = val
}

type MePutReq struct {
	Password OptString `json:"password"`
	Fullname OptString `json:"fullname"`
	Status   OptString `json:"status"`
}

// GetPassword returns the value of Password.
func (s *MePutReq) GetPassword() OptString {
	return s.Password
}

// GetFullname returns the value of Fullname.
func (s *MePutReq) GetFullname() OptString {
	return s.Fullname
}

// GetStatus returns the value of Status.
func (s *MePutReq) GetStatus() OptString {
	return s.Status
}

// SetPassword sets the value of Password.
func (s *MePutReq) SetPassword(val OptString) {
	s.Password = val
}

// SetFullname sets the value of Fullname.
func (s *MePutReq) SetFullname(val OptString) {
	s.Fullname = val
}

// SetStatus sets the value of Status.
func (s *MePutReq) SetStatus(val OptString) {
	s.Status = val
}

// MeRoomsDeleteNoContent is response for MeRoomsDelete operation.
type MeRoomsDeleteNoContent struct{}

type MeRoomsDeleteReq struct {
	ID string `json:"id"`
}

// GetID returns the value of ID.
func (s *MeRoomsDeleteReq) GetID() string {
	return s.ID
}

// SetID sets the value of ID.
func (s *MeRoomsDeleteReq) SetID(val string) {
	s.ID = val
}

type MeRoomsGetSelector struct {
	Order  OptOrder `json:"order"`
	Limit  OptInt   `json:"limit"`
	Offset OptInt   `json:"offset"`
}

// GetOrder returns the value of Order.
func (s *MeRoomsGetSelector) GetOrder() OptOrder {
	return s.Order
}

// GetLimit returns the value of Limit.
func (s *MeRoomsGetSelector) GetLimit() OptInt {
	return s.Limit
}

// GetOffset returns the value of Offset.
func (s *MeRoomsGetSelector) GetOffset() OptInt {
	return s.Offset
}

// SetOrder sets the value of Order.
func (s *MeRoomsGetSelector) SetOrder(val OptOrder) {
	s.Order = val
}

// SetLimit sets the value of Limit.
func (s *MeRoomsGetSelector) SetLimit(val OptInt) {
	s.Limit = val
}

// SetOffset sets the value of Offset.
func (s *MeRoomsGetSelector) SetOffset(val OptInt) {
	s.Offset = val
}

type MeRoomsIdMessagesGetSelector struct {
	Order  OptOrder `json:"order"`
	Limit  OptInt   `json:"limit"`
	Offset OptInt   `json:"offset"`
}

// GetOrder returns the value of Order.
func (s *MeRoomsIdMessagesGetSelector) GetOrder() OptOrder {
	return s.Order
}

// GetLimit returns the value of Limit.
func (s *MeRoomsIdMessagesGetSelector) GetLimit() OptInt {
	return s.Limit
}

// GetOffset returns the value of Offset.
func (s *MeRoomsIdMessagesGetSelector) GetOffset() OptInt {
	return s.Offset
}

// SetOrder sets the value of Order.
func (s *MeRoomsIdMessagesGetSelector) SetOrder(val OptOrder) {
	s.Order = val
}

// SetLimit sets the value of Limit.
func (s *MeRoomsIdMessagesGetSelector) SetLimit(val OptInt) {
	s.Limit = val
}

// SetOffset sets the value of Offset.
func (s *MeRoomsIdMessagesGetSelector) SetOffset(val OptInt) {
	s.Offset = val
}

type MeRoomsPostReq struct {
	Name        string    `json:"name"`
	Description OptString `json:"description"`
}

// GetName returns the value of Name.
func (s *MeRoomsPostReq) GetName() string {
	return s.Name
}

// GetDescription returns the value of Description.
func (s *MeRoomsPostReq) GetDescription() OptString {
	return s.Description
}

// SetName sets the value of Name.
func (s *MeRoomsPostReq) SetName(val string) {
	s.Name = val
}

// SetDescription sets the value of Description.
func (s *MeRoomsPostReq) SetDescription(val OptString) {
	s.Description = val
}

type MeRoomsPutReq struct {
	ID          string    `json:"id"`
	Name        OptString `json:"name"`
	Description OptString `json:"description"`
}

// GetID returns the value of ID.
func (s *MeRoomsPutReq) GetID() string {
	return s.ID
}

// GetName returns the value of Name.
func (s *MeRoomsPutReq) GetName() OptString {
	return s.Name
}

// GetDescription returns the value of Description.
func (s *MeRoomsPutReq) GetDescription() OptString {
	return s.Description
}

// SetID sets the value of ID.
func (s *MeRoomsPutReq) SetID(val string) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *MeRoomsPutReq) SetName(val OptString) {
	s.Name = val
}

// SetDescription sets the value of Description.
func (s *MeRoomsPutReq) SetDescription(val OptString) {
	s.Description = val
}

// NewOptInt returns new OptInt with value set to v.
func NewOptInt(v int) OptInt {
	return OptInt{
		Value: v,
		Set:   true,
	}
}

// OptInt is optional int.
type OptInt struct {
	Value int
	Set   bool
}

// IsSet returns true if OptInt was set.
func (o OptInt) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt) Reset() {
	var v int
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt) SetTo(v int) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt) Get() (v int, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt) Or(d int) int {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptOrder returns new OptOrder with value set to v.
func NewOptOrder(v Order) OptOrder {
	return OptOrder{
		Value: v,
		Set:   true,
	}
}

// OptOrder is optional Order.
type OptOrder struct {
	Value Order
	Set   bool
}

// IsSet returns true if OptOrder was set.
func (o OptOrder) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptOrder) Reset() {
	var v Order
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptOrder) SetTo(v Order) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptOrder) Get() (v Order, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptOrder) Or(d Order) Order {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// Ref: #/components/schemas/order
type Order string

const (
	OrderAsc  Order = "asc"
	OrderDesc Order = "desc"
)

// AllValues returns all Order values.
func (Order) AllValues() []Order {
	return []Order{
		OrderAsc,
		OrderDesc,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s Order) MarshalText() ([]byte, error) {
	switch s {
	case OrderAsc:
		return []byte(s), nil
	case OrderDesc:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *Order) UnmarshalText(data []byte) error {
	switch Order(data) {
	case OrderAsc:
		*s = OrderAsc
		return nil
	case OrderDesc:
		*s = OrderDesc
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// Ref: #/components/schemas/room
type Room struct {
	ID            string    `json:"id"`
	CreatorUserID string    `json:"creator_user_id"`
	Name          string    `json:"name"`
	Description   OptString `json:"description"`
	CreatedAt     time.Time `json:"created_at"`
}

// GetID returns the value of ID.
func (s *Room) GetID() string {
	return s.ID
}

// GetCreatorUserID returns the value of CreatorUserID.
func (s *Room) GetCreatorUserID() string {
	return s.CreatorUserID
}

// GetName returns the value of Name.
func (s *Room) GetName() string {
	return s.Name
}

// GetDescription returns the value of Description.
func (s *Room) GetDescription() OptString {
	return s.Description
}

// GetCreatedAt returns the value of CreatedAt.
func (s *Room) GetCreatedAt() time.Time {
	return s.CreatedAt
}

// SetID sets the value of ID.
func (s *Room) SetID(val string) {
	s.ID = val
}

// SetCreatorUserID sets the value of CreatorUserID.
func (s *Room) SetCreatorUserID(val string) {
	s.CreatorUserID = val
}

// SetName sets the value of Name.
func (s *Room) SetName(val string) {
	s.Name = val
}

// SetDescription sets the value of Description.
func (s *Room) SetDescription(val OptString) {
	s.Description = val
}

// SetCreatedAt sets the value of CreatedAt.
func (s *Room) SetCreatedAt(val time.Time) {
	s.CreatedAt = val
}

// Ref: #/components/schemas/room_message
type RoomMessage struct {
	ID        string    `json:"id"`
	RoomID    string    `json:"room_id"`
	UserID    string    `json:"user_id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

// GetID returns the value of ID.
func (s *RoomMessage) GetID() string {
	return s.ID
}

// GetRoomID returns the value of RoomID.
func (s *RoomMessage) GetRoomID() string {
	return s.RoomID
}

// GetUserID returns the value of UserID.
func (s *RoomMessage) GetUserID() string {
	return s.UserID
}

// GetText returns the value of Text.
func (s *RoomMessage) GetText() string {
	return s.Text
}

// GetCreatedAt returns the value of CreatedAt.
func (s *RoomMessage) GetCreatedAt() time.Time {
	return s.CreatedAt
}

// SetID sets the value of ID.
func (s *RoomMessage) SetID(val string) {
	s.ID = val
}

// SetRoomID sets the value of RoomID.
func (s *RoomMessage) SetRoomID(val string) {
	s.RoomID = val
}

// SetUserID sets the value of UserID.
func (s *RoomMessage) SetUserID(val string) {
	s.UserID = val
}

// SetText sets the value of Text.
func (s *RoomMessage) SetText(val string) {
	s.Text = val
}

// SetCreatedAt sets the value of CreatedAt.
func (s *RoomMessage) SetCreatedAt(val time.Time) {
	s.CreatedAt = val
}

type RoomMessages struct {
	RoomMessages []RoomMessage `json:"room_messages"`
}

// GetRoomMessages returns the value of RoomMessages.
func (s *RoomMessages) GetRoomMessages() []RoomMessage {
	return s.RoomMessages
}

// SetRoomMessages sets the value of RoomMessages.
func (s *RoomMessages) SetRoomMessages(val []RoomMessage) {
	s.RoomMessages = val
}

type Rooms struct {
	Rooms []Room `json:"rooms"`
}

// GetRooms returns the value of Rooms.
func (s *Rooms) GetRooms() []Room {
	return s.Rooms
}

// SetRooms sets the value of Rooms.
func (s *Rooms) SetRooms(val []Room) {
	s.Rooms = val
}

type SignInPostOK struct {
	UserResponse User   `json:"user_response"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// GetUserResponse returns the value of UserResponse.
func (s *SignInPostOK) GetUserResponse() User {
	return s.UserResponse
}

// GetAccessToken returns the value of AccessToken.
func (s *SignInPostOK) GetAccessToken() string {
	return s.AccessToken
}

// GetRefreshToken returns the value of RefreshToken.
func (s *SignInPostOK) GetRefreshToken() string {
	return s.RefreshToken
}

// SetUserResponse sets the value of UserResponse.
func (s *SignInPostOK) SetUserResponse(val User) {
	s.UserResponse = val
}

// SetAccessToken sets the value of AccessToken.
func (s *SignInPostOK) SetAccessToken(val string) {
	s.AccessToken = val
}

// SetRefreshToken sets the value of RefreshToken.
func (s *SignInPostOK) SetRefreshToken(val string) {
	s.RefreshToken = val
}

type SignInPostReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// GetEmail returns the value of Email.
func (s *SignInPostReq) GetEmail() string {
	return s.Email
}

// GetPassword returns the value of Password.
func (s *SignInPostReq) GetPassword() string {
	return s.Password
}

// SetEmail sets the value of Email.
func (s *SignInPostReq) SetEmail(val string) {
	s.Email = val
}

// SetPassword sets the value of Password.
func (s *SignInPostReq) SetPassword(val string) {
	s.Password = val
}

type SignRefreshPostOK struct {
	AccessToken string `json:"access_token"`
}

// GetAccessToken returns the value of AccessToken.
func (s *SignRefreshPostOK) GetAccessToken() string {
	return s.AccessToken
}

// SetAccessToken sets the value of AccessToken.
func (s *SignRefreshPostOK) SetAccessToken(val string) {
	s.AccessToken = val
}

type SignRefreshPostReq struct {
	RefreshToken string `json:"refresh_token"`
}

// GetRefreshToken returns the value of RefreshToken.
func (s *SignRefreshPostReq) GetRefreshToken() string {
	return s.RefreshToken
}

// SetRefreshToken sets the value of RefreshToken.
func (s *SignRefreshPostReq) SetRefreshToken(val string) {
	s.RefreshToken = val
}

// SignUpPostCreated is response for SignUpPost operation.
type SignUpPostCreated struct{}

type SignUpPostReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// GetUsername returns the value of Username.
func (s *SignUpPostReq) GetUsername() string {
	return s.Username
}

// GetEmail returns the value of Email.
func (s *SignUpPostReq) GetEmail() string {
	return s.Email
}

// GetPassword returns the value of Password.
func (s *SignUpPostReq) GetPassword() string {
	return s.Password
}

// SetUsername sets the value of Username.
func (s *SignUpPostReq) SetUsername(val string) {
	s.Username = val
}

// SetEmail sets the value of Email.
func (s *SignUpPostReq) SetEmail(val string) {
	s.Email = val
}

// SetPassword sets the value of Password.
func (s *SignUpPostReq) SetPassword(val string) {
	s.Password = val
}

// Ref: #/components/schemas/user
type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Fullname  OptString `json:"fullname"`
	Status    OptString `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

// GetID returns the value of ID.
func (s *User) GetID() string {
	return s.ID
}

// GetUsername returns the value of Username.
func (s *User) GetUsername() string {
	return s.Username
}

// GetEmail returns the value of Email.
func (s *User) GetEmail() string {
	return s.Email
}

// GetFullname returns the value of Fullname.
func (s *User) GetFullname() OptString {
	return s.Fullname
}

// GetStatus returns the value of Status.
func (s *User) GetStatus() OptString {
	return s.Status
}

// GetCreatedAt returns the value of CreatedAt.
func (s *User) GetCreatedAt() time.Time {
	return s.CreatedAt
}

// SetID sets the value of ID.
func (s *User) SetID(val string) {
	s.ID = val
}

// SetUsername sets the value of Username.
func (s *User) SetUsername(val string) {
	s.Username = val
}

// SetEmail sets the value of Email.
func (s *User) SetEmail(val string) {
	s.Email = val
}

// SetFullname sets the value of Fullname.
func (s *User) SetFullname(val OptString) {
	s.Fullname = val
}

// SetStatus sets the value of Status.
func (s *User) SetStatus(val OptString) {
	s.Status = val
}

// SetCreatedAt sets the value of CreatedAt.
func (s *User) SetCreatedAt(val time.Time) {
	s.CreatedAt = val
}
