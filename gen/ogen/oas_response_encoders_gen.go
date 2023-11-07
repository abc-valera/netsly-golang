// Code generated by ogen, DO NOT EDIT.

package ogen

import (
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	ht "github.com/ogen-go/ogen/http"
)

func encodeCommentsByJokeIDGetResponse(response *Comments, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))

	e := new(jx.Encoder)
	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeLikesByJokeIDGetResponse(response int, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))

	e := new(jx.Encoder)
	e.Int(response)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeMeCommentsDelResponse(response *MeCommentsDelNoContent, w http.ResponseWriter, span trace.Span) error {
	w.WriteHeader(204)
	span.SetStatus(codes.Ok, http.StatusText(204))

	return nil
}

func encodeMeCommentsPostResponse(response *MeCommentsPostOK, w http.ResponseWriter, span trace.Span) error {
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))

	return nil
}

func encodeMeCommentsPutResponse(response *MeCommentsPutOK, w http.ResponseWriter, span trace.Span) error {
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))

	return nil
}

func encodeMeDelResponse(response *MeDelNoContent, w http.ResponseWriter, span trace.Span) error {
	w.WriteHeader(204)
	span.SetStatus(codes.Ok, http.StatusText(204))

	return nil
}

func encodeMeGetResponse(response *User, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))

	e := new(jx.Encoder)
	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeMeJokesDelResponse(response *MeJokesDelNoContent, w http.ResponseWriter, span trace.Span) error {
	w.WriteHeader(204)
	span.SetStatus(codes.Ok, http.StatusText(204))

	return nil
}

func encodeMeJokesGetResponse(response *Jokes, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))

	e := new(jx.Encoder)
	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeMeJokesPostResponse(response *MeJokesPostCreated, w http.ResponseWriter, span trace.Span) error {
	w.WriteHeader(201)
	span.SetStatus(codes.Ok, http.StatusText(201))

	return nil
}

func encodeMeJokesPutResponse(response *MeJokesPutCreated, w http.ResponseWriter, span trace.Span) error {
	w.WriteHeader(201)
	span.SetStatus(codes.Ok, http.StatusText(201))

	return nil
}

func encodeMeLikesDelResponse(response *MeLikesDelNoContent, w http.ResponseWriter, span trace.Span) error {
	w.WriteHeader(204)
	span.SetStatus(codes.Ok, http.StatusText(204))

	return nil
}

func encodeMeLikesPostResponse(response *MeLikesPostCreated, w http.ResponseWriter, span trace.Span) error {
	w.WriteHeader(201)
	span.SetStatus(codes.Ok, http.StatusText(201))

	return nil
}

func encodeMePutResponse(response *MePutCreated, w http.ResponseWriter, span trace.Span) error {
	w.WriteHeader(201)
	span.SetStatus(codes.Ok, http.StatusText(201))

	return nil
}

func encodeSignInPostResponse(response *SignInPostOK, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))

	e := new(jx.Encoder)
	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeSignRefreshPostResponse(response *SignRefreshPostOK, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))

	e := new(jx.Encoder)
	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeSignUpPostResponse(response *SignUpPostCreated, w http.ResponseWriter, span trace.Span) error {
	w.WriteHeader(201)
	span.SetStatus(codes.Ok, http.StatusText(201))

	return nil
}

func encodeErrorResponse(response *CodeErrorStatusCode, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	code := response.StatusCode
	if code == 0 {
		// Set default status code.
		code = http.StatusOK
	}
	w.WriteHeader(code)
	st := http.StatusText(code)
	if code >= http.StatusBadRequest {
		span.SetStatus(codes.Error, st)
	} else {
		span.SetStatus(codes.Ok, st)
	}

	e := new(jx.Encoder)
	response.Response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	if code >= http.StatusInternalServerError {
		return errors.Wrapf(ht.ErrInternalServerErrorResponse, "code: %d, message: %s", code, http.StatusText(code))
	}
	return nil

}
