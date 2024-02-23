// Code generated by ogen, DO NOT EDIT.

package ogen

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// CommentsByJokeIDGet implements CommentsByJokeIDGet operation.
	//
	// Returns comments of the joke.
	//
	// GET /comments/{joke_id}
	CommentsByJokeIDGet(ctx context.Context, params CommentsByJokeIDGetParams) (*Comments, error)
	// LikesByJokeIDGet implements LikesByJokeIDGet operation.
	//
	// Counts likes of the joke.
	//
	// GET /likes/{joke_id}
	LikesByJokeIDGet(ctx context.Context, params LikesByJokeIDGetParams) (int, error)
	// MeCommentsDel implements MeCommentsDel operation.
	//
	// Deletes a comment of the current user.
	//
	// DELETE /me/comments
	MeCommentsDel(ctx context.Context, req *MeCommentsDelReq) error
	// MeCommentsPost implements MeCommentsPost operation.
	//
	// Creates a comment for the current user and the current joke.
	//
	// POST /me/comments
	MeCommentsPost(ctx context.Context, req *MeCommentsPostReq) (*Comment, error)
	// MeCommentsPut implements MeCommentsPut operation.
	//
	// Updates a comment of the current user.
	//
	// PUT /me/comments
	MeCommentsPut(ctx context.Context, req *MeCommentsPutReq) (*Comment, error)
	// MeDel implements MeDel operation.
	//
	// Deletes current user profile.
	//
	// DELETE /me
	MeDel(ctx context.Context, req *MeDelReq) error
	// MeGet implements MeGet operation.
	//
	// Returns current user profile.
	//
	// GET /me
	MeGet(ctx context.Context) (*User, error)
	// MeJokesDel implements MeJokesDel operation.
	//
	// Deletes joke for current user.
	//
	// DELETE /me/jokes
	MeJokesDel(ctx context.Context, req *MeJokesDelReq) error
	// MeJokesGet implements MeJokesGet operation.
	//
	// Returns jokes of the current user.
	//
	// GET /me/jokes
	MeJokesGet(ctx context.Context, params MeJokesGetParams) (*Jokes, error)
	// MeJokesPost implements MeJokesPost operation.
	//
	// Creates a new joke for current user.
	//
	// POST /me/jokes
	MeJokesPost(ctx context.Context, req *MeJokesPostReq) (*Joke, error)
	// MeJokesPut implements MeJokesPut operation.
	//
	// Updates joke for current user.
	//
	// PUT /me/jokes
	MeJokesPut(ctx context.Context, req *MeJokesPutReq) (*Joke, error)
	// MeLikesDel implements MeLikesDel operation.
	//
	// Deletes a like of the current user.
	//
	// DELETE /me/likes
	MeLikesDel(ctx context.Context, req *MeLikesDelReq) error
	// MeLikesPost implements MeLikesPost operation.
	//
	// Creates a like for a joke for the current user.
	//
	// POST /me/likes
	MeLikesPost(ctx context.Context, req *MeLikesPostReq) error
	// MePut implements MePut operation.
	//
	// Updates current user profile.
	//
	// PUT /me
	MePut(ctx context.Context, req *MePutReq) (*User, error)
	// SignInPost implements SignInPost operation.
	//
	// Performs user authentication.
	//
	// POST /sign/in
	SignInPost(ctx context.Context, req *SignInPostReq) (*SignInPostOK, error)
	// SignRefreshPost implements SignRefreshPost operation.
	//
	// Exchanges a refresh token for an access token.
	//
	// POST /sign/refresh
	SignRefreshPost(ctx context.Context, req *SignRefreshPostReq) (*SignRefreshPostOK, error)
	// SignUpPost implements SignUpPost operation.
	//
	// Performs user registration.
	//
	// POST /sign/up
	SignUpPost(ctx context.Context, req *SignUpPostReq) error
	// NewError creates *CodeErrorStatusCode from error returned by handler.
	//
	// Used for common default response.
	NewError(ctx context.Context, err error) *CodeErrorStatusCode
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h   Handler
	sec SecurityHandler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, sec SecurityHandler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		sec:        sec,
		baseServer: s,
	}, nil
}
