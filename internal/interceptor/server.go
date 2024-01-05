package interceptor

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

type UnaryServer func(cctx context.Context, req any) (any, error)

func ChainServer (interceptors ...UnaryServer) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error){
		start := time.Now()
		
		for _, i := range interceptors {
			h, err := i(ctx, req)
			if err != nil {
				return nil, err
			}
			if h != nil {
				return h, nil
			}
		}
	
		h, err := handler(ctx, req)
		
		log.Printf("Request - Method:%s\tDuration:%s\tError:%v\n", info.FullMethod, time.Since(start), err)

		return h, err
	}
}