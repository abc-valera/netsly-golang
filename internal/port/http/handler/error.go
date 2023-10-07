package handler

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/gen/ogen"
	"github.com/abc-valera/flugo-api-golang/internal/domain/codeerr"
	"github.com/abc-valera/flugo-api-golang/internal/domain/service"
)

type ErrorHandler struct {
	log service.Logger
}

func NewErrorHandler(log service.Logger) ErrorHandler {
	return ErrorHandler{
		log: log,
	}
}

func (h *ErrorHandler) NewError(ctx context.Context, err error) *ogen.CodeErrorStatusCode {
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

	h.log.Error("REQUEST_ERROR", "err", err.Error())
	return &ogen.CodeErrorStatusCode{
		StatusCode: 500,
		Response:   codeError,
	}
}
