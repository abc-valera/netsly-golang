// Code generated by ogen, DO NOT EDIT.

package ogen

import (
	"fmt"
)

func (s *CodeErrorStatusCode) Error() string {
	return fmt.Sprintf("code %d: %+v", s.StatusCode, s.Response)
}

// Ref: #/components/schemas/CodeError
type CodeError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// GetCode returns the value of Code.
func (s *CodeError) GetCode() string {
	return s.Code
}

// GetMessage returns the value of Message.
func (s *CodeError) GetMessage() string {
	return s.Message
}

// SetCode sets the value of Code.
func (s *CodeError) SetCode(val string) {
	s.Code = val
}

// SetMessage sets the value of Message.
func (s *CodeError) SetMessage(val string) {
	s.Message = val
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

// Ref: #/components/schemas/SignInRequest
type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// GetEmail returns the value of Email.
func (s *SignInRequest) GetEmail() string {
	return s.Email
}

// GetPassword returns the value of Password.
func (s *SignInRequest) GetPassword() string {
	return s.Password
}

// SetEmail sets the value of Email.
func (s *SignInRequest) SetEmail(val string) {
	s.Email = val
}

// SetPassword sets the value of Password.
func (s *SignInRequest) SetPassword(val string) {
	s.Password = val
}

// Ref: #/components/schemas/SignInResponse
type SignInResponse struct {
	UserResponse UserResponse `json:"userResponse"`
	AccessToken  string       `json:"accessToken"`
	RefreshToken string       `json:"refreshToken"`
}

// GetUserResponse returns the value of UserResponse.
func (s *SignInResponse) GetUserResponse() UserResponse {
	return s.UserResponse
}

// GetAccessToken returns the value of AccessToken.
func (s *SignInResponse) GetAccessToken() string {
	return s.AccessToken
}

// GetRefreshToken returns the value of RefreshToken.
func (s *SignInResponse) GetRefreshToken() string {
	return s.RefreshToken
}

// SetUserResponse sets the value of UserResponse.
func (s *SignInResponse) SetUserResponse(val UserResponse) {
	s.UserResponse = val
}

// SetAccessToken sets the value of AccessToken.
func (s *SignInResponse) SetAccessToken(val string) {
	s.AccessToken = val
}

// SetRefreshToken sets the value of RefreshToken.
func (s *SignInResponse) SetRefreshToken(val string) {
	s.RefreshToken = val
}

// SignUpCreated is response for SignUp operation.
type SignUpCreated struct{}

// Ref: #/components/schemas/SignUpRequest
type SignUpRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// GetUsername returns the value of Username.
func (s *SignUpRequest) GetUsername() string {
	return s.Username
}

// GetEmail returns the value of Email.
func (s *SignUpRequest) GetEmail() string {
	return s.Email
}

// GetPassword returns the value of Password.
func (s *SignUpRequest) GetPassword() string {
	return s.Password
}

// SetUsername sets the value of Username.
func (s *SignUpRequest) SetUsername(val string) {
	s.Username = val
}

// SetEmail sets the value of Email.
func (s *SignUpRequest) SetEmail(val string) {
	s.Email = val
}

// SetPassword sets the value of Password.
func (s *SignUpRequest) SetPassword(val string) {
	s.Password = val
}

// Ref: #/components/schemas/UserResponse
type UserResponse struct {
	ID       string    `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Fullname OptString `json:"fullname"`
	Status   OptString `json:"status"`
}

// GetID returns the value of ID.
func (s *UserResponse) GetID() string {
	return s.ID
}

// GetUsername returns the value of Username.
func (s *UserResponse) GetUsername() string {
	return s.Username
}

// GetEmail returns the value of Email.
func (s *UserResponse) GetEmail() string {
	return s.Email
}

// GetFullname returns the value of Fullname.
func (s *UserResponse) GetFullname() OptString {
	return s.Fullname
}

// GetStatus returns the value of Status.
func (s *UserResponse) GetStatus() OptString {
	return s.Status
}

// SetID sets the value of ID.
func (s *UserResponse) SetID(val string) {
	s.ID = val
}

// SetUsername sets the value of Username.
func (s *UserResponse) SetUsername(val string) {
	s.Username = val
}

// SetEmail sets the value of Email.
func (s *UserResponse) SetEmail(val string) {
	s.Email = val
}

// SetFullname sets the value of Fullname.
func (s *UserResponse) SetFullname(val OptString) {
	s.Fullname = val
}

// SetStatus sets the value of Status.
func (s *UserResponse) SetStatus(val OptString) {
	s.Status = val
}
