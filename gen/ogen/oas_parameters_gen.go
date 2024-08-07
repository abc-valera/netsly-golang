// Code generated by ogen, DO NOT EDIT.

package ogen

import (
	"net/http"
	"net/url"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// CommentsByJokeIDGetParams is parameters of CommentsByJokeIDGet operation.
type CommentsByJokeIDGetParams struct {
	// ID of the joke.
	JokeID string
	// Fields to specify select parameters.
	Selector QuerySelectParams
}

func unpackCommentsByJokeIDGetParams(packed middleware.Parameters) (params CommentsByJokeIDGetParams) {
	{
		key := middleware.ParameterKey{
			Name: "joke_id",
			In:   "path",
		}
		params.JokeID = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "selector",
			In:   "query",
		}
		params.Selector = packed[key].(QuerySelectParams)
	}
	return params
}

func decodeCommentsByJokeIDGetParams(args [1]string, argsEscaped bool, r *http.Request) (params CommentsByJokeIDGetParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode path: joke_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "joke_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.JokeID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "joke_id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode query: selector.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "selector",
			Style:   uri.QueryStyleForm,
			Explode: true,
			Fields:  []uri.QueryParameterObjectField{{Name: "order", Required: false}, {Name: "limit", Required: false}, {Name: "offset", Required: false}},
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				return params.Selector.DecodeURI(d)
			}); err != nil {
				return err
			}
			if err := func() error {
				if err := params.Selector.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "selector",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// LikesByJokeIDGetParams is parameters of LikesByJokeIDGet operation.
type LikesByJokeIDGetParams struct {
	// ID of the joke.
	JokeID string
}

func unpackLikesByJokeIDGetParams(packed middleware.Parameters) (params LikesByJokeIDGetParams) {
	{
		key := middleware.ParameterKey{
			Name: "joke_id",
			In:   "path",
		}
		params.JokeID = packed[key].(string)
	}
	return params
}

func decodeLikesByJokeIDGetParams(args [1]string, argsEscaped bool, r *http.Request) (params LikesByJokeIDGetParams, _ error) {
	// Decode path: joke_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "joke_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.JokeID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "joke_id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// MeJokesGetParams is parameters of MeJokesGet operation.
type MeJokesGetParams struct {
	// Fields to specify select parameters.
	Selector QuerySelectParams
}

func unpackMeJokesGetParams(packed middleware.Parameters) (params MeJokesGetParams) {
	{
		key := middleware.ParameterKey{
			Name: "selector",
			In:   "query",
		}
		params.Selector = packed[key].(QuerySelectParams)
	}
	return params
}

func decodeMeJokesGetParams(args [0]string, argsEscaped bool, r *http.Request) (params MeJokesGetParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: selector.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "selector",
			Style:   uri.QueryStyleForm,
			Explode: true,
			Fields:  []uri.QueryParameterObjectField{{Name: "order", Required: false}, {Name: "limit", Required: false}, {Name: "offset", Required: false}},
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				return params.Selector.DecodeURI(d)
			}); err != nil {
				return err
			}
			if err := func() error {
				if err := params.Selector.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "selector",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// MeRoomsGetParams is parameters of MeRoomsGet operation.
type MeRoomsGetParams struct {
	// Fields to specify select parameters.
	Selector QuerySelectParams
}

func unpackMeRoomsGetParams(packed middleware.Parameters) (params MeRoomsGetParams) {
	{
		key := middleware.ParameterKey{
			Name: "selector",
			In:   "query",
		}
		params.Selector = packed[key].(QuerySelectParams)
	}
	return params
}

func decodeMeRoomsGetParams(args [0]string, argsEscaped bool, r *http.Request) (params MeRoomsGetParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: selector.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "selector",
			Style:   uri.QueryStyleForm,
			Explode: true,
			Fields:  []uri.QueryParameterObjectField{{Name: "order", Required: false}, {Name: "limit", Required: false}, {Name: "offset", Required: false}},
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				return params.Selector.DecodeURI(d)
			}); err != nil {
				return err
			}
			if err := func() error {
				if err := params.Selector.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "selector",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// MeRoomsIdMessagesGetParams is parameters of MeRoomsIdMessagesGet operation.
type MeRoomsIdMessagesGetParams struct {
	// ID of the room.
	RoomID string
	// Fields to specify select parameters.
	Selector QuerySelectParams
}

func unpackMeRoomsIdMessagesGetParams(packed middleware.Parameters) (params MeRoomsIdMessagesGetParams) {
	{
		key := middleware.ParameterKey{
			Name: "room_id",
			In:   "path",
		}
		params.RoomID = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "selector",
			In:   "query",
		}
		params.Selector = packed[key].(QuerySelectParams)
	}
	return params
}

func decodeMeRoomsIdMessagesGetParams(args [1]string, argsEscaped bool, r *http.Request) (params MeRoomsIdMessagesGetParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode path: room_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "room_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.RoomID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "room_id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode query: selector.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "selector",
			Style:   uri.QueryStyleForm,
			Explode: true,
			Fields:  []uri.QueryParameterObjectField{{Name: "order", Required: false}, {Name: "limit", Required: false}, {Name: "offset", Required: false}},
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				return params.Selector.DecodeURI(d)
			}); err != nil {
				return err
			}
			if err := func() error {
				if err := params.Selector.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "selector",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}
