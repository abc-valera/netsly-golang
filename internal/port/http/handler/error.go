package handler

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/gen/ogen"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/codeerr"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/service"
)

type ErrorHandler struct {
}

func NewErrorHandler() ErrorHandler {
	return ErrorHandler{}
}

func (h ErrorHandler) NewError(ctx context.Context, err error) *ogen.CodeErrorStatusCode {
	code := codeerr.ErrorCode(err)
	codeError := ogen.CodeError{
		Code:    ogen.CodeErrorCode(code),
		Message: codeerr.ErrorMessage(err),
	}

	if code == codeerr.CodeInvalidArgument ||
		code == codeerr.CodeNotFound ||
		code == codeerr.CodeAlreadyExists {
		return &ogen.CodeErrorStatusCode{
			StatusCode: 400,
			Response:   codeError,
		}
	}

	if code == codeerr.CodePermissionDenied ||
		code == codeerr.CodeUnauthenticated {
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
