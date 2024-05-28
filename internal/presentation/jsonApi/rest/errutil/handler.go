package errutil

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/ogen"
	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
)

type Handler struct{}

func NewHandler() Handler {
	return Handler{}
}

func (h Handler) NewError(ctx context.Context, err error) *ogen.CodeErrorStatusCode {
	var codeError ogen.CodeError
	if coderr.ErrorCode(err) == coderr.CodeInternal {
		codeError = ogen.CodeError{
			Code:         "internal",
			ErrorMessage: "Internal error",
		}
	} else {
		codeError = ogen.CodeError{
			Code:         ogen.CodeErrorCode(coderr.ErrorCode(err)),
			ErrorMessage: err.Error(),
		}
	}

	switch coderr.ErrorCode(err) {
	case coderr.CodeInvalidArgument, coderr.CodeNotFound, coderr.CodeAlreadyExists:
		return &ogen.CodeErrorStatusCode{
			StatusCode: 400,
			Response:   codeError,
		}
	case coderr.CodeUnauthenticated:
		return &ogen.CodeErrorStatusCode{
			StatusCode: 401,
			Response:   codeError,
		}
	case coderr.CodePermissionDenied:
		return &ogen.CodeErrorStatusCode{
			StatusCode: 403,
			Response:   codeError,
		}
	default:
		global.Log().Error("REQUEST_ERROR", "err", err.Error())
		return &ogen.CodeErrorStatusCode{
			StatusCode: 500,
			Response:   codeError,
		}
	}
}
