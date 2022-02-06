package interceptor

import (
	"context"
	"google.golang.org/grpc"
)

func AuthInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		return handler(ctx, req) // Do RPC
	}
}
