package interceptor

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

type UnaryClient func(ctx *context.Context, method string, req, reply any, cc *grpc.ClientConn, opts ...grpc.CallOption) error
func ChainClient (interceptors ...UnaryClient) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		start := time.Now()
		
		for _, i := range interceptors {
			err := i(&ctx, method, req, reply, cc)
			if err != nil {
				return err
			}
		}
	
		err := invoker(ctx, method, req, reply, cc, opts...)
		
		log.Printf("Invoked RPC - Method:%s\tDuration:%s\tError:%v\n", method, time.Since(start), err)

		return err
	}
}