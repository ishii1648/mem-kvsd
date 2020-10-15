package app

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// server-side error
var (
	ErrGRPCEmptyKey = status.New(codes.InvalidArgument, "etcdserver: key is not provided").Err()
)
