package server

import (
	"context"

	"github.com/NaKa2355/aim/internal/app/aim/usecases/boundary"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func UnaryErrInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	res, err := handler(ctx, req)
	if err == nil {
		return res, err
	}

	if bdyErr, ok := err.(boundary.Error); ok {
		switch bdyErr.Code {
		case boundary.CodeAlreadyExists:
			err = status.Error(codes.AlreadyExists, err.Error())
		case boundary.CodeDatabase:
			err = status.Error(codes.Internal, err.Error())
		case boundary.CodeNotFound:
			err = status.Error(codes.NotFound, err.Error())
		case boundary.CodeInvaildInput:
			err = status.Error(codes.InvalidArgument, err.Error())
		case boundary.CodeInvaildOperation:
			err = status.Error(codes.Unimplemented, err.Error())
		}
	} else {
		err = status.Error(codes.Unknown, err.Error())
	}

	return res, err
}
