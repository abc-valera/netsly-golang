package handler

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/gen/ogen"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/coderr"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/service"
)

type ErrorHandler struct {
}

func NewErrorHandler() ErrorHandler {
	return ErrorHandler{}
}

func (h ErrorHandler) NewError(ctx context.Context, err error) *ogen.CodeErrorStatusCode {
	code := coderr.ErrorCode(err)
	codeError := ogen.CodeError{
		Code:    ogen.CodeErrorCode(code),
		Message: coderr.ErrorMessage(err),
	}

	if code == coderr.CodeInvalidArgument ||
		code == coderr.CodeNotFound ||
		code == coderr.CodeAlreadyExists {
		return &ogen.CodeErrorStatusCode{
			StatusCode: 400,
			Response:   codeError,
		}
	}

	if code == coderr.CodePermissionDenied ||
		code == coderr.CodeUnauthenticated {
		return &ogen.CodeErrorStatusCode{
			StatusCode: 401,
			Response:   codeError,
		}
	}

	service.Log.Error("REQUEST_ERROR", "err", err.Error())
	return &ogen.CodeErrorStatusCode{
		StatusCode: 500,
		Response:   codeError,
	}
}
