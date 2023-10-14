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
	Code    CodeErrorCode `json:"code"`
	Message string        `json:"message"`
}

// GetCode returns the value of Code.
func (s *CodeError) GetCode() CodeErrorCode {
	return s.Code
}

// GetMessage returns the value of Message.
func (s *CodeError) GetMessage() string {
	return s.Message
}

// SetCode sets the value of Code.
func (s *CodeError) SetCode(val CodeErrorCode) {
	s.Code = val
}

// SetMessage sets the value of Message.
func (s *CodeError) SetMessage(val string) {
	s.Message = val
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

type CommentsByJokeIDGetSelectParams struct {
	OrderBy OptString `json:"order_by"`
	Order   OptOrder  `json:"order"`
	Limit   int       `json:"limit"`
	Offset  int       `json:"offset"`
}

// GetOrderBy returns the value of OrderBy.
func (s *CommentsByJokeIDGetSelectParams) GetOrderBy() OptString {
	return s.OrderBy
}

// GetOrder returns the value of Order.
func (s *CommentsByJokeIDGetSelectParams) GetOrder() OptOrder {
	return s.Order
}

// GetLimit returns the value of Limit.
func (s *CommentsByJokeIDGetSelectParams) GetLimit() int {
	return s.Limit
}

// GetOffset returns the value of Offset.
func (s *CommentsByJokeIDGetSelectParams) GetOffset() int {
	return s.Offset
}

// SetOrderBy sets the value of OrderBy.
func (s *CommentsByJokeIDGetSelectParams) SetOrderBy(val OptString) {
	s.OrderBy = val
}

// SetOrder sets the value of Order.
func (s *CommentsByJokeIDGetSelectParams) SetOrder(val OptOrder) {
	s.Order = val
}

// SetLimit sets the value of Limit.
func (s *CommentsByJokeIDGetSelectParams) SetLimit(val int) {
	s.Limit = val
}

// SetOffset sets the value of Offset.
func (s *CommentsByJokeIDGetSelectParams) SetOffset(val int) {
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

// MeCommentsPostOK is response for MeCommentsPost operation.
type MeCommentsPostOK struct{}

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

type MeCommentsPostSelectParams struct {
	OrderBy OptString `json:"order_by"`
	Order   OptOrder  `json:"order"`
	Limit   int       `json:"limit"`
	Offset  int       `json:"offset"`
}

// GetOrderBy returns the value of OrderBy.
func (s *MeCommentsPostSelectParams) GetOrderBy() OptString {
	return s.OrderBy
}

// GetOrder returns the value of Order.
func (s *MeCommentsPostSelectParams) GetOrder() OptOrder {
	return s.Order
}

// GetLimit returns the value of Limit.
func (s *MeCommentsPostSelectParams) GetLimit() int {
	return s.Limit
}

// GetOffset returns the value of Offset.
func (s *MeCommentsPostSelectParams) GetOffset() int {
	return s.Offset
}

// SetOrderBy sets the value of OrderBy.
func (s *MeCommentsPostSelectParams) SetOrderBy(val OptString) {
	s.OrderBy = val
}

// SetOrder sets the value of Order.
func (s *MeCommentsPostSelectParams) SetOrder(val OptOrder) {
	s.Order = val
}

// SetLimit sets the value of Limit.
func (s *MeCommentsPostSelectParams) SetLimit(val int) {
	s.Limit = val
}

// SetOffset sets the value of Offset.
func (s *MeCommentsPostSelectParams) SetOffset(val int) {
	s.Offset = val
}

// MeCommentsPutOK is response for MeCommentsPut operation.
type MeCommentsPutOK struct{}

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
	UserID string `json:"user_id"`
}

// GetUserID returns the value of UserID.
func (s *MeDelReq) GetUserID() string {
	return s.UserID
}

// SetUserID sets the value of UserID.
func (s *MeDelReq) SetUserID(val string) {
	s.UserID = val
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

type MeJokesGetSelectParams struct {
	OrderBy OptString `json:"order_by"`
	Order   OptOrder  `json:"order"`
	Limit   int       `json:"limit"`
	Offset  int       `json:"offset"`
}

// GetOrderBy returns the value of OrderBy.
func (s *MeJokesGetSelectParams) GetOrderBy() OptString {
	return s.OrderBy
}

// GetOrder returns the value of Order.
func (s *MeJokesGetSelectParams) GetOrder() OptOrder {
	return s.Order
}

// GetLimit returns the value of Limit.
func (s *MeJokesGetSelectParams) GetLimit() int {
	return s.Limit
}

// GetOffset returns the value of Offset.
func (s *MeJokesGetSelectParams) GetOffset() int {
	return s.Offset
}

// SetOrderBy sets the value of OrderBy.
func (s *MeJokesGetSelectParams) SetOrderBy(val OptString) {
	s.OrderBy = val
}

// SetOrder sets the value of Order.
func (s *MeJokesGetSelectParams) SetOrder(val OptOrder) {
	s.Order = val
}

// SetLimit sets the value of Limit.
func (s *MeJokesGetSelectParams) SetLimit(val int) {
	s.Limit = val
}

// SetOffset sets the value of Offset.
func (s *MeJokesGetSelectParams) SetOffset(val int) {
	s.Offset = val
}

// MeJokesPostCreated is response for MeJokesPost operation.
type MeJokesPostCreated struct{}

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

// MeJokesPutCreated is response for MeJokesPut operation.
type MeJokesPutCreated struct{}

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

// MePutNoContent is response for MePut operation.
type MePutNoContent struct{}

type MePutReq struct {
	UserID   string    `json:"user_id"`
	Username OptString `json:"username"`
	Fullname OptString `json:"fullname"`
	Status   OptString `json:"status"`
}

// GetUserID returns the value of UserID.
func (s *MePutReq) GetUserID() string {
	return s.UserID
}

// GetUsername returns the value of Username.
func (s *MePutReq) GetUsername() OptString {
	return s.Username
}

// GetFullname returns the value of Fullname.
func (s *MePutReq) GetFullname() OptString {
	return s.Fullname
}

// GetStatus returns the value of Status.
func (s *MePutReq) GetStatus() OptString {
	return s.Status
}

// SetUserID sets the value of UserID.
func (s *MePutReq) SetUserID(val string) {
	s.UserID = val
}

// SetUsername sets the value of Username.
func (s *MePutReq) SetUsername(val OptString) {
	s.Username = val
}

// SetFullname sets the value of Fullname.
func (s *MePutReq) SetFullname(val OptString) {
	s.Fullname = val
}

// SetStatus sets the value of Status.
func (s *MePutReq) SetStatus(val OptString) {
	s.Status = val
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

type SignInPostOK struct {
	UserResponse User   `json:"userResponse"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
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
	AccessToken string `json:"accessToken"`
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
	RefreshToken string `json:"refreshToken"`
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
