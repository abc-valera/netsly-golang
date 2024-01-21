package handler

import (
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/coderr"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func handleErr(err error) error {
	if err == nil {
		return nil
	}

	code := coderr.ErrorCode(err)
	msg := coderr.ErrorMessage(err)

	if code == coderr.CodeInvalidArgument {
		return status.Error(codes.InvalidArgument, msg)
	}

	if code == coderr.CodeNotFound {
		return status.Error(codes.NotFound, msg)
	}

	if code == coderr.CodeAlreadyExists {
		return status.Error(codes.AlreadyExists, msg)
	}

	if code == coderr.CodePermissionDenied {
		return status.Error(codes.PermissionDenied, msg)
	}

	if code == coderr.CodeUnauthenticated {
		return status.Error(codes.Unauthenticated, msg)
	}

	return status.Error(codes.Internal, msg)
}
