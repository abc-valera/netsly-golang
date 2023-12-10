package codeerr

import "errors"

// codeerr package provides a way to represent errors as a code and a message.

const (
	// CodeUnauthenticated indicates the request does not have valid authentication credentials for the operation
	CodeUnauthenticated Code = "unauthenticated"

	// CodeInvalidArgument indicates client specified an invalid argument
	CodeInvalidArgument Code = "invalid_argument"

	// CodeNotFound means requested entity was not found
	CodeNotFound Code = "not_found"

	// CodeAlreadyExists means an attempt to create an entity failed because one already exists
	CodeAlreadyExists Code = "already_exists"

	// CodePermissionDenied indicates the caller does not have permission to execute the specified operation
	CodePermissionDenied Code = "permission_denied"

	// CodeInternal means an internal error occured
	CodeInternal Code = "internal"
)

type Code string

// ErrorCode returns the code of the root error, if available, otherwise returns Internal.
func ErrorCode(err error) Code {
	var messageErrTarget *MessageErr
	var internalErrTarget *InternalErr
	if err == nil {
		return ""
	} else if e, ok := err.(*MessageErr); ok && e.Code != "" {
		return e.Code
	} else if e, ok := err.(*InternalErr); ok && e.Code != "" {
		return e.Code
	} else if errors.As(err, &messageErrTarget) {
		return messageErrTarget.Code
	} else if errors.As(err, &internalErrTarget) {
		return internalErrTarget.Code
	} else if ok && e.Err != nil {
		return ErrorCode(e.Err)
	}
	return CodeInternal
}
