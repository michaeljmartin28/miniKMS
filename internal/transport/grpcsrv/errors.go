package grpcsrv

import (
	"errors"

	"github.com/michaeljmartin28/minikms/internal/core"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func mapErrorToGRPC(err error) error {
	switch {
	case errors.Is(err, core.ErrBadAlgorithm):
		return status.Error(codes.InvalidArgument, err.Error())

	case errors.Is(err, core.ErrInvalidVersion):
		return status.Error(codes.InvalidArgument, err.Error())

	case errors.Is(err, core.ErrKeyNotFound):
		return status.Error(codes.NotFound, err.Error())

	case errors.Is(err, core.ErrKeyDisabled):
		return status.Error(codes.FailedPrecondition, err.Error())

	default:
		return status.Error(codes.Internal, "internal error")
	}
}
