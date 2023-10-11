// Code generated by ogen, DO NOT EDIT.

package ogen

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// MeGet implements GET /me operation.
	//
	// Returns current user profile.
	//
	// GET /me
	MeGet(ctx context.Context) (*User, error)
	// MeJokesGet implements GET /me/jokes operation.
	//
	// Returns jokes of the current user.
	//
	// GET /me/jokes
	MeJokesGet(ctx context.Context) (*Jokes, error)
	// MeJokesPost implements POST /me/jokes operation.
	//
	// Creates a new joke for current user.
	//
	// POST /me/jokes
	MeJokesPost(ctx context.Context, req *MeJokesPostReq) error
	// SignInPost implements POST /sign_in operation.
	//
	// Performs user authentication.
	//
	// POST /sign_in
	SignInPost(ctx context.Context, req *SignInPostReq) (*SignInPostOK, error)
	// SignRefreshPost implements POST /sign_refresh operation.
	//
	// Exchanges a refresh token for an access token.
	//
	// POST /sign_refresh
	SignRefreshPost(ctx context.Context, req *SignRefreshPostReq) (*SignRefreshPostOK, error)
	// SignUpPost implements POST /sign_up operation.
	//
	// Performs user registration.
	//
	// POST /sign_up
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
