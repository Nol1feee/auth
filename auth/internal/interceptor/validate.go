package interceptor

import (
	"context"
	"google.golang.org/grpc"
)

type Validator interface {
	Validate() error
}

func InterceptorValidate(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	if val, ok := req.(Validator); ok {
		if err := val.Validate(); err != nil {
			return nil, err
		}
	}
	return handler(ctx, req)
}
