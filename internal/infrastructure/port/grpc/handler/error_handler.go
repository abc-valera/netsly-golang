package handler

import (
	"github.com/abc-valera/flugo-api-golang/internal/domain/codeerr"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func handleErr(err error) error {
	if err == nil {
		return nil
	}

	code := codeerr.ErrorCode(err)
	msg := codeerr.ErrorMessage(err)

	if code == codeerr.CodeInvalidArgument {
		return status.Error(codes.InvalidArgument, msg)
	}

	if code == codeerr.CodeNotFound {
		return status.Error(codes.NotFound, msg)
	}

	if code == codeerr.CodeAlreadyExists {
		return status.Error(codes.AlreadyExists, msg)
	}

	if code == codeerr.CodePermissionDenied {
		return status.Error(codes.PermissionDenied, msg)
	}

	if code == codeerr.CodeUnauthenticated {
		return status.Error(codes.Unauthenticated, msg)
	}

	return status.Error(codes.Internal, msg)
}
