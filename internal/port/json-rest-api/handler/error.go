package handler

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/ogen"
	"github.com/abc-valera/netsly-api-golang/internal/domain/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
)

type ErrorHandler struct {
}

func NewErrorHandler() ErrorHandler {
	return ErrorHandler{}
}

func (h ErrorHandler) NewError(ctx context.Context, err error) *ogen.CodeErrorStatusCode {
	codeError := ogen.CodeError{
		Code:         ogen.CodeErrorCode(coderr.ErrorCode(err)),
		ErrorMessage: coderr.ErrorMessage(err),
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
